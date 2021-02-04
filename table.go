package mockdynamodb

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Table struct {
	sync.RWMutex
	items []dynamodb.PutItemInput
}

func newTable() *Table {
	return &Table{
		items: nil,
	}
}

// Items returns the PutItemInput submitted by all calls to PutItem()
func (t *Table) Items() *[]dynamodb.PutItemInput {
	t.Lock()
	defer t.Unlock()

	tmp := make([]dynamodb.PutItemInput, len(t.items))
	copy(tmp, t.items)
	return &tmp
}
