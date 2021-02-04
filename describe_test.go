package mockdynamodb_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/stretchr/testify/require"

	"github.com/ebh/mockdynamodb"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDb_DescribeBackup(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeBackup is not implemented", func() {
		_, err := mockdynamodb.New().DescribeBackup(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeBackupWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeBackupWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeBackupWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeBackupRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeBackupRequest is not implemented", func() {
		mockdynamodb.New().DescribeBackupRequest(nil)
	})
}

func TestDynamoDb_DescribeContinuousBackups(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeContinuousBackups is not implemented", func() {
		_, err := mockdynamodb.New().DescribeContinuousBackups(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeContinuousBackupsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeContinuousBackupsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeContinuousBackupsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeContinuousBackupsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeContinuousBackupsRequest is not implemented", func() {
		mockdynamodb.New().DescribeContinuousBackupsRequest(nil)
	})
}

func TestDynamoDb_DescribeContributorInsights(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeContributorInsights is not implemented", func() {
		_, err := mockdynamodb.New().DescribeContributorInsights(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeContributorInsightsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeContributorInsightsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeContributorInsightsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeContributorInsightsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeContributorInsightsRequest is not implemented", func() {
		mockdynamodb.New().DescribeContributorInsightsRequest(nil)
	})
}

func TestDynamoDb_DescribeEndpoints(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeEndpoints is not implemented", func() {
		_, err := mockdynamodb.New().DescribeEndpoints(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeEndpointsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeEndpointsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeEndpointsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeEndpointsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeEndpointsRequest is not implemented", func() {
		mockdynamodb.New().DescribeEndpointsRequest(nil)
	})
}

func TestDynamoDb_DescribeExport(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeExport is not implemented", func() {
		_, err := mockdynamodb.New().DescribeExport(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeExportWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeExportWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeExportWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeExportRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeExportRequest is not implemented", func() {
		mockdynamodb.New().DescribeExportRequest(nil)
	})
}

func TestDynamoDb_DescribeGlobalTable(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeGlobalTable is not implemented", func() {
		_, err := mockdynamodb.New().DescribeGlobalTable(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeGlobalTableWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeGlobalTableWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeGlobalTableWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeGlobalTableRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeGlobalTableRequest is not implemented", func() {
		mockdynamodb.New().DescribeGlobalTableRequest(nil)
	})
}

func TestDynamoDb_DescribeGlobalTableSettings(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeGlobalTableSettings is not implemented", func() {
		_, err := mockdynamodb.New().DescribeGlobalTableSettings(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeGlobalTableSettingsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeGlobalTableSettingsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeGlobalTableSettingsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeGlobalTableSettingsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeGlobalTableSettingsRequest is not implemented", func() {
		mockdynamodb.New().DescribeGlobalTableSettingsRequest(nil)
	})
}

func TestDynamoDb_DescribeKinesisStreamingDestination(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeKinesisStreamingDestination is not implemented", func() {
		_, err := mockdynamodb.New().DescribeKinesisStreamingDestination(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeKinesisStreamingDestinationWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeKinesisStreamingDestinationWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeKinesisStreamingDestinationWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeKinesisStreamingDestinationRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeKinesisStreamingDestinationRequest is not implemented", func() {
		mockdynamodb.New().DescribeKinesisStreamingDestinationRequest(nil)
	})
}

func TestDynamoDb_DescribeLimits(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeLimits is not implemented", func() {
		_, err := mockdynamodb.New().DescribeLimits(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeLimitsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeLimitsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeLimitsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeLimitsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeLimitsRequest is not implemented", func() {
		mockdynamodb.New().DescribeLimitsRequest(nil)
	})
}

func TestDynamoDb_DescribeTable(t *testing.T) {
	t.Run("ReturnsTable", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		table, err := db.DescribeTable(&dynamodb.DescribeTableInput{TableName: &tableName})

		require.NoError(t, err)
		assert.Equal(t, tableName, *table.Table.TableName)
	})

	t.Run("TableDoesNotExist", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{"anotherMockTable"})

		_, err := db.DescribeTable(&dynamodb.DescribeTableInput{TableName: &tableName})

		assertRegexpError(t, err, "AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.\\n\\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("TableNameNotSpecified", func(t *testing.T) {
		tableName := "mockTable"
		db := mockdynamodb.NewWithTables([]string{tableName})

		_, err := db.DescribeTable(&dynamodb.DescribeTableInput{})

		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, DescribeTableInput.TableName.\n")
	})
}

func TestDynamoDb_DescribeTableWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeTableWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeTableWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeTableRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeTableRequest is not implemented", func() {
		mockdynamodb.New().DescribeTableRequest(nil)
	})
}

func TestDynamoDb_DescribeTableReplicaAutoScaling(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeTableReplicaAutoScaling is not implemented", func() {
		_, err := mockdynamodb.New().DescribeTableReplicaAutoScaling(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeTableReplicaAutoScalingWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeTableReplicaAutoScalingWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeTableReplicaAutoScalingWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeTableReplicaAutoScalingRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeTableReplicaAutoScalingRequest is not implemented", func() {
		mockdynamodb.New().DescribeTableReplicaAutoScalingRequest(nil)
	})
}

func TestDynamoDb_DescribeTimeToLive(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeTimeToLive is not implemented", func() {
		_, err := mockdynamodb.New().DescribeTimeToLive(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeTimeToLiveWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeTimeToLiveWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DescribeTimeToLiveWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DescribeTimeToLiveRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DescribeTimeToLiveRequest is not implemented", func() {
		mockdynamodb.New().DescribeTimeToLiveRequest(nil)
	})
}
