package mockdynamodb

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Table struct {
	sync.RWMutex

	receivedPutItemInputs []dynamodb.PutItemInput
	returnPutItemOutputs  []*dynamodb.PutItemOutput

	receivedQueryInputs []dynamodb.QueryInput
	returnQueryOutputs  []*dynamodb.QueryOutput
}

func newTable() *Table {
	return &Table{
		returnPutItemOutputs: []*dynamodb.PutItemOutput{},
		returnQueryOutputs:   []*dynamodb.QueryOutput{},
	}
}

// ReceivedPutItemInputs returns the PutItemInputs submitted by all calls to PutItem()
// The order of the elements returned is the order in which they were received by PutItem()
func (t *Table) ReceivedPutItemInputs() *[]dynamodb.PutItemInput {
	t.Lock()
	defer t.Unlock()

	tmp := make([]dynamodb.PutItemInput, len(t.receivedPutItemInputs))
	copy(tmp, t.receivedPutItemInputs)
	return &tmp
}

// AddReturnPutItemOutput pushes a PutItemOutput that will then be returned by calls to PutItem()
// PutItemOutputs are returned in the same order in which they are pushed
// If nil is pushed then PutItem() will return an error
func (t *Table) AddReturnPutItemOutput(outputs ...*dynamodb.PutItemOutput) {
	t.Lock()
	defer t.Unlock()

	t.returnPutItemOutputs = append(t.returnPutItemOutputs, outputs...)
}

func (t *Table) popReturnPutItemOutput() *dynamodb.PutItemOutput {
	x, a := t.returnPutItemOutputs[0], t.returnPutItemOutputs[1:]
	t.returnPutItemOutputs = a
	return x
}

// ReceivedQueryInputs returns the QueryInputs submitted by all calls to Query()
// The order of the elements returned is the order in which they were received by Query()
func (t *Table) ReceivedQueryInputs() *[]dynamodb.QueryInput {
	t.Lock()
	defer t.Unlock()

	tmp := make([]dynamodb.QueryInput, len(t.receivedQueryInputs))
	copy(tmp, t.receivedQueryInputs)
	return &tmp
}

// AddReturnQueryOutput pushes a QueryOutput that will then be returned by calls to Query()
// QueryOutputs are returned in the same order in which they are pushed
// If nil is pushed then Query() will return an error
func (t *Table) AddReturnQueryOutput(outputs ...*dynamodb.QueryOutput) {
	t.Lock()
	defer t.Unlock()

	t.returnQueryOutputs = append(t.returnQueryOutputs, outputs...)
}

func (t *Table) popReturnQueryOutput() *dynamodb.QueryOutput {
	x, a := t.returnQueryOutputs[0], t.returnQueryOutputs[1:]
	t.returnQueryOutputs = a
	return x
}

func (t *Table) moreQueryOutputs() bool {
	return len(t.returnQueryOutputs) > 0
}
