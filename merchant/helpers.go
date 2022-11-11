package merchant

import (
	"database/sql"
	"embed"
	"fmt"
	"io"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/cockroachdb"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	merchpb "merchant-service/proto"
)

//go:embed migrations/*.sql
var fs embed.FS

// version defines the current migration version. This ensures the app
// is always compatible with the version of the database.
const version = 1

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

func partnerPostgresToProto(pgPartner Partner) (*merchpb.Partner, error) {
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
	return &merchpb.Partner{
		Id:         userID,
		CreateTime: timestamppb.New(pgPartner.CreateTime),
		FullName:   pgPartner.FullName,
		Url:        pgPartner.Url,
		ApiToken:   apiToken,
		Role:       protoRole,
	}, nil
}

func merchantPostgresToProto(pgMerchant Merchant) (*merchpb.Merchant, error) {
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
	return &merchpb.Merchant{
		Id:         merchantID,
		CreateTime: timestamppb.New(pgMerchant.CreateTime),
		FullName:   pgMerchant.FullName,
		Url:        pgMerchant.Url,
		PartnerId:  partnerID,
	}, nil
}

func rolePostgresToProto(pgRole Role) (merchpb.Role, error) {
	switch pgRole {
	case RoleBasePartner:
		return merchpb.Role_BASE_PARTNER, nil
	case RolePartner:
		return merchpb.Role_PARTNER, nil
	case RoleAdmin:
		return merchpb.Role_ADMIN, nil
	case RoleGuest:
		return merchpb.Role_GUEST, nil
	default:
		return 0, status.Errorf(codes.Internal, "unknown role type %q", pgRole)
	}
}

func roleProtoToPostgres(pbRole merchpb.Role) (Role, error) {
	switch pbRole {
	case merchpb.Role_GUEST:
		return RoleGuest, nil
	case merchpb.Role_ADMIN:
		return RoleAdmin, nil
	case merchpb.Role_PARTNER:
		return RolePartner, nil
	case merchpb.Role_BASE_PARTNER:
		return RoleBasePartner, nil
	default:
		return "", status.Errorf(codes.InvalidArgument, "unknown role type %q", pbRole)
	}
}

var _ pgx.CopyFromSource = (*partnersSource)(nil)

type partnersSource struct {
	getPartner func() (*merchpb.AddPartnerRequest, error)
	nextValues []interface{}
	err        error
}

func (p *partnersSource) Next() bool {
	if p.err != nil {
		return false
	}
	var req *merchpb.AddPartnerRequest
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
