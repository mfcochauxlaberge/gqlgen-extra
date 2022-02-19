package store

// Driver defines the interface the implement to make a data source compatible
// with this library.
type Driver interface {
	// GetOne should populate v with data retrieved from the
	// data source using the given *Query q.
	// At most one document can be retrieved.
	GetOne(q *Query, v interface{}) error

	// GetMany should populate v with data retrieved from the
	// data source using the given *Query q.
	GetMany(q *Query, v interface{}) error

	// Put should insert a new document into the data source
	// with the given id and attributes.
	Put(id string, attrs interface{}) error

	// Update should update a document into the data source
	// with the given id and attributes.
	Update(id string, attrs interface{}) error

	// Delete should delete a document from the data source
	// with the given id.
	Delete(id string, attrs interface{}) error
}
