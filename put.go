package mockdynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// PutItem is implemented.
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
		table.items = append(table.items, *input)
		return &dynamodb.PutItemOutput{}, nil
	}

	return &dynamodb.PutItemOutput{}, errorNonExistentTable()
}
