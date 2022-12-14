package merchant

import (
	"database/sql"
	"fmt"
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
	connURL := *pgURL
	if connURL.Scheme == "cockroachdb" {
		// Overwrite the scheme before parsing with pgx, since
		// it doesn't support the "cockroachdb" scheme.
		connURL.Scheme = "merchant"
	}
	c, err := pgx.ParseConfig(connURL.String())
	if err != nil {
		return nil, fmt.Errorf("parsing merchant URI: %w", err)
	}

	c.Logger = logrusadapter.NewLogger(logger)
	db := stdlib.OpenDB(*c)

	err = validateSchema(db, pgURL.Scheme)
	if err != nil {
		Log().format("NewDirectory validateSchema error: %s\n", err.Error())
		return nil, fmt.Errorf("validating schema: %s", err.Error())
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
