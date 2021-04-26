package mockdynamodb_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/ebh/mockdynamodb"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDb_Scan(t *testing.T) {
	t.Run("TableNameNotSet", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.Scan(&dynamodb.ScanInput{})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, ScanInput.TableName.\n")
	})

	t.Run("NonExistentTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		_, err := db.Scan(&dynamodb.ScanInput{TableName: &tableName})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("ScanInputRecorded", func(t *testing.T) {
		tableName := "mockTable"
		ScanInput := dynamodb.ScanInput{TableName: &tableName}
		ScanOutput := dynamodb.ScanOutput{}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(&ScanOutput)

		output, err := db.Scan(&ScanInput)

		assert.NoError(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{ScanInput}, db.GetTable(tableName).ReceivedScanInputs())
		assert.Equal(t, &ScanOutput, output)
	})

	t.Run("ScanOutputIsNull", func(t *testing.T) {
		tableName := "mockTable"
		Scan := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(nil)

		_, err := db.Scan(&Scan)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{Scan}, db.GetTable(tableName).ReceivedScanInputs())
	})

	t.Run("ScanOutputEmpty", func(t *testing.T) {
		tableName := "mockTable"
		Scan := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.Scan(&Scan)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{Scan}, db.GetTable(tableName).ReceivedScanInputs())
	})
}

func TestDynamoDb_ScanWithContext(t *testing.T) {
	t.Run("TableNameNotSet", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.ScanWithContext(context.TODO(), &dynamodb.ScanInput{})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, ScanInput.TableName.\n")
	})

	t.Run("NonExistentTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		_, err := db.ScanWithContext(context.TODO(), &dynamodb.ScanInput{TableName: &tableName})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("ScanInputRecorded", func(t *testing.T) {
		tableName := "mockTable"
		ScanInput := dynamodb.ScanInput{TableName: &tableName}
		ScanOutput := dynamodb.ScanOutput{}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(&ScanOutput)

		output, err := db.ScanWithContext(context.TODO(), &ScanInput)

		assert.NoError(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{ScanInput}, db.GetTable(tableName).ReceivedScanInputs())
		assert.Equal(t, &ScanOutput, output)
	})

	t.Run("ScanOutputIsNull", func(t *testing.T) {
		tableName := "mockTable"
		Scan := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(nil)

		_, err := db.ScanWithContext(context.TODO(), &Scan)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{Scan}, db.GetTable(tableName).ReceivedScanInputs())
	})

	t.Run("ScanOutputEmpty", func(t *testing.T) {
		tableName := "mockTable"
		Scan := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.ScanWithContext(context.TODO(), &Scan)

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{Scan}, db.GetTable(tableName).ReceivedScanInputs())
	})
}

func TestDynamoDb_ScanRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ScanRequest is not implemented", func() {
		mockdynamodb.New().ScanRequest(nil)
	})
}

func TestDynamoDb_ScanPages(t *testing.T) {
	t.Run("TableNameNotSet", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		err := db.ScanPages(&dynamodb.ScanInput{}, func(output *dynamodb.ScanOutput, b bool) bool {
			return true
		})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, ScanInput.TableName.\n")
	})

	t.Run("NonExistentTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		err := db.ScanPages(&dynamodb.ScanInput{TableName: &tableName}, func(output *dynamodb.ScanOutput, b bool) bool {
			return true
		})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("ScanInputRecorded", func(t *testing.T) {
		tableName := "mockTable"
		ScanInput := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(scanOutputs(3)...)

		var outputs []dynamodb.ScanOutput

		err := db.ScanPages(&ScanInput, func(output *dynamodb.ScanOutput, b bool) bool {
			outputs = append(outputs, *output)
			return true
		})

		assert.NoError(t, err, "error")
		assert.Equal(t, &[]dynamodb.ScanInput{ScanInput}, db.GetTable(tableName).ReceivedScanInputs(), "inputs")
		assert.Equal(t, 3, len(outputs), "output count")
	})

	t.Run("ScanOutputIsNull", func(t *testing.T) {
		tableName := "mockTable"
		Scan := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(nil)

		var outputs []dynamodb.ScanOutput

		err := db.ScanPages(&Scan, func(output *dynamodb.ScanOutput, b bool) bool {
			outputs = append(outputs, *output)
			return true
		})

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{Scan}, db.GetTable(tableName).ReceivedScanInputs())
	})

	t.Run("ScanOutputEmpty", func(t *testing.T) {
		tableName := "mockTable"
		Scan := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		var outputs []dynamodb.ScanOutput

		err := db.ScanPages(&Scan, func(output *dynamodb.ScanOutput, b bool) bool {
			outputs = append(outputs, *output)
			return true
		})

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{Scan}, db.GetTable(tableName).ReceivedScanInputs())
	})

	t.Run("Output handler func returns false", func(t *testing.T) {
		tableName := "mockTable"
		ScanInput := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(scanOutputs(3)...)

		var outputs []dynamodb.ScanOutput

		err := db.ScanPages(&ScanInput, func(output *dynamodb.ScanOutput, b bool) bool {
			outputs = append(outputs, *output)
			return len(outputs) < 2
		})

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{ScanInput}, db.GetTable(tableName).ReceivedScanInputs(), "inputs")
		assert.Equal(t, 2, len(outputs), "output count")
	})
}

func TestDynamoDb_ScanPagesWithContext(t *testing.T) {
	t.Run("TableNameNotSet", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		err := db.ScanPagesWithContext(context.TODO(), &dynamodb.ScanInput{}, func(output *dynamodb.ScanOutput, b bool) bool {
			return true
		})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, ScanInput.TableName.\n")
	})

	t.Run("NonExistentTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		err := db.ScanPagesWithContext(context.TODO(), &dynamodb.ScanInput{TableName: &tableName}, func(output *dynamodb.ScanOutput, b bool) bool {
			return true
		})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("ScanInputRecorded", func(t *testing.T) {
		tableName := "mockTable"
		ScanInput := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(scanOutputs(3)...)

		var outputs []dynamodb.ScanOutput

		err := db.ScanPagesWithContext(context.TODO(), &ScanInput, func(output *dynamodb.ScanOutput, b bool) bool {
			outputs = append(outputs, *output)
			return true
		})

		assert.NoError(t, err, "error")
		assert.Equal(t, &[]dynamodb.ScanInput{ScanInput}, db.GetTable(tableName).ReceivedScanInputs(), "inputs")
		assert.Equal(t, 3, len(outputs), "output count")
	})

	t.Run("ScanOutputIsNull", func(t *testing.T) {
		tableName := "mockTable"
		Scan := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(nil)

		var outputs []dynamodb.ScanOutput

		err := db.ScanPagesWithContext(context.TODO(), &Scan, func(output *dynamodb.ScanOutput, b bool) bool {
			outputs = append(outputs, *output)
			return true
		})

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{Scan}, db.GetTable(tableName).ReceivedScanInputs())
	})

	t.Run("ScanOutputEmpty", func(t *testing.T) {
		tableName := "mockTable"
		Scan := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})

		var outputs []dynamodb.ScanOutput

		err := db.ScanPagesWithContext(context.TODO(), &Scan, func(output *dynamodb.ScanOutput, b bool) bool {
			outputs = append(outputs, *output)
			return true
		})

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{Scan}, db.GetTable(tableName).ReceivedScanInputs())
	})

	t.Run("Output handler func returns false", func(t *testing.T) {
		tableName := "mockTable"
		ScanInput := dynamodb.ScanInput{TableName: &tableName}

		db := mockdynamodb.NewWithTables([]string{tableName})
		db.GetTable(tableName).AddReturnScanOutput(scanOutputs(3)...)

		var outputs []dynamodb.ScanOutput

		err := db.ScanPagesWithContext(context.TODO(), &ScanInput, func(output *dynamodb.ScanOutput, b bool) bool {
			outputs = append(outputs, *output)
			return len(outputs) < 2
		})

		assert.Error(t, err)
		assert.Equal(t, &[]dynamodb.ScanInput{ScanInput}, db.GetTable(tableName).ReceivedScanInputs(), "inputs")
		assert.Equal(t, 2, len(outputs), "output count")
	})
}

func scanOutputs(n int) []*dynamodb.ScanOutput {
	var a []*dynamodb.ScanOutput
	for i := 0; i < n; i++ {
		a = append(a, &dynamodb.ScanOutput{})
	}
	return a
}
