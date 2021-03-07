package mockdynamodb_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"

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
		putItemInput := dynamodb.PutItemInput{TableName: &tableName}
		putItemOutput := dynamodb.PutItemOutput{}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnPutItemOutput(&putItemOutput)

		output, err := db.PutItem(&putItemInput)

		assert.NoError(t, err)
		assert.Equal(t, &[]dynamodb.PutItemInput{putItemInput}, db.GetTable(tableName).ReceivedPutItemInputs())
		assert.Equal(t, &putItemOutput, output)
	})

	t.Run("PutItemOutputReturnedInOrder", func(t *testing.T) {
		tableName := "mockTable"
		putItemInput := dynamodb.PutItemInput{TableName: &tableName}

		putItemOutput1 := dynamodb.PutItemOutput{}
		putItemOutput2 := dynamodb.PutItemOutput{}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnPutItemOutput(&putItemOutput1)
		db.GetTable(tableName).AddReturnPutItemOutput(&putItemOutput2)

		output1, err1 := db.PutItem(&putItemInput)
		output2, err2 := db.PutItem(&putItemInput)

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Same(t, &putItemOutput1, output1)
		assert.Same(t, &putItemOutput2, output2)
	})

	t.Run("PutItemOutputIsNull", func(t *testing.T) {
		tableName := "mockTable"
		putItem := dynamodb.PutItemInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnPutItemOutput(nil)

		_, err := db.PutItem(&putItem)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.PutItemInput{putItem}, db.GetTable(tableName).ReceivedPutItemInputs())
	})

	t.Run("PutItemOutputEmpty", func(t *testing.T) {
		tableName := "mockTable"
		putItem := dynamodb.PutItemInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.PutItem(&putItem)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.PutItemInput{putItem}, db.GetTable(tableName).ReceivedPutItemInputs())
	})
}

func TestDynamoDb_PutItemWithContext(t *testing.T) {
	t.Run("TableNameNotSet", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.PutItemWithContext(context.TODO(), &dynamodb.PutItemInput{})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, PutItemInput.TableName.\n")
	})

	t.Run("NonExistentTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		_, err := db.PutItemWithContext(context.TODO(), &dynamodb.PutItemInput{TableName: &tableName})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("PutItemInputRecorded", func(t *testing.T) {
		tableName := "mockTable"
		putItemInput := dynamodb.PutItemInput{TableName: &tableName}
		putItemOutput := dynamodb.PutItemOutput{}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnPutItemOutput(&putItemOutput)

		output, err := db.PutItemWithContext(context.TODO(), &putItemInput)

		assert.NoError(t, err)
		assert.Equal(t, &[]dynamodb.PutItemInput{putItemInput}, db.GetTable(tableName).ReceivedPutItemInputs())
		assert.Equal(t, &putItemOutput, output)
	})

	t.Run("PutItemOutputReturnedInOrder", func(t *testing.T) {
		tableName := "mockTable"
		putItemInput := dynamodb.PutItemInput{TableName: &tableName}

		putItemOutput1 := dynamodb.PutItemOutput{}
		putItemOutput2 := dynamodb.PutItemOutput{}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnPutItemOutput(&putItemOutput1)
		db.GetTable(tableName).AddReturnPutItemOutput(&putItemOutput2)

		output1, err1 := db.PutItemWithContext(context.TODO(), &putItemInput)
		output2, err2 := db.PutItemWithContext(context.TODO(), &putItemInput)

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Same(t, &putItemOutput1, output1)
		assert.Same(t, &putItemOutput2, output2)
	})

	t.Run("PutItemOutputIsNull", func(t *testing.T) {
		tableName := "mockTable"
		putItem := dynamodb.PutItemInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnPutItemOutput(nil)

		_, err := db.PutItemWithContext(context.TODO(), &putItem)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.PutItemInput{putItem}, db.GetTable(tableName).ReceivedPutItemInputs())
	})

	t.Run("PutItemOutputEmpty", func(t *testing.T) {
		tableName := "mockTable"
		putItem := dynamodb.PutItemInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.PutItemWithContext(context.TODO(), &putItem)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.PutItemInput{putItem}, db.GetTable(tableName).ReceivedPutItemInputs())
	})
}

func TestDynamoDb_PutItemRequest(t *testing.T) {
	assert.PanicsWithValue(t, "PutItemRequest is not implemented", func() {
		mockdynamodb.New().PutItemRequest(nil)
	})
}
