package mockdynamodb_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/stretchr/testify/require"

	"github.com/ebh/mockdynamodb"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDb_PutItem(t *testing.T) {
	t.Run("TableNameNotSet", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.PutItem(&dynamodb.PutItemInput{})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, PutItemInput.TableName.\n")
	})

	t.Run("NonExistentTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		_, err := db.PutItem(&dynamodb.PutItemInput{TableName: &tableName})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("PutItemInputRecorded", func(t *testing.T) {
		tableName := "mockTable"
		putItem := dynamodb.PutItemInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.PutItem(&putItem)

		require.NoError(t, err)
		assert.Equal(t, &[]dynamodb.PutItemInput{putItem}, db.GetTable(tableName).Items())
	})
}

func TestDynamoDb_PutItemWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "PutItemWithContext is not implemented", func() {
		_, err := mockdynamodb.New().PutItemWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_PutItemRequest(t *testing.T) {
	assert.PanicsWithValue(t, "PutItemRequest is not implemented", func() {
		mockdynamodb.New().PutItemRequest(nil)
	})
}
