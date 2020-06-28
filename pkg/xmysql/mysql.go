package xmysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	errors "git.gmtshenzhen.com/geeky-medic/x-lite/xerrors"
)

type DB struct {
	db *sql.DB
}

// Open a new connect
func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, errors.By(err)
	}
	return &DB{db: db}, nil
}

func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.db.SetConnMaxLifetime(d)
}

func (db *DB) SetMaxIdleConns(n int) {
	db.db.SetMaxIdleConns(n)
}

func (db *DB) SetMaxOpenConns(n int) {
	db.db.SetMaxOpenConns(n)
}

func (db *DB) PingContext(ctx context.Context) error {
	err := db.db.PingContext(ctx)
	return errors.By(err)
}

func (db *DB) Close() error {
	return errors.By(db.db.Close())
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := db.db.ExecContext(ctx, query, args...)
	return result, errors.By(err)
}

func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error) {
	result, err := db.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.By(err)
	}
	return &Stmt{stmt: result}, nil
}

type Stmt struct {
	stmt *sql.Stmt
}

func (stmt *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	result, err := stmt.stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, errors.By(err)
	}
	return result, nil
}

func (stmt *Stmt) QueryContext(ctx context.Context, args ...interface{}) (*Rows, error) {
	rows, err := stmt.stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.By(err)
	}
	return &Rows{rows: rows}, nil
}

func (stmt *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) *Row {
	row := stmt.stmt.QueryRowContext(ctx, args...)
	return &Row{row: row}
}

func (stmt *Stmt) Close() error {
	return errors.By(stmt.stmt.Close())
}

type Rows struct {
	rows *sql.Rows
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	if args == nil {
		args = []interface{}{}
	}
	rows, err := db.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.By(err)
	}
	return &Rows{rows: rows}, nil
}

func (rows *Rows) Scan(dest ...interface{}) error {
	err := rows.rows.Scan(dest...)
	return errors.By(err)
}

func (rows *Rows) Next() bool {
	ok := rows.rows.Next()
	return ok
}

func (rows *Rows) NextResultSet() bool {
	ok := rows.rows.NextResultSet()
	return ok
}

func (rows *Rows) Columns() ([]string, error) {
	strs, err := rows.rows.Columns()
	return strs, errors.By(err)
}

func (rows *Rows) Close() error {
	return errors.By(rows.rows.Close())
}

func (rows *Rows) Err() error {
	return errors.By(rows.rows.Err())
}

type ColumnType struct {
	ct *sql.ColumnType
}

func (rows *Rows) ColumnTypes() ([]*ColumnType, error) {
	col, err := rows.rows.ColumnTypes()
	if err != nil {
		return nil, errors.By(err)
	}
	var result []*ColumnType
	for _, c := range col {
		result = append(result, &ColumnType{c})
	}
	return result, nil
}

type Row struct {
	row *sql.Row
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row {
	if len(args) == 0 {
		args = []interface{}{}
	}
	row := db.db.QueryRowContext(ctx, query, args...)
	return &Row{row: row}
}

func (r *Row) Scan(dest ...interface{}) error {
	err := r.row.Scan(dest...)
	if err != nil && err == sql.ErrNoRows {
		return err
	}
	return errors.By(err)
}

type Tx struct {
	tx *sql.Tx
}

//open transaction
func (db *DB) Begin() (*Tx, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return nil, errors.By(err)
	}
	return &Tx{tx}, nil
}

type TxOptions = sql.TxOptions

func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error) {
	tx, err := db.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, errors.By(err)
	}
	return &Tx{tx}, nil
}

//commit transaction
func (tx *Tx) Commit() error {
	err := tx.tx.Commit()
	return errors.By(err)
}

func (tx *Tx) Rollback() error {
	err := tx.tx.Rollback()
	return errors.By(err)
}

func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error) {
	stmt, err := tx.tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.By(err)
	}
	return &Stmt{stmt}, nil
}

func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt {
	st := tx.tx.StmtContext(ctx, stmt.stmt)
	return &Stmt{st}
}

func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if args == nil {
		args = []interface{}{}
	}
	result, err := tx.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, errors.By(err)
	}
	return result, nil
}

func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	if args == nil {
		args = []interface{}{}
	}
	rows, err := tx.tx.QueryContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		err = errors.By(err)
	}
	return &Rows{rows}, err
}

func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row {
	if args == nil {
		args = []interface{}{}
	}
	row := tx.tx.QueryRowContext(ctx, query, args...)
	return &Row{row}
}

type Conn struct {
	c *sql.Conn
}

func (db *DB) Conn(ctx context.Context) (*Conn, error) {
	conn, err := db.db.Conn(ctx)
	if err != nil && err != sql.ErrConnDone {
		err = errors.By(err)
	}
	return &Conn{conn}, err
}

func (c *Conn) PingContext(ctx context.Context) error {
	err := c.c.PingContext(ctx)
	return errors.By(err)
}

func (c *Conn) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if args == nil {
		args = []interface{}{}
	}
	result, err := c.c.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, errors.By(err)
	}
	return result, nil
}

func (c *Conn) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	if args == nil {
		args = []interface{}{}
	}
	rows, err := c.c.QueryContext(ctx, query, args...)

	if err != nil && err != sql.ErrNoRows {
		err = errors.By(err)
	}
	return &Rows{rows}, err
}

func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row {
	if args == nil {
		args = []interface{}{}
	}
	row := c.c.QueryRowContext(ctx, query, args...)
	return &Row{row}
}

func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error) {
	st, err := c.c.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.By(err)
	}
	return &Stmt{st}, nil
}

func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error) {
	tx, err := c.c.BeginTx(ctx, opts)
	if err != nil {
		return nil, errors.By(err)
	}
	return &Tx{tx}, nil
}

func (c *Conn) Close() error {
	err := c.c.Close()
	return errors.By(err)
}

func (db *DB) Stats() sql.DBStats {
	dbs := db.db.Stats()
	return dbs
}

func (db *DB) Driver() driver.Driver {
	driver := db.db.Driver()
	return driver
}

func Named(name string, value interface{}) sql.NamedArg {
	nd := sql.Named(name, value)
	return nd
}

type IsolationLevel sql.IsolationLevel

/*
//注册
func Register(name string, driver driver.Driver) {
	sql.Register(name, driver)
}
func Drivers() []string {
	list := sql.Drivers()
	return list
}
type NullString struct {
	ns *sql.NullString
}
func (ns *NullString) Scan(value interface{}) error {
	err := ns.ns.Scan(value)
	if err != nil {
		//
	}
	return err
}
type NullInt64 struct {
	n *sql.NullInt64
}
func (n *NullInt64) Scan(value interface{}) error {
	err := n.n.Scan(value)
	if err != nil {
		//
	}
	return err
}
type NullFloat64 struct {
	n *sql.NullFloat64
}
func (n *NullFloat64) Scan(value interface{}) error {
	err := n.n.Scan(value)
	if err != nil {
		//
	}
	return err
}
type NullBool struct {
	n *sql.NullBool
}
func (n *NullBool) Scan(value interface{}) error {
	err := n.n.Scan(value)
	if err != nil {
		//
	}
	return err
}
*/

func IterRowsFn(fn func() error) error {
	return fn()
}