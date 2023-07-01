package storagecore

import (
	"context"
	"feather-ai/service-core/lib/config"
	"log"
	"time"

	"feather-ai/service-core/lib/storage"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PostgresDB struct {
	DB *sqlx.DB
}

type PostgresPreparedStatement struct {
	Query *sqlx.NamedStmt
}

func NewPostgresDB(options config.Options) *PostgresDB {
	storage := &PostgresDB{}

	for i := 0; i < 10; i++ {
		// Connect to Postgres
		db, err := sqlx.Open("pgx", options.DBurl)
		if err != nil {
			log.Fatalln(err)
		}

		err = db.Ping()
		if err != nil {
			db.Close()
			logrus.Errorf("Cannot connect to DB (will retry): %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		storage.DB = db
		break
	}

	if storage.DB == nil {
		logrus.Fatalf("Cannot connect to DB at all! Aborting")
	}

	return storage
}

func (m *PostgresDB) NewTransaction(ctx context.Context) (storage.DBTransaction, error) {
	tx, err := m.DB.Beginx()
	if err != nil {
		return nil, err
	}

	return &PostgresTransaction{
		tx: tx,
	}, nil
}

func (m *PostgresDB) PrepareQuery(query string) storage.PreparedStatement {
	q, err := m.DB.PrepareNamed(query)
	abortIfFailed(err, query)

	return &PostgresPreparedStatement{
		Query: q,
	}
}

func (m *PostgresDB) DynamicQuery(ctx context.Context, query string, results interface{}) error {
	err := m.DB.SelectContext(ctx, results, query)
	if err != nil {
		logrus.Errorf("DB.DynamicQuery: %v", err)
		return err
	}
	return nil
}

/*
* Prepared statements
 */

func abortIfFailed(err error, name string) {
	if err != nil {
		logrus.Errorf("Could not prepare statement %s. Err=%v\n", name, err)
		time.Sleep(10 * time.Second)
		logrus.Fatal("Aborting")
	}
}

func (stmt *PostgresPreparedStatement) Write(ctx context.Context, args interface{}) error {
	_, err := stmt.Query.ExecContext(ctx, args)
	return err
}

func (stmt *PostgresPreparedStatement) WriteReturnID(ctx context.Context, args interface{}) (int64, error) {
	var rowID int64
	err := stmt.Query.QueryRowxContext(ctx, args).Scan(&rowID)
	if err != nil {
		return -1, err
	}

	return rowID, nil
}

func (stmt *PostgresPreparedStatement) QueryMany(ctx context.Context, args interface{}, results interface{}) error {
	err := stmt.Query.SelectContext(ctx, results, args)
	return err
}

/*
* Helpers for working with DB transactions.
 */
type PostgresTransaction struct {
	tx  *sqlx.Tx
	err error
}

func (m *PostgresTransaction) Execute(ctx context.Context, statement storage.PreparedStatement, args interface{}) {
	if m.err != nil {
		return
	}

	txStmt := m.tx.NamedStmtContext(ctx, statement.(*PostgresPreparedStatement).Query)
	_, m.err = txStmt.ExecContext(ctx, args)
}

func (m *PostgresTransaction) ExecuteMany(ctx context.Context, statement storage.PreparedStatement, args []interface{}) {
	if m.err != nil {
		return
	}

	txStmt := m.tx.NamedStmtContext(ctx, statement.(*PostgresPreparedStatement).Query)
	for arg := range args {
		_, m.err = txStmt.ExecContext(ctx, arg)

		if m.err != nil {
			return
		}
	}
}

func (m *PostgresTransaction) Commit(ctx context.Context) {
	if m.err != nil {
		m.tx.Rollback()
		return
	}
	m.err = m.tx.Commit()
}

func (m *PostgresTransaction) Rollback(ctx context.Context) {
	m.tx.Rollback()
}

func (m *PostgresTransaction) Error() error {
	return m.err
}

/*
func BeginDBTransaction(ctx context.Context, db *sqlx.DB) (*sqlx.Tx, error) {
	return db.Beginx()
}

func EndDBTransaction(ctx context.Context, tx *sqlx.Tx, err error) error {
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func ExecuteDBTransaction(ctx context.Context, tx *sqlx.Tx, preparedStatement *sqlx.NamedStmt, payload interface{}, err error) error {
	if err != nil {
		return err
	}
	txStmt := tx.NamedStmt(preparedStatement)
	_, err = txStmt.Exec(payload)
	return err
}*/
