package mockdynamodb_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/ebh/mockdynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDynamoDb_Query(t *testing.T) {
	t.Run("TableNameNotSet", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.Query(&dynamodb.QueryInput{})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, QueryInput.TableName.\n")
	})

	t.Run("NonExistentTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		_, err := db.Query(&dynamodb.QueryInput{TableName: &tableName})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("QueryInputRecorded", func(t *testing.T) {
		tableName := "mockTable"
		QueryInput := dynamodb.QueryInput{TableName: &tableName}
		QueryOutput := dynamodb.QueryOutput{}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnQueryOutput(&QueryOutput)

		output, err := db.Query(&QueryInput)

		assert.NoError(t, err)
		assert.Equal(t, &[]dynamodb.QueryInput{QueryInput}, db.GetTable(tableName).ReceivedQueryInputs())
		assert.Equal(t, &QueryOutput, output)
	})

	t.Run("QueryOutputIsNull", func(t *testing.T) {
		tableName := "mockTable"
		Query := dynamodb.QueryInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnQueryOutput(nil)

		_, err := db.Query(&Query)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.QueryInput{Query}, db.GetTable(tableName).ReceivedQueryInputs())
	})

	t.Run("QueryOutputEmpty", func(t *testing.T) {
		tableName := "mockTable"
		Query := dynamodb.QueryInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.Query(&Query)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.QueryInput{Query}, db.GetTable(tableName).ReceivedQueryInputs())
	})
}

func TestDynamoDb_QueryWithContext(t *testing.T) {
	t.Run("TableNameNotSet", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.QueryWithContext(context.TODO(), &dynamodb.QueryInput{})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, QueryInput.TableName.\n")
	})

	t.Run("NonExistentTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		_, err := db.QueryWithContext(context.TODO(), &dynamodb.QueryInput{TableName: &tableName})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("QueryInputRecorded", func(t *testing.T) {
		tableName := "mockTable"
		QueryInput := dynamodb.QueryInput{TableName: &tableName}
		QueryOutput := dynamodb.QueryOutput{}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnQueryOutput(&QueryOutput)

		output, err := db.QueryWithContext(context.TODO(), &QueryInput)

		assert.NoError(t, err)
		assert.Equal(t, &[]dynamodb.QueryInput{QueryInput}, db.GetTable(tableName).ReceivedQueryInputs())
		assert.Equal(t, &QueryOutput, output)
	})

	t.Run("QueryOutputIsNull", func(t *testing.T) {
		tableName := "mockTable"
		Query := dynamodb.QueryInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnQueryOutput(nil)

		_, err := db.QueryWithContext(context.TODO(), &Query)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.QueryInput{Query}, db.GetTable(tableName).ReceivedQueryInputs())
	})

	t.Run("QueryOutputEmpty", func(t *testing.T) {
		tableName := "mockTable"
		Query := dynamodb.QueryInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.QueryWithContext(context.TODO(), &Query)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.QueryInput{Query}, db.GetTable(tableName).ReceivedQueryInputs())
	})
}

func TestDynamoDb_QueryRequest(t *testing.T) {
	assert.PanicsWithValue(t, "QueryRequest is not implemented", func() {
		mockdynamodb.New().QueryRequest(nil)
	})
}

func TestDynamoDb_QueryPages(t *testing.T) {
	assert.PanicsWithValue(t, "QueryPages is not implemented", func() {
		err := mockdynamodb.New().QueryPages(nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_QueryPagesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "QueryPagesWithContext is not implemented", func() {
		err := mockdynamodb.New().QueryPagesWithContext(context.TODO(), nil, nil)
		require.NoError(t, err)
	})
}
