package mockdynamodb

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// PutItem is implemented.
// Use AddReturnPutItemOutput() set PutItemOutput for PutItem() to return.
// PutItem() will return PutItemOutputs in the same order in which AddReturnPutItemOutput() is called.
// If a nil value is given to AddReturnPutItemOutput() then PutItem() will return an error.
// Use ReceivedPutItemInputs() retrieve the inputs given to PutItem(), use for asserting.
func (db *DynamoDb) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	db.Lock()
	defer db.Unlock()

	return db.putItem(input)
}

// PutItemWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) PutItemWithContext(aws.Context, *dynamodb.PutItemInput, ...request.Option) (*dynamodb.PutItemOutput, error) {
	panic("PutItemWithContext is not implemented")
}

// PutItemRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) PutItemRequest(*dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	panic("PutItemRequest is not implemented")
}

func (db *DynamoDb) putItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	err := checkRequiredFields(map[string]interface{}{
		"PutItemInput.TableName": input.TableName,
	})
	if err != nil {
		return &dynamodb.PutItemOutput{}, err
	}

	if table := db.GetTable(*input.TableName); table != nil {
		table.receivedPutItemInputs = append(table.receivedPutItemInputs, *input)
		if len(table.returnPutItemOutputs) > 0 {
			o := table.popReturnPutItemOutput()
			if o != nil {
				return o, nil
			}
			return nil, errors.New("nil PutItemOutput")
		}

		return nil, errors.New("no PutItemOutputs to return")
	}

	return &dynamodb.PutItemOutput{}, errorNonExistentTable()
}
