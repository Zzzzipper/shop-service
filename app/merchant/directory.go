package merchant

import (
	"database/sql"
	"net/url"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

// Directory stores a directory of partners.
type Directory struct {
	logger  *logrus.Logger
	db      *sql.DB
	sb      squirrel.StatementBuilderType
	querier Querier
}

// NewDirectory creates a new Directory, connecting it to the postgres server on
// the URL provided.
func NewDirectory(logger *logrus.Logger, pgURL *url.URL) (*Directory, error) {

	Log().Format("Start NewDirectory, Postgres URL: %v\n", pgURL)

	c, err := pgx.ParseConfig(pgURL.String())
	if err != nil {
		Log().Format("Error parsing merchant URI: %v\n", err)
		return nil, Log().Errorf("Parsing merchant URI: %v", err)
	}

	c.Logger = logrusadapter.NewLogger(logger)
	db := stdlib.OpenDB(*c)

	err = validateSchema(db, pgURL.Scheme)
	if err != nil {
		Log().Format("NewDirectory validateSchema error: %s\n", err.Error())
		return nil, Log().Errorf("Validating schema: %s", err.Error())
	}

	return &Directory{
		logger:  logger,
		db:      db,
		sb:      squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
		querier: New(db),
	}, nil
}

// Close releases any resources.
func (d Directory) Close() error {
	return d.db.Close()
}
