package merchant

import (
	"database/sql"
	"embed"
	"fmt"
	"io"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/cockroachdb"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.mapcard.pro/external-map-team/api-proto/merchant/api"
)

//go:embed migrations/*.sql
var fs embed.FS

// version defines the current migration version. This ensures the app
// is always compatible with the version of the database.
const version = 3

// Migrate migrates the Postgres schema to the current version.
func validateSchema(db *sql.DB, scheme string) error {
	sourceInstance, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	var driverInstance database.Driver
	switch scheme {
	case "postgres", "postgresql":
		driverInstance, err = postgres.WithInstance(db, new(postgres.Config))
	case "cockroachdb":
		driverInstance, err = cockroachdb.WithInstance(db, new(cockroachdb.Config))
	default:
		return fmt.Errorf("unknown scheme: %q", scheme)
	}
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", sourceInstance, scheme, driverInstance)
	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}

func partnerPostgresToProto(pgPartner Partner) (*api.Partner, error) {
	protoRole, err := rolePostgresToProto(pgPartner.Role)
	if err != nil {
		return nil, err
	}
	var userID string
	err = pgPartner.ID.AssignTo(&userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign UUID to string: %s", err.Error())
	}
	var apiToken string
	err = pgPartner.ApiToken.AssignTo(&apiToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign Api token to string: %s", err.Error())
	}
	return &api.Partner{
		Id:         userID,
		CreateTime: timestamppb.New(pgPartner.CreateTime),
		FullName:   pgPartner.FullName,
		Url:        pgPartner.Url,
		ApiToken:   apiToken,
		Role:       protoRole,
	}, nil
}

func merchantPostgresToProto(pgMerchant Merchant) (*api.Merchant, error) {
	var merchantID string
	err := pgMerchant.ID.AssignTo(&merchantID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign merchant UUID to string: %s", err.Error())
	}
	var partnerID string
	err = pgMerchant.PartnerID.AssignTo(&partnerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign partner UUID to string: %s", err.Error())
	}
	return &api.Merchant{
		Id:         merchantID,
		CreateTime: timestamppb.New(pgMerchant.CreateTime),
		FullName:   pgMerchant.FullName,
		Url:        pgMerchant.Url,
		PartnerId:  partnerID,
	}, nil
}

func shopPostgresToProto(pgShop Shop) (*api.Shop, error) {
	var shopID string
	err := pgShop.ID.AssignTo(&shopID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign shop UUID to string: %s", err.Error())
	}
	var merchantID string
	err = pgShop.MerchantID.AssignTo(&merchantID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign merchant UUID to string: %s", err.Error())
	}
	return &api.Shop{
		Id:         shopID,
		CreateTime: timestamppb.New(pgShop.CreateTime),
		FullName:   pgShop.FullName,
		Url:        pgShop.Url,
		MerchantId: merchantID,
		Login:      pgShop.Login,
		Password:   pgShop.Password,
	}, nil
}

func terminalPostgresToProto(pgTerminal Terminal) (*api.Terminal, error) {
	var terminalID string
	err := pgTerminal.ID.AssignTo(&terminalID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign terminal UUID to string: %s", err.Error())
	}
	var shopID string
	err = pgTerminal.ShopID.AssignTo(&shopID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign shop UUID to string: %s", err.Error())
	}
	return &api.Terminal{
		Id:         shopID,
		CreateTime: timestamppb.New(pgTerminal.CreateTime),
		FullName:   pgTerminal.FullName,
		Url:        pgTerminal.Url,
		ShopId:     shopID,
		Login:      pgTerminal.Login,
		Password:   pgTerminal.Password,
	}, nil
}

func authinfoPostgresToProto(pgSellerinfo AuthRow) (*api.SellerInfo, error) {
	fmt.Println(pgSellerinfo.TerminalID.Status)
	if pgSellerinfo.TerminalID.Status == pgtype.Undefined {
		return &api.SellerInfo{
			Success:    false,
			ErrCode:    "NOT_FOUND",
			ErrMessage: "Не найдено: терминал",
		}, nil

	}
	protoRole, err := rolePostgresToProto(pgSellerinfo.PartnerRole)
	if err != nil {
		return nil, err
	}
	var partnerID string
	err = pgSellerinfo.PartnerID.AssignTo(&partnerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign partner UUID to string: %s", err.Error())
	}
	var merchantID string
	err = pgSellerinfo.MerchantID.AssignTo(&merchantID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign merchant UUID to string: %s", err.Error())
	}
	var shopID string
	err = pgSellerinfo.ShopID.AssignTo(&shopID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign shop UUID to string: %s", err.Error())
	}
	var terminalID string
	err = pgSellerinfo.TerminalID.AssignTo(&terminalID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to assign terminal UUID to string: %s", err.Error())
	}
	return &api.SellerInfo{
		Success: true,
		PartnerInfo: &api.Partner{
			Id:       partnerID,
			FullName: pgSellerinfo.PartnerFullName,
			Url:      pgSellerinfo.PartnerUrl,
			Role:     protoRole,
		},
		MerchantInfo: &api.Merchant{
			Id:       merchantID,
			FullName: pgSellerinfo.MerchantFullName,
			Url:      pgSellerinfo.MerchantUrl,
		},
		ShopInfo: &api.Shop{
			Id:       shopID,
			FullName: pgSellerinfo.ShopFullName,
			Url:      pgSellerinfo.ShopUrl,
		},
		TerminalInfo: &api.Terminal{
			Id:       terminalID,
			FullName: pgSellerinfo.TerminalFullName,
			Url:      pgSellerinfo.TerminalUrl,
			Login:    pgSellerinfo.TermnalLogin,
		},
	}, nil
}

func rolePostgresToProto(pgRole Role) (api.Role, error) {
	switch pgRole {
	case RoleBasePartner:
		return api.Role_BASE_PARTNER, nil
	case RolePartner:
		return api.Role_PARTNER, nil
	case RoleAdmin:
		return api.Role_ADMIN, nil
	case RoleGuest:
		return api.Role_GUEST, nil
	default:
		return 0, status.Errorf(codes.Internal, "unknown role type %q", pgRole)
	}
}

func roleProtoToPostgres(pbRole api.Role) (Role, error) {
	switch pbRole {
	case api.Role_GUEST:
		return RoleGuest, nil
	case api.Role_ADMIN:
		return RoleAdmin, nil
	case api.Role_PARTNER:
		return RolePartner, nil
	case api.Role_BASE_PARTNER:
		return RoleBasePartner, nil
	default:
		return "", status.Errorf(codes.InvalidArgument, "unknown role type %q", pbRole)
	}
}

var _ pgx.CopyFromSource = (*partnersSource)(nil)

type partnersSource struct {
	getPartner func() (*api.AddPartnerRequest, error)
	nextValues []interface{}
	err        error
}

func (p *partnersSource) Next() bool {
	if p.err != nil {
		return false
	}
	var req *api.AddPartnerRequest
	req, p.err = p.getPartner()
	if p.err != nil {
		return false
	}
	var pgRole Role
	pgRole, p.err = roleProtoToPostgres(req.Role)
	if p.err != nil {
		return false
	}
	p.nextValues = []interface{}{req.FullName, req.Url, pgRole}
	return true
}

func (p *partnersSource) Values() ([]interface{}, error) {
	return p.nextValues, nil
}

func (p *partnersSource) Err() error {
	if p.err == io.EOF {
		// This is actually success, so we don't want to return an error
		return nil
	}
	return p.err
}

const defaultPgDriver = "postgres"
const defaultPgPort = "5432"

func CreateURL_FromEnvParts() (string, error) {
	var pgUrl string = ""
	pgHost := os.Getenv("DB_HOST")
	if pgHost == "" {
		return "", fmt.Errorf("Postgres host must be set")
	}
	pgDriver := os.Getenv("DB_DRIVER")
	if pgDriver == "" {
		pgDriver = defaultPgDriver
	}
	pgUser := os.Getenv("DB_USER")
	if pgUser == "" {
		return "", fmt.Errorf("Postgres user must be set")
	}
	pgPassword := os.Getenv("DB_PASSWORD")
	if pgPassword == "" {
		return "", fmt.Errorf("Postgres password must be set")
	}
	pgDbName := os.Getenv("DB_NAME")
	if pgDbName == "" {
		return "", fmt.Errorf("Postgres database name must be set")
	}
	pgPort := os.Getenv("DB_PORT")
	if pgPort == "" {
		pgPort = defaultPgPort
	}
	pgUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		pgUser,
		pgPassword,
		pgHost,
		pgPort,
		pgDbName,
	)

	return pgUrl, nil
}
