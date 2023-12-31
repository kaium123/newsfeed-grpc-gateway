// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldParentCommentID holds the string denoting the parent_comment_id field in the database.
	FieldParentCommentID = "parent_comment_id"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeReacts holds the string denoting the reacts edge name in mutations.
	EdgeReacts = "reacts"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// AttachmentsTable is the table that holds the attachments relation/edge. The primary key declared below.
	AttachmentsTable = "comment_attachments"
	// AttachmentsInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentsInverseTable = "attachments"
	// PostTable is the table that holds the post relation/edge. The primary key declared below.
	PostTable = "post_comments"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// ReactsTable is the table that holds the reacts relation/edge. The primary key declared below.
	ReactsTable = "comment_reacts"
	// ReactsInverseTable is the table name for the React entity.
	// It exists in this package in order to avoid circular dependency with the "react" package.
	ReactsInverseTable = "reacts"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldContent,
	FieldParentCommentID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_comments",
}

var (
	// AttachmentsPrimaryKey and AttachmentsColumn2 are the table columns denoting the
	// primary key for the attachments relation (M2M).
	AttachmentsPrimaryKey = []string{"comment_id", "attachment_id"}
	// PostPrimaryKey and PostColumn2 are the table columns denoting the
	// primary key for the post relation (M2M).
	PostPrimaryKey = []string{"post_id", "comment_id"}
	// ReactsPrimaryKey and ReactsColumn2 are the table columns denoting the
	// primary key for the reacts relation (M2M).
	ReactsPrimaryKey = []string{"comment_id", "react_id"}
)

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
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultParentCommentID holds the default value on creation for the "parent_comment_id" field.
	DefaultParentCommentID int
)

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByParentCommentID orders the results by the parent_comment_id field.
func ByParentCommentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentCommentID, opts...).ToFunc()
}

// ByAttachmentsCount orders the results by attachments count.
func ByAttachmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttachmentsStep(), opts...)
	}
}

// ByAttachments orders the results by attachments terms.
func ByAttachments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttachmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostCount orders the results by post count.
func ByPostCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostStep(), opts...)
	}
}

// ByPost orders the results by post terms.
func ByPost(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReactsCount orders the results by reacts count.
func ByReactsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReactsStep(), opts...)
	}
}

// ByReacts orders the results by reacts terms.
func ByReacts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReactsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAttachmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttachmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AttachmentsTable, AttachmentsPrimaryKey...),
	)
}
func newPostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PostTable, PostPrimaryKey...),
	)
}
func newReactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReactsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ReactsTable, ReactsPrimaryKey...),
	)
}
