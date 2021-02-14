package mockdynamodb

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Query is implemented.
// Use AddReturnQueryOutput() set QueryOutput for Query() to return.
// Query() will return QueryOutputs in the same order in which AddReturnQueryOutput() is called.
// If a nil value is given to AddReturnQueryOutput() then Query() will return an error.
// Use ReceivedQueryInputs() retrieve the inputs given to Query(), use for asserting.
func (db *DynamoDb) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	db.Lock()
	defer db.Unlock()

	return db.query(input)
}

// QueryWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) QueryWithContext(aws.Context, *dynamodb.QueryInput, ...request.Option) (*dynamodb.QueryOutput, error) {
	panic("QueryWithContext is not implemented")
}

// QueryRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) QueryRequest(*dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	panic("QueryRequest is not implemented")
}

// QueryPages is not implemented. It will panic in all cases.
func (db *DynamoDb) QueryPages(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error {
	panic("QueryPages is not implemented")
}

// QueryPagesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) QueryPagesWithContext(aws.Context, *dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool, ...request.Option) error {
	panic("QueryPagesWithContext is not implemented")
}

func (db *DynamoDb) query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	err := checkRequiredFields(map[string]interface{}{
		"QueryInput.TableName": input.TableName,
	})
	if err != nil {
		return &dynamodb.QueryOutput{}, err
	}

	if table := db.GetTable(*input.TableName); table != nil {
		table.receivedQueryInputs = append(table.receivedQueryInputs, *input)
		if len(table.returnQueryOutputs) > 0 {
			o := table.popReturnQueryOutput()
			if o != nil {
				return o, nil
			}
			return nil, errors.New("nil QueryOutput")
		}

		return nil, errors.New("no QueryOutputs to return")
	}

	return &dynamodb.QueryOutput{}, errorNonExistentTable()
}
