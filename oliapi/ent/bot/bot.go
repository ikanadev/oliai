// Code generated by ent, DO NOT EDIT.

package bot

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the bot type in the database.
	Label = "bot"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGreetingMessage holds the string denoting the greeting_message field in the database.
	FieldGreetingMessage = "greeting_message"
	// FieldCustomPropmt holds the string denoting the custom_propmt field in the database.
	FieldCustomPropmt = "custom_propmt"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldArchivedAt holds the string denoting the archived_at field in the database.
	FieldArchivedAt = "archived_at"
	// EdgeBlocks holds the string denoting the blocks edge name in mutations.
	EdgeBlocks = "blocks"
	// EdgeCompany holds the string denoting the company edge name in mutations.
	EdgeCompany = "company"
	// Table holds the table name of the bot in the database.
	Table = "bots"
	// BlocksTable is the table that holds the blocks relation/edge.
	BlocksTable = "block_categories"
	// BlocksInverseTable is the table name for the BlockCategory entity.
	// It exists in this package in order to avoid circular dependency with the "blockcategory" package.
	BlocksInverseTable = "block_categories"
	// BlocksColumn is the table column denoting the blocks relation/edge.
	BlocksColumn = "bot_blocks"
	// CompanyTable is the table that holds the company relation/edge.
	CompanyTable = "bots"
	// CompanyInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompanyInverseTable = "companies"
	// CompanyColumn is the table column denoting the company relation/edge.
	CompanyColumn = "company_bots"
)

// Columns holds all SQL columns for bot fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldGreetingMessage,
	FieldCustomPropmt,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldArchivedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "bots"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"company_bots",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Bot queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByGreetingMessage orders the results by the greeting_message field.
func ByGreetingMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGreetingMessage, opts...).ToFunc()
}

// ByCustomPropmt orders the results by the custom_propmt field.
func ByCustomPropmt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomPropmt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByArchivedAt orders the results by the archived_at field.
func ByArchivedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArchivedAt, opts...).ToFunc()
}

// ByBlocksCount orders the results by blocks count.
func ByBlocksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBlocksStep(), opts...)
	}
}

// ByBlocks orders the results by blocks terms.
func ByBlocks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlocksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCompanyField orders the results by company field.
func ByCompanyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCompanyStep(), sql.OrderByField(field, opts...))
	}
}
func newBlocksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlocksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BlocksTable, BlocksColumn),
	)
}
func newCompanyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CompanyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
	)
}
