package mockdynamodb

import "sync"

type DynamoDb struct {
	sync.RWMutex
	tables sync.Map
}

// New creates a new mockdynamodb client
func New() *DynamoDb {
	return &DynamoDb{}
}

// NewWithTables creates a new mockdynamodb client with tables
func NewWithTables(tables []string) *DynamoDb {
	db := New()

	for _, table := range tables {
		db.tables.Store(table, newTable())
	}

	return db
}

// GetTable returns the table with the given name. If no table exists with with the given name `nil` is returned
func (db *DynamoDb) GetTable(name string) *Table {
	t, ok := db.tables.Load(name)
	if !ok {
		return nil
	}
	return t.(*Table)
}
