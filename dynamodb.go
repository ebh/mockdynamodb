package mockdynamodb

import "sync"

type DynamoDb struct {
	sync.RWMutex
	tables sync.Map
}

func New() *DynamoDb {
	return &DynamoDb{}
}

func NewWithTables(tables []string) *DynamoDb {
	db := New()

	for _, table := range tables {
		db.tables.Store(table, newTable())
	}

	return db
}

func (db *DynamoDb) GetTable(name string) *Table {
	t, ok := db.tables.Load(name)
	if !ok {
		return nil
	}
	return t.(*Table)
}
