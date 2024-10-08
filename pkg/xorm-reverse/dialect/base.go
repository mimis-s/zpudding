package dialect

import (
	"context"
	"fmt"
	"strings"

	"xorm.io/xorm/dialects"

	"xorm.io/xorm/core"
	"xorm.io/xorm/schemas"
)

// Base represents a basic dialect and all real dialects could embed this struct
type Base struct {
	dialect dialects.Dialect
	uri     *dialects.URI
	quoter  schemas.Quoter
}

// Alias returned col itself
func (db *Base) Alias(col string) string {
	return col
}

// Quoter returns the current database Quoter
func (db *Base) Quoter() schemas.Quoter {
	return db.quoter
}

// Init initialize the dialect
func (db *Base) Init(dialect dialects.Dialect, uri *dialects.URI) error {
	db.dialect, db.uri = dialect, uri
	return nil
}

// URI returns the uri of database
func (db *Base) URI() *dialects.URI {
	return db.uri
}

// CreateTableSQL implements Dialect
func (db *Base) CreateTableSQL(ctx context.Context, queryer core.Queryer, table *schemas.Table, tableName string) (string, bool, error) {
	if tableName == "" {
		tableName = table.Name
	}

	quoter := db.dialect.Quoter()
	var b strings.Builder
	b.WriteString("CREATE TABLE IF NOT EXISTS ")
	if err := quoter.QuoteTo(&b, tableName); err != nil {
		return "", false, err
	}
	b.WriteString(" (")

	for i, colName := range table.ColumnsSeq() {
		col := table.GetColumn(colName)
		s, _ := dialects.ColumnString(db.dialect, col, col.IsPrimaryKey && len(table.PrimaryKeys) == 1, false)
		b.WriteString(s)

		if i != len(table.ColumnsSeq())-1 {
			b.WriteString(", ")
		}
	}

	if len(table.PrimaryKeys) > 1 {
		b.WriteString(", PRIMARY KEY (")
		b.WriteString(quoter.Join(table.PrimaryKeys, ","))
		b.WriteString(")")
	}

	b.WriteString(")")

	return b.String(), false, nil
}

func (db *Base) CreateSequenceSQL(ctx context.Context, queryer core.Queryer, seqName string) (string, error) {
	return fmt.Sprintf(`CREATE SEQUENCE %s 
	minvalue 1
	   nomaxvalue
	   start with 1
	   increment by 1
	   nocycle
	nocache`, seqName), nil
}

func (db *Base) IsSequenceExist(ctx context.Context, queryer core.Queryer, seqName string) (bool, error) {
	return false, fmt.Errorf("unsupported sequence feature")
}

func (db *Base) DropSequenceSQL(seqName string) (string, error) {
	return fmt.Sprintf("DROP SEQUENCE %s", seqName), nil
}

// DropTableSQL returns drop table SQL
func (db *Base) DropTableSQL(tableName string) (string, bool) {
	quote := db.dialect.Quoter().Quote
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", quote(tableName)), true
}

// HasRecords returns true if the SQL has records returned
func (db *Base) HasRecords(queryer core.Queryer, ctx context.Context, query string, args ...interface{}) (bool, error) {
	rows, err := queryer.QueryContext(ctx, query, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, rows.Err()
}

// IsColumnExist returns true if the column of the table exist
func (db *Base) IsColumnExist(queryer core.Queryer, ctx context.Context, tableName, colName string) (bool, error) {
	quote := db.dialect.Quoter().Quote
	query := fmt.Sprintf(
		"SELECT %v FROM %v.%v WHERE %v = ? AND %v = ? AND %v = ?",
		quote("COLUMN_NAME"),
		quote("INFORMATION_SCHEMA"),
		quote("COLUMNS"),
		quote("TABLE_SCHEMA"),
		quote("TABLE_NAME"),
		quote("COLUMN_NAME"),
	)
	return db.HasRecords(queryer, ctx, query, db.uri.DBName, tableName, colName)
}

// AddColumnSQL returns a SQL to add a column
func (db *Base) AddColumnSQL(tableName string, col *schemas.Column) string {
	s, _ := dialects.ColumnString(db.dialect, col, true, false)
	return fmt.Sprintf("ALTER TABLE %s ADD %s", db.dialect.Quoter().Quote(tableName), s)
}

// CreateIndexSQL returns a SQL to create index
func (db *Base) CreateIndexSQL(tableName string, index *schemas.Index) string {
	quoter := db.dialect.Quoter()
	var unique string
	var idxName string
	if index.Type == schemas.UniqueType {
		unique = " UNIQUE"
	}
	idxName = index.XName(tableName)
	return fmt.Sprintf("CREATE%s INDEX %v ON %v (%v)", unique,
		quoter.Quote(idxName), quoter.Quote(tableName),
		quoter.Join(index.Cols, ","))
}

// DropIndexSQL returns a SQL to drop index
func (db *Base) DropIndexSQL(tableName string, index *schemas.Index) string {
	quote := db.dialect.Quoter().Quote
	var name string
	if index.IsRegular {
		name = index.XName(tableName)
	} else {
		name = index.Name
	}
	return fmt.Sprintf("DROP INDEX %v ON %s", quote(name), quote(tableName))
}

// ModifyColumnSQL returns a SQL to modify SQL
func (db *Base) ModifyColumnSQL(tableName string, col *schemas.Column) string {
	s, _ := dialects.ColumnString(db.dialect, col, false, false)
	return fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s", db.quoter.Quote(tableName), s)
}

// SetParams set params
func (db *Base) SetParams(params map[string]string) {
}
