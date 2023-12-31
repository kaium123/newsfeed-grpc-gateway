// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeReacts holds the string denoting the reacts edge name in mutations.
	EdgeReacts = "reacts"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// AttachmentsTable is the table that holds the attachments relation/edge. The primary key declared below.
	AttachmentsTable = "post_attachments"
	// AttachmentsInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentsInverseTable = "attachments"
	// CommentsTable is the table that holds the comments relation/edge. The primary key declared below.
	CommentsTable = "post_comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// ReactsTable is the table that holds the reacts relation/edge. The primary key declared below.
	ReactsTable = "post_reacts"
	// ReactsInverseTable is the table name for the React entity.
	// It exists in this package in order to avoid circular dependency with the "react" package.
	ReactsInverseTable = "reacts"
	// AuthorTable is the table that holds the author relation/edge. The primary key declared below.
	AuthorTable = "user_posts"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldContent,
}

var (
	// AttachmentsPrimaryKey and AttachmentsColumn2 are the table columns denoting the
	// primary key for the attachments relation (M2M).
	AttachmentsPrimaryKey = []string{"post_id", "attachment_id"}
	// CommentsPrimaryKey and CommentsColumn2 are the table columns denoting the
	// primary key for the comments relation (M2M).
	CommentsPrimaryKey = []string{"post_id", "comment_id"}
	// ReactsPrimaryKey and ReactsColumn2 are the table columns denoting the
	// primary key for the reacts relation (M2M).
	ReactsPrimaryKey = []string{"post_id", "react_id"}
	// AuthorPrimaryKey and AuthorColumn2 are the table columns denoting the
	// primary key for the author relation (M2M).
	AuthorPrimaryKey = []string{"user_id", "post_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
)

// OrderOption defines the ordering options for the Post queries.
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

// ByCommentsCount orders the results by comments count.
func ByCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentsStep(), opts...)
	}
}

// ByComments orders the results by comments terms.
func ByComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByAuthorCount orders the results by author count.
func ByAuthorCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorStep(), opts...)
	}
}

// ByAuthor orders the results by author terms.
func ByAuthor(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAttachmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttachmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AttachmentsTable, AttachmentsPrimaryKey...),
	)
}
func newCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CommentsTable, CommentsPrimaryKey...),
	)
}
func newReactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReactsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ReactsTable, ReactsPrimaryKey...),
	)
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AuthorTable, AuthorPrimaryKey...),
	)
}
