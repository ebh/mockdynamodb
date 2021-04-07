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

// QueryWithContext not implemented.
// Uses same logic as Query() and thus does not record context or options
func (db *DynamoDb) QueryWithContext(_ aws.Context, input *dynamodb.QueryInput, _ ...request.Option) (*dynamodb.QueryOutput, error) {
	db.Lock()
	defer db.Unlock()

	return db.query(input)
}

// QueryRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) QueryRequest(*dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	panic("QueryRequest is not implemented")
}

// QueryPages is not implemented. It will panic in all cases.
func (db *DynamoDb) QueryPages(i *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error {
	db.Lock()
	defer db.Unlock()

	return db.queryPages(i, fn)
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
		if table.moreQueryOutputs() {
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

func (db *DynamoDb) queryPages(input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error {
	err := checkRequiredFields(map[string]interface{}{
		"QueryInput.TableName": input.TableName,
	})
	if err != nil {
		return err
	}

	if table := db.GetTable(*input.TableName); table != nil {
		table.receivedQueryInputs = append(table.receivedQueryInputs, *input)
		if table.moreQueryOutputs() {
			for ok := true; ok; ok = table.moreQueryOutputs() {
				output := table.popReturnQueryOutput()

				if output == nil {
					return errors.New("nil QueryOutput")
				}

				if !fn(output, false) {
					return errors.New("output handler function returned false")
				}
			}

			return nil
		}

		return errors.New("no QueryOutputs to return")
	}

	return errorNonExistentTable()
}
