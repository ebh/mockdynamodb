package mockdynamodb_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ebh/mockdynamodb"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDb_BatchExecuteStatement(t *testing.T) {
	assert.PanicsWithValue(t, "BatchExecuteStatement is not implemented", func() {
		_, err := mockdynamodb.New().BatchExecuteStatement(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_BatchExecuteStatementWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "BatchExecuteStatementWithContext is not implemented", func() {
		_, err := mockdynamodb.New().BatchExecuteStatementWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_BatchExecuteStatementRequest(t *testing.T) {
	assert.PanicsWithValue(t, "BatchExecuteStatementRequest is not implemented", func() {
		mockdynamodb.New().BatchExecuteStatementRequest(nil)
	})
}

func TestDynamoDb_BatchGetItem(t *testing.T) {
	assert.PanicsWithValue(t, "BatchGetItem is not implemented", func() {
		_, err := mockdynamodb.New().BatchGetItem(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_BatchGetItemWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "BatchGetItemWithContext is not implemented", func() {
		_, err := mockdynamodb.New().BatchGetItemWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_BatchGetItemRequest(t *testing.T) {
	assert.PanicsWithValue(t, "BatchGetItemRequest is not implemented", func() {
		mockdynamodb.New().BatchGetItemRequest(nil)
	})
}

func TestDynamoDb_BatchGetItemPages(t *testing.T) {
	assert.PanicsWithValue(t, "BatchGetItemPages is not implemented", func() {
		err := mockdynamodb.New().BatchGetItemPages(nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_BatchGetItemPagesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "BatchGetItemPagesWithContext is not implemented", func() {
		err := mockdynamodb.New().BatchGetItemPagesWithContext(context.TODO(), nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_BatchWriteItem(t *testing.T) {
	assert.PanicsWithValue(t, "BatchWriteItem is not implemented", func() {
		_, err := mockdynamodb.New().BatchWriteItem(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_BatchWriteItemWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "BatchWriteItemWithContext is not implemented", func() {
		_, err := mockdynamodb.New().BatchWriteItemWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_BatchWriteItemRequest(t *testing.T) {
	assert.PanicsWithValue(t, "BatchWriteItemRequest is not implemented", func() {
		mockdynamodb.New().BatchWriteItemRequest(nil)
	})
}

func TestDynamoDb_CreateBackup(t *testing.T) {
	assert.PanicsWithValue(t, "CreateBackup is not implemented", func() {
		_, err := mockdynamodb.New().CreateBackup(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_CreateBackupWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "CreateBackupWithContext is not implemented", func() {
		_, err := mockdynamodb.New().CreateBackupWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_CreateBackupRequest(t *testing.T) {
	assert.PanicsWithValue(t, "CreateBackupRequest is not implemented", func() {
		mockdynamodb.New().CreateBackupRequest(nil)
	})
}

func TestDynamoDb_CreateGlobalTable(t *testing.T) {
	assert.PanicsWithValue(t, "CreateGlobalTable is not implemented", func() {
		_, err := mockdynamodb.New().CreateGlobalTable(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_CreateGlobalTableWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "CreateGlobalTableWithContext is not implemented", func() {
		_, err := mockdynamodb.New().CreateGlobalTableWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_CreateGlobalTableRequest(t *testing.T) {
	assert.PanicsWithValue(t, "CreateGlobalTableRequest is not implemented", func() {
		mockdynamodb.New().CreateGlobalTableRequest(nil)
	})
}

func TestDynamoDb_CreateTable(t *testing.T) {
	assert.PanicsWithValue(t, "CreateTable is not implemented", func() {
		_, err := mockdynamodb.New().CreateTable(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_CreateTableWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "CreateTableWithContext is not implemented", func() {
		_, err := mockdynamodb.New().CreateTableWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_CreateTableRequest(t *testing.T) {
	assert.PanicsWithValue(t, "CreateTableRequest is not implemented", func() {
		mockdynamodb.New().CreateTableRequest(nil)
	})
}

func TestDynamoDb_DeleteBackup(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteBackup is not implemented", func() {
		_, err := mockdynamodb.New().DeleteBackup(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DeleteBackupWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteBackupWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DeleteBackupWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DeleteBackupRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteBackupRequest is not implemented", func() {
		mockdynamodb.New().DeleteBackupRequest(nil)
	})
}

func TestDynamoDb_DeleteItem(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteItem is not implemented", func() {
		_, err := mockdynamodb.New().DeleteItem(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DeleteItemWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteItemWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DeleteItemWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DeleteItemRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteItemRequest is not implemented", func() {
		mockdynamodb.New().DeleteItemRequest(nil)
	})
}

func TestDynamoDb_DeleteTable(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteTable is not implemented", func() {
		_, err := mockdynamodb.New().DeleteTable(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DeleteTableWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteTableWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DeleteTableWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DeleteTableRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteTableRequest is not implemented", func() {
		mockdynamodb.New().DeleteTableRequest(nil)
	})
}

func TestDynamoDb_DisableKinesisStreamingDestination(t *testing.T) {
	assert.PanicsWithValue(t, "DisableKinesisStreamingDestination is not implemented", func() {
		_, err := mockdynamodb.New().DisableKinesisStreamingDestination(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DisableKinesisStreamingDestinationWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DisableKinesisStreamingDestinationWithContext is not implemented", func() {
		_, err := mockdynamodb.New().DisableKinesisStreamingDestinationWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_DisableKinesisStreamingDestinationRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DisableKinesisStreamingDestinationRequest is not implemented", func() {
		mockdynamodb.New().DisableKinesisStreamingDestinationRequest(nil)
	})
}

func TestDynamoDb_EnableKinesisStreamingDestination(t *testing.T) {
	assert.PanicsWithValue(t, "EnableKinesisStreamingDestination is not implemented", func() {
		_, err := mockdynamodb.New().EnableKinesisStreamingDestination(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_EnableKinesisStreamingDestinationWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "EnableKinesisStreamingDestinationWithContext is not implemented", func() {
		_, err := mockdynamodb.New().EnableKinesisStreamingDestinationWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_EnableKinesisStreamingDestinationRequest(t *testing.T) {
	assert.PanicsWithValue(t, "EnableKinesisStreamingDestinationRequest is not implemented", func() {
		mockdynamodb.New().EnableKinesisStreamingDestinationRequest(nil)
	})
}

func TestDynamoDb_ExecuteStatement(t *testing.T) {
	assert.PanicsWithValue(t, "ExecuteStatement is not implemented", func() {
		_, err := mockdynamodb.New().ExecuteStatement(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ExecuteStatementWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ExecuteStatementWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ExecuteStatementWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ExecuteStatementRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ExecuteStatementRequest is not implemented", func() {
		mockdynamodb.New().ExecuteStatementRequest(nil)
	})
}

func TestDynamoDb_ExecuteTransaction(t *testing.T) {
	assert.PanicsWithValue(t, "ExecuteTransaction is not implemented", func() {
		_, err := mockdynamodb.New().ExecuteTransaction(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ExecuteTransactionWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ExecuteTransactionWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ExecuteTransactionWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ExecuteTransactionRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ExecuteTransactionRequest is not implemented", func() {
		mockdynamodb.New().ExecuteTransactionRequest(nil)
	})
}

func TestDynamoDb_ExportTableToPointInTime(t *testing.T) {
	assert.PanicsWithValue(t, "ExportTableToPointInTime is not implemented", func() {
		_, err := mockdynamodb.New().ExportTableToPointInTime(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ExportTableToPointInTimeWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ExportTableToPointInTimeWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ExportTableToPointInTimeWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ExportTableToPointInTimeRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ExportTableToPointInTimeRequest is not implemented", func() {
		mockdynamodb.New().ExportTableToPointInTimeRequest(nil)
	})
}

func TestDynamoDb_GetItem(t *testing.T) {
	assert.PanicsWithValue(t, "GetItem is not implemented", func() {
		_, err := mockdynamodb.New().GetItem(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_GetItemWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "GetItemWithContext is not implemented", func() {
		_, err := mockdynamodb.New().GetItemWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_GetItemRequest(t *testing.T) {
	assert.PanicsWithValue(t, "GetItemRequest is not implemented", func() {
		mockdynamodb.New().GetItemRequest(nil)
	})
}

func TestDynamoDb_ListBackups(t *testing.T) {
	assert.PanicsWithValue(t, "ListBackups is not implemented", func() {
		_, err := mockdynamodb.New().ListBackups(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListBackupsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListBackupsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ListBackupsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListBackupsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListBackupsRequest is not implemented", func() {
		mockdynamodb.New().ListBackupsRequest(nil)
	})
}

func TestDynamoDb_ListContributorInsights(t *testing.T) {
	assert.PanicsWithValue(t, "ListContributorInsights is not implemented", func() {
		_, err := mockdynamodb.New().ListContributorInsights(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListContributorInsightsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListContributorInsightsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ListContributorInsightsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListContributorInsightsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListContributorInsightsRequest is not implemented", func() {
		mockdynamodb.New().ListContributorInsightsRequest(nil)
	})
}

func TestDynamoDb_ListContributorInsightsPages(t *testing.T) {
	assert.PanicsWithValue(t, "ListContributorInsightsPages is not implemented", func() {
		err := mockdynamodb.New().ListContributorInsightsPages(nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListContributorInsightsPagesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListContributorInsightsPagesWithContext is not implemented", func() {
		err := mockdynamodb.New().ListContributorInsightsPagesWithContext(context.TODO(), nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListExports(t *testing.T) {
	assert.PanicsWithValue(t, "ListExports is not implemented", func() {
		_, err := mockdynamodb.New().ListExports(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListExportsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListExportsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ListExportsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListExportsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListExportsRequest is not implemented", func() {
		mockdynamodb.New().ListExportsRequest(nil)
	})
}

func TestDynamoDb_ListExportsPages(t *testing.T) {
	assert.PanicsWithValue(t, "ListExportsPages is not implemented", func() {
		err := mockdynamodb.New().ListExportsPages(nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListExportsPagesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListExportsPagesWithContext is not implemented", func() {
		err := mockdynamodb.New().ListExportsPagesWithContext(context.TODO(), nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListGlobalTables(t *testing.T) {
	assert.PanicsWithValue(t, "ListGlobalTables is not implemented", func() {
		_, err := mockdynamodb.New().ListGlobalTables(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListGlobalTablesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListGlobalTablesWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ListGlobalTablesWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListGlobalTablesRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListGlobalTablesRequest is not implemented", func() {
		mockdynamodb.New().ListGlobalTablesRequest(nil)
	})
}

func TestDynamoDb_ListTables(t *testing.T) {
	assert.PanicsWithValue(t, "ListTables is not implemented", func() {
		_, err := mockdynamodb.New().ListTables(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListTablesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListTablesWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ListTablesWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListTablesRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListTablesRequest is not implemented", func() {
		mockdynamodb.New().ListTablesRequest(nil)
	})
}

func TestDynamoDb_ListTablesPages(t *testing.T) {
	assert.PanicsWithValue(t, "ListTablesPages is not implemented", func() {
		err := mockdynamodb.New().ListTablesPages(nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListTablesPagesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListTablesPagesWithContext is not implemented", func() {
		err := mockdynamodb.New().ListTablesPagesWithContext(context.TODO(), nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListTagsOfResource(t *testing.T) {
	assert.PanicsWithValue(t, "ListTagsOfResource is not implemented", func() {
		_, err := mockdynamodb.New().ListTagsOfResource(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListTagsOfResourceWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListTagsOfResourceWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ListTagsOfResourceWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ListTagsOfResourceRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListTagsOfResourceRequest is not implemented", func() {
		mockdynamodb.New().ListTagsOfResourceRequest(nil)
	})
}

func TestDynamoDb_Query(t *testing.T) {
	assert.PanicsWithValue(t, "Query is not implemented", func() {
		_, err := mockdynamodb.New().Query(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_QueryWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "QueryWithContext is not implemented", func() {
		_, err := mockdynamodb.New().QueryWithContext(context.TODO(), nil)
		require.NoError(t, err)
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

func TestDynamoDb_RestoreTableFromBackup(t *testing.T) {
	assert.PanicsWithValue(t, "RestoreTableFromBackup is not implemented", func() {
		_, err := mockdynamodb.New().RestoreTableFromBackup(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_RestoreTableFromBackupWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "RestoreTableFromBackupWithContext is not implemented", func() {
		_, err := mockdynamodb.New().RestoreTableFromBackupWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_RestoreTableFromBackupRequest(t *testing.T) {
	assert.PanicsWithValue(t, "RestoreTableFromBackupRequest is not implemented", func() {
		mockdynamodb.New().RestoreTableFromBackupRequest(nil)
	})
}

func TestDynamoDb_RestoreTableToPointInTime(t *testing.T) {
	assert.PanicsWithValue(t, "RestoreTableToPointInTime is not implemented", func() {
		_, err := mockdynamodb.New().RestoreTableToPointInTime(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_RestoreTableToPointInTimeWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "RestoreTableToPointInTimeWithContext is not implemented", func() {
		_, err := mockdynamodb.New().RestoreTableToPointInTimeWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_RestoreTableToPointInTimeRequest(t *testing.T) {
	assert.PanicsWithValue(t, "RestoreTableToPointInTimeRequest is not implemented", func() {
		mockdynamodb.New().RestoreTableToPointInTimeRequest(nil)
	})
}

func TestDynamoDb_Scan(t *testing.T) {
	assert.PanicsWithValue(t, "Scan is not implemented", func() {
		_, err := mockdynamodb.New().Scan(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ScanWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ScanWithContext is not implemented", func() {
		_, err := mockdynamodb.New().ScanWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ScanRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ScanRequest is not implemented", func() {
		mockdynamodb.New().ScanRequest(nil)
	})
}

func TestDynamoDb_ScanPages(t *testing.T) {
	assert.PanicsWithValue(t, "ScanPages is not implemented", func() {
		err := mockdynamodb.New().ScanPages(nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_ScanPagesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ScanPagesWithContext is not implemented", func() {
		err := mockdynamodb.New().ScanPagesWithContext(context.TODO(), nil, nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_TagResource(t *testing.T) {
	assert.PanicsWithValue(t, "TagResource is not implemented", func() {
		_, err := mockdynamodb.New().TagResource(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_TagResourceWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "TagResourceWithContext is not implemented", func() {
		_, err := mockdynamodb.New().TagResourceWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_TagResourceRequest(t *testing.T) {
	assert.PanicsWithValue(t, "TagResourceRequest is not implemented", func() {
		mockdynamodb.New().TagResourceRequest(nil)
	})
}

func TestDynamoDb_TransactGetItems(t *testing.T) {
	assert.PanicsWithValue(t, "TransactGetItems is not implemented", func() {
		_, err := mockdynamodb.New().TransactGetItems(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_TransactGetItemsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "TransactGetItemsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().TransactGetItemsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_TransactGetItemsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "TransactGetItemsRequest is not implemented", func() {
		mockdynamodb.New().TransactGetItemsRequest(nil)
	})
}

func TestDynamoDb_TransactWriteItems(t *testing.T) {
	assert.PanicsWithValue(t, "TransactWriteItems is not implemented", func() {
		_, err := mockdynamodb.New().TransactWriteItems(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_TransactWriteItemsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "TransactWriteItemsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().TransactWriteItemsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_TransactWriteItemsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "TransactWriteItemsRequest is not implemented", func() {
		mockdynamodb.New().TransactWriteItemsRequest(nil)
	})
}

func TestDynamoDb_UntagResource(t *testing.T) {
	assert.PanicsWithValue(t, "UntagResource is not implemented", func() {
		_, err := mockdynamodb.New().UntagResource(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UntagResourceWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UntagResourceWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UntagResourceWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UntagResourceRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UntagResourceRequest is not implemented", func() {
		mockdynamodb.New().UntagResourceRequest(nil)
	})
}

func TestDynamoDb_UpdateContinuousBackups(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateContinuousBackups is not implemented", func() {
		_, err := mockdynamodb.New().UpdateContinuousBackups(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateContinuousBackupsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateContinuousBackupsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UpdateContinuousBackupsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateContinuousBackupsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateContinuousBackupsRequest is not implemented", func() {
		mockdynamodb.New().UpdateContinuousBackupsRequest(nil)
	})
}

func TestDynamoDb_UpdateContributorInsights(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateContributorInsights is not implemented", func() {
		_, err := mockdynamodb.New().UpdateContributorInsights(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateContributorInsightsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateContributorInsightsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UpdateContributorInsightsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateContributorInsightsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateContributorInsightsRequest is not implemented", func() {
		mockdynamodb.New().UpdateContributorInsightsRequest(nil)
	})
}

func TestDynamoDb_UpdateGlobalTable(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateGlobalTable is not implemented", func() {
		_, err := mockdynamodb.New().UpdateGlobalTable(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateGlobalTableWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateGlobalTableWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UpdateGlobalTableWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateGlobalTableRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateGlobalTableRequest is not implemented", func() {
		mockdynamodb.New().UpdateGlobalTableRequest(nil)
	})
}

func TestDynamoDb_UpdateGlobalTableSettings(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateGlobalTableSettings is not implemented", func() {
		_, err := mockdynamodb.New().UpdateGlobalTableSettings(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateGlobalTableSettingsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateGlobalTableSettingsWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UpdateGlobalTableSettingsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateGlobalTableSettingsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateGlobalTableSettingsRequest is not implemented", func() {
		mockdynamodb.New().UpdateGlobalTableSettingsRequest(nil)
	})
}

func TestDynamoDb_UpdateItem(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateItem is not implemented", func() {
		_, err := mockdynamodb.New().UpdateItem(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateItemWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateItemWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UpdateItemWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateItemRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateItemRequest is not implemented", func() {
		mockdynamodb.New().UpdateItemRequest(nil)
	})
}

func TestDynamoDb_UpdateTable(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTable is not implemented", func() {
		_, err := mockdynamodb.New().UpdateTable(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateTableWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTableWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UpdateTableWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateTableRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTableRequest is not implemented", func() {
		mockdynamodb.New().UpdateTableRequest(nil)
	})
}

func TestDynamoDb_UpdateTableReplicaAutoScaling(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTableReplicaAutoScaling is not implemented", func() {
		_, err := mockdynamodb.New().UpdateTableReplicaAutoScaling(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateTableReplicaAutoScalingWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTableReplicaAutoScalingWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UpdateTableReplicaAutoScalingWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateTableReplicaAutoScalingRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTableReplicaAutoScalingRequest is not implemented", func() {
		mockdynamodb.New().UpdateTableReplicaAutoScalingRequest(nil)
	})
}

func TestDynamoDb_UpdateTimeToLive(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTimeToLive is not implemented", func() {
		_, err := mockdynamodb.New().UpdateTimeToLive(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateTimeToLiveWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTimeToLiveWithContext is not implemented", func() {
		_, err := mockdynamodb.New().UpdateTimeToLiveWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_UpdateTimeToLiveRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UpdateTimeToLiveRequest is not implemented", func() {
		mockdynamodb.New().UpdateTimeToLiveRequest(nil)
	})
}

func TestDynamoDb_WaitUntilTableExists(t *testing.T) {
	assert.PanicsWithValue(t, "WaitUntilTableExists is not implemented", func() {
		err := mockdynamodb.New().WaitUntilTableExists(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_WaitUntilTableExistsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "WaitUntilTableExistsWithContext is not implemented", func() {
		err := mockdynamodb.New().WaitUntilTableExistsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_WaitUntilTableNotExists(t *testing.T) {
	assert.PanicsWithValue(t, "WaitUntilTableNotExists is not implemented", func() {
		err := mockdynamodb.New().WaitUntilTableNotExists(nil)
		require.NoError(t, err)
	})
}

func TestDynamoDb_WaitUntilTableNotExistsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "WaitUntilTableNotExistsWithContext is not implemented", func() {
		err := mockdynamodb.New().WaitUntilTableNotExistsWithContext(context.TODO(), nil)
		require.NoError(t, err)
	})
}
