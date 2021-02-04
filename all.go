package mockdynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// BatchExecuteStatement is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchExecuteStatement(*dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput, error) {
	panic("BatchExecuteStatement is not implemented")
}

// BatchExecuteStatementWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchExecuteStatementWithContext(aws.Context, *dynamodb.BatchExecuteStatementInput, ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error) {
	panic("BatchExecuteStatementWithContext is not implemented")
}

// BatchExecuteStatementRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchExecuteStatementRequest(*dynamodb.BatchExecuteStatementInput) (*request.Request, *dynamodb.BatchExecuteStatementOutput) {
	panic("BatchExecuteStatementRequest is not implemented")
}

// BatchGetItem is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchGetItem(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	panic("BatchGetItem is not implemented")
}

// BatchGetItemWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchGetItemWithContext(aws.Context, *dynamodb.BatchGetItemInput, ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	panic("BatchGetItemWithContext is not implemented")
}

// BatchGetItemRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchGetItemRequest(*dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	panic("BatchGetItemRequest is not implemented")
}

// BatchGetItemPages is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchGetItemPages(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	panic("BatchGetItemPages is not implemented")
}

// BatchGetItemPagesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchGetItemPagesWithContext(aws.Context, *dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool, ...request.Option) error {
	panic("BatchGetItemPagesWithContext is not implemented")
}

// BatchWriteItem is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchWriteItem(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	panic("BatchWriteItem is not implemented")
}

// BatchWriteItemWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchWriteItemWithContext(aws.Context, *dynamodb.BatchWriteItemInput, ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	panic("BatchWriteItemWithContext is not implemented")
}

// BatchWriteItemRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) BatchWriteItemRequest(*dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	panic("BatchWriteItemRequest is not implemented")
}

// CreateBackup is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateBackup(*dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	panic("CreateBackup is not implemented")
}

// CreateBackupWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateBackupWithContext(aws.Context, *dynamodb.CreateBackupInput, ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	panic("CreateBackupWithContext is not implemented")
}

// CreateBackupRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateBackupRequest(*dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	panic("CreateBackupRequest is not implemented")
}

// CreateGlobalTable is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateGlobalTable(*dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	panic("CreateGlobalTable is not implemented")
}

// CreateGlobalTableWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateGlobalTableWithContext(aws.Context, *dynamodb.CreateGlobalTableInput, ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	panic("CreateGlobalTableWithContext is not implemented")
}

// CreateGlobalTableRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateGlobalTableRequest(*dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	panic("CreateGlobalTableRequest is not implemented")
}

// CreateTable is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	panic("CreateTable is not implemented")
}

// CreateTableWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateTableWithContext(aws.Context, *dynamodb.CreateTableInput, ...request.Option) (*dynamodb.CreateTableOutput, error) {
	panic("CreateTableWithContext is not implemented")
}

// CreateTableRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) CreateTableRequest(*dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	panic("CreateTableRequest is not implemented")
}

// DeleteBackup is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteBackup(*dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	panic("DeleteBackup is not implemented")
}

// DeleteBackupWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteBackupWithContext(aws.Context, *dynamodb.DeleteBackupInput, ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	panic("DeleteBackupWithContext is not implemented")
}

// DeleteBackupRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteBackupRequest(*dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	panic("DeleteBackupRequest is not implemented")
}

// DeleteItem is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	panic("DeleteItem is not implemented")
}

// DeleteItemWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteItemWithContext(aws.Context, *dynamodb.DeleteItemInput, ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	panic("DeleteItemWithContext is not implemented")
}

// DeleteItemRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteItemRequest(*dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	panic("DeleteItemRequest is not implemented")
}

// DeleteTable is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteTable(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	panic("DeleteTable is not implemented")
}

// DeleteTableWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteTableWithContext(aws.Context, *dynamodb.DeleteTableInput, ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	panic("DeleteTableWithContext is not implemented")
}

// DeleteTableRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) DeleteTableRequest(*dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	panic("DeleteTableRequest is not implemented")
}

// DisableKinesisStreamingDestination is not implemented. It will panic in all cases.
func (db *DynamoDb) DisableKinesisStreamingDestination(*dynamodb.DisableKinesisStreamingDestinationInput) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	panic("DisableKinesisStreamingDestination is not implemented")
}

// DisableKinesisStreamingDestinationWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) DisableKinesisStreamingDestinationWithContext(aws.Context, *dynamodb.DisableKinesisStreamingDestinationInput, ...request.Option) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	panic("DisableKinesisStreamingDestinationWithContext is not implemented")
}

// DisableKinesisStreamingDestinationRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) DisableKinesisStreamingDestinationRequest(*dynamodb.DisableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.DisableKinesisStreamingDestinationOutput) {
	panic("DisableKinesisStreamingDestinationRequest is not implemented")
}

// EnableKinesisStreamingDestination is not implemented. It will panic in all cases.
func (db *DynamoDb) EnableKinesisStreamingDestination(*dynamodb.EnableKinesisStreamingDestinationInput) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	panic("EnableKinesisStreamingDestination is not implemented")
}

// EnableKinesisStreamingDestinationWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) EnableKinesisStreamingDestinationWithContext(aws.Context, *dynamodb.EnableKinesisStreamingDestinationInput, ...request.Option) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	panic("EnableKinesisStreamingDestinationWithContext is not implemented")
}

// EnableKinesisStreamingDestinationRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) EnableKinesisStreamingDestinationRequest(*dynamodb.EnableKinesisStreamingDestinationInput) (*request.Request, *dynamodb.EnableKinesisStreamingDestinationOutput) {
	panic("EnableKinesisStreamingDestinationRequest is not implemented")
}

// ExecuteStatement is not implemented. It will panic in all cases.
func (db *DynamoDb) ExecuteStatement(*dynamodb.ExecuteStatementInput) (*dynamodb.ExecuteStatementOutput, error) {
	panic("ExecuteStatement is not implemented")
}

// ExecuteStatementWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ExecuteStatementWithContext(aws.Context, *dynamodb.ExecuteStatementInput, ...request.Option) (*dynamodb.ExecuteStatementOutput, error) {
	panic("ExecuteStatementWithContext is not implemented")
}

// ExecuteStatementRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ExecuteStatementRequest(*dynamodb.ExecuteStatementInput) (*request.Request, *dynamodb.ExecuteStatementOutput) {
	panic("ExecuteStatementRequest is not implemented")
}

// ExecuteTransaction is not implemented. It will panic in all cases.
func (db *DynamoDb) ExecuteTransaction(*dynamodb.ExecuteTransactionInput) (*dynamodb.ExecuteTransactionOutput, error) {
	panic("ExecuteTransaction is not implemented")
}

// ExecuteTransactionWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ExecuteTransactionWithContext(aws.Context, *dynamodb.ExecuteTransactionInput, ...request.Option) (*dynamodb.ExecuteTransactionOutput, error) {
	panic("ExecuteTransactionWithContext is not implemented")
}

// ExecuteTransactionRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ExecuteTransactionRequest(*dynamodb.ExecuteTransactionInput) (*request.Request, *dynamodb.ExecuteTransactionOutput) {
	panic("ExecuteTransactionRequest is not implemented")
}

// ExportTableToPointInTime is not implemented. It will panic in all cases.
func (db *DynamoDb) ExportTableToPointInTime(*dynamodb.ExportTableToPointInTimeInput) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	panic("ExportTableToPointInTime is not implemented")
}

// ExportTableToPointInTimeWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ExportTableToPointInTimeWithContext(aws.Context, *dynamodb.ExportTableToPointInTimeInput, ...request.Option) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	panic("ExportTableToPointInTimeWithContext is not implemented")
}

// ExportTableToPointInTimeRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ExportTableToPointInTimeRequest(*dynamodb.ExportTableToPointInTimeInput) (*request.Request, *dynamodb.ExportTableToPointInTimeOutput) {
	panic("ExportTableToPointInTimeRequest is not implemented")
}

// GetItem is not implemented. It will panic in all cases.
func (db *DynamoDb) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	panic("GetItem is not implemented")
}

// GetItemWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) GetItemWithContext(aws.Context, *dynamodb.GetItemInput, ...request.Option) (*dynamodb.GetItemOutput, error) {
	panic("GetItemWithContext is not implemented")
}

// GetItemRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) GetItemRequest(*dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	panic("GetItemRequest is not implemented")
}

// ListBackups is not implemented. It will panic in all cases.
func (db *DynamoDb) ListBackups(*dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	panic("ListBackups is not implemented")
}

// ListBackupsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListBackupsWithContext(aws.Context, *dynamodb.ListBackupsInput, ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	panic("ListBackupsWithContext is not implemented")
}

// ListBackupsRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ListBackupsRequest(*dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	panic("ListBackupsRequest is not implemented")
}

// ListContributorInsights is not implemented. It will panic in all cases.
func (db *DynamoDb) ListContributorInsights(*dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	panic("ListContributorInsights is not implemented")
}

// ListContributorInsightsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListContributorInsightsWithContext(aws.Context, *dynamodb.ListContributorInsightsInput, ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	panic("ListContributorInsightsWithContext is not implemented")
}

// ListContributorInsightsRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ListContributorInsightsRequest(*dynamodb.ListContributorInsightsInput) (*request.Request, *dynamodb.ListContributorInsightsOutput) {
	panic("ListContributorInsightsRequest is not implemented")
}

// ListContributorInsightsPages is not implemented. It will panic in all cases.
func (db *DynamoDb) ListContributorInsightsPages(*dynamodb.ListContributorInsightsInput, func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	panic("ListContributorInsightsPages is not implemented")
}

// ListContributorInsightsPagesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListContributorInsightsPagesWithContext(aws.Context, *dynamodb.ListContributorInsightsInput, func(*dynamodb.ListContributorInsightsOutput, bool) bool, ...request.Option) error {
	panic("ListContributorInsightsPagesWithContext is not implemented")
}

// ListExports is not implemented. It will panic in all cases.
func (db *DynamoDb) ListExports(*dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput, error) {
	panic("ListExports is not implemented")
}

// ListExportsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListExportsWithContext(aws.Context, *dynamodb.ListExportsInput, ...request.Option) (*dynamodb.ListExportsOutput, error) {
	panic("ListExportsWithContext is not implemented")
}

// ListExportsRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ListExportsRequest(*dynamodb.ListExportsInput) (*request.Request, *dynamodb.ListExportsOutput) {
	panic("ListExportsRequest is not implemented")
}

// ListExportsPages is not implemented. It will panic in all cases.
func (db *DynamoDb) ListExportsPages(*dynamodb.ListExportsInput, func(*dynamodb.ListExportsOutput, bool) bool) error {
	panic("ListExportsPages is not implemented")
}

// ListExportsPagesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListExportsPagesWithContext(aws.Context, *dynamodb.ListExportsInput, func(*dynamodb.ListExportsOutput, bool) bool, ...request.Option) error {
	panic("ListExportsPagesWithContext is not implemented")
}

// ListGlobalTables is not implemented. It will panic in all cases.
func (db *DynamoDb) ListGlobalTables(*dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	panic("ListGlobalTables is not implemented")
}

// ListGlobalTablesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListGlobalTablesWithContext(aws.Context, *dynamodb.ListGlobalTablesInput, ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	panic("ListGlobalTablesWithContext is not implemented")
}

// ListGlobalTablesRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ListGlobalTablesRequest(*dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	panic("ListGlobalTablesRequest is not implemented")
}

// ListTables is not implemented. It will panic in all cases.
func (db *DynamoDb) ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	panic("ListTables is not implemented")
}

// ListTablesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListTablesWithContext(aws.Context, *dynamodb.ListTablesInput, ...request.Option) (*dynamodb.ListTablesOutput, error) {
	panic("ListTablesWithContext is not implemented")
}

// ListTablesRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ListTablesRequest(*dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	panic("ListTablesRequest is not implemented")
}

// ListTablesPages is not implemented. It will panic in all cases.
func (db *DynamoDb) ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error {
	panic("ListTablesPages is not implemented")
}

// ListTablesPagesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListTablesPagesWithContext(aws.Context, *dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool, ...request.Option) error {
	panic("ListTablesPagesWithContext is not implemented")
}

// ListTagsOfResource is not implemented. It will panic in all cases.
func (db *DynamoDb) ListTagsOfResource(*dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	panic("ListTagsOfResource is not implemented")
}

// ListTagsOfResourceWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ListTagsOfResourceWithContext(aws.Context, *dynamodb.ListTagsOfResourceInput, ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	panic("ListTagsOfResourceWithContext is not implemented")
}

// ListTagsOfResourceRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ListTagsOfResourceRequest(*dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	panic("ListTagsOfResourceRequest is not implemented")
}

// Query is not implemented. It will panic in all cases.
func (db *DynamoDb) Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	panic("Query is not implemented")
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

// RestoreTableFromBackup is not implemented. It will panic in all cases.
func (db *DynamoDb) RestoreTableFromBackup(*dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	panic("RestoreTableFromBackup is not implemented")
}

// RestoreTableFromBackupWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) RestoreTableFromBackupWithContext(aws.Context, *dynamodb.RestoreTableFromBackupInput, ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	panic("RestoreTableFromBackupWithContext is not implemented")
}

// RestoreTableFromBackupRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) RestoreTableFromBackupRequest(*dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	panic("RestoreTableFromBackupRequest is not implemented")
}

// RestoreTableToPointInTime is not implemented. It will panic in all cases.
func (db *DynamoDb) RestoreTableToPointInTime(*dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	panic("RestoreTableToPointInTime is not implemented")
}

// RestoreTableToPointInTimeWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) RestoreTableToPointInTimeWithContext(aws.Context, *dynamodb.RestoreTableToPointInTimeInput, ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	panic("RestoreTableToPointInTimeWithContext is not implemented")
}

// RestoreTableToPointInTimeRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) RestoreTableToPointInTimeRequest(*dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	panic("RestoreTableToPointInTimeRequest is not implemented")
}

// Scan is not implemented. It will panic in all cases.
func (db *DynamoDb) Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	panic("Scan is not implemented")
}

// ScanWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ScanWithContext(aws.Context, *dynamodb.ScanInput, ...request.Option) (*dynamodb.ScanOutput, error) {
	panic("ScanWithContext is not implemented")
}

// ScanRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ScanRequest(*dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	panic("ScanRequest is not implemented")
}

// ScanPages is not implemented. It will panic in all cases.
func (db *DynamoDb) ScanPages(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error {
	panic("ScanPages is not implemented")
}

// ScanPagesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ScanPagesWithContext(aws.Context, *dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool, ...request.Option) error {
	panic("ScanPagesWithContext is not implemented")
}

// TagResource is not implemented. It will panic in all cases.
func (db *DynamoDb) TagResource(*dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	panic("TagResource is not implemented")
}

// TagResourceWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) TagResourceWithContext(aws.Context, *dynamodb.TagResourceInput, ...request.Option) (*dynamodb.TagResourceOutput, error) {
	panic("TagResourceWithContext is not implemented")
}

// TagResourceRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) TagResourceRequest(*dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	panic("TagResourceRequest is not implemented")
}

// TransactGetItems is not implemented. It will panic in all cases.
func (db *DynamoDb) TransactGetItems(*dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	panic("TransactGetItems is not implemented")
}

// TransactGetItemsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) TransactGetItemsWithContext(aws.Context, *dynamodb.TransactGetItemsInput, ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	panic("TransactGetItemsWithContext is not implemented")
}

// TransactGetItemsRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) TransactGetItemsRequest(*dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	panic("TransactGetItemsRequest is not implemented")
}

// TransactWriteItems is not implemented. It will panic in all cases.
func (db *DynamoDb) TransactWriteItems(*dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("TransactWriteItems is not implemented")
}

// TransactWriteItemsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) TransactWriteItemsWithContext(aws.Context, *dynamodb.TransactWriteItemsInput, ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	panic("TransactWriteItemsWithContext is not implemented")
}

// TransactWriteItemsRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) TransactWriteItemsRequest(*dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	panic("TransactWriteItemsRequest is not implemented")
}

// UntagResource is not implemented. It will panic in all cases.
func (db *DynamoDb) UntagResource(*dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	panic("UntagResource is not implemented")
}

// UntagResourceWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UntagResourceWithContext(aws.Context, *dynamodb.UntagResourceInput, ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	panic("UntagResourceWithContext is not implemented")
}

// UntagResourceRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UntagResourceRequest(*dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	panic("UntagResourceRequest is not implemented")
}

// UpdateContinuousBackups is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateContinuousBackups(*dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	panic("UpdateContinuousBackups is not implemented")
}

// UpdateContinuousBackupsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateContinuousBackupsWithContext(aws.Context, *dynamodb.UpdateContinuousBackupsInput, ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	panic("UpdateContinuousBackupsWithContext is not implemented")
}

// UpdateContinuousBackupsRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateContinuousBackupsRequest(*dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	panic("UpdateContinuousBackupsRequest is not implemented")
}

// UpdateContributorInsights is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateContributorInsights(*dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
	panic("UpdateContributorInsights is not implemented")
}

// UpdateContributorInsightsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateContributorInsightsWithContext(aws.Context, *dynamodb.UpdateContributorInsightsInput, ...request.Option) (*dynamodb.UpdateContributorInsightsOutput, error) {
	panic("UpdateContributorInsightsWithContext is not implemented")
}

// UpdateContributorInsightsRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateContributorInsightsRequest(*dynamodb.UpdateContributorInsightsInput) (*request.Request, *dynamodb.UpdateContributorInsightsOutput) {
	panic("UpdateContributorInsightsRequest is not implemented")
}

// UpdateGlobalTable is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateGlobalTable(*dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	panic("UpdateGlobalTable is not implemented")
}

// UpdateGlobalTableWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateGlobalTableWithContext(aws.Context, *dynamodb.UpdateGlobalTableInput, ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	panic("UpdateGlobalTableWithContext is not implemented")
}

// UpdateGlobalTableRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateGlobalTableRequest(*dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	panic("UpdateGlobalTableRequest is not implemented")
}

// UpdateGlobalTableSettings is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateGlobalTableSettings(*dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	panic("UpdateGlobalTableSettings is not implemented")
}

// UpdateGlobalTableSettingsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateGlobalTableSettingsWithContext(aws.Context, *dynamodb.UpdateGlobalTableSettingsInput, ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	panic("UpdateGlobalTableSettingsWithContext is not implemented")
}

// UpdateGlobalTableSettingsRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateGlobalTableSettingsRequest(*dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	panic("UpdateGlobalTableSettingsRequest is not implemented")
}

// UpdateItem is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	panic("UpdateItem is not implemented")
}

// UpdateItemWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateItemWithContext(aws.Context, *dynamodb.UpdateItemInput, ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	panic("UpdateItemWithContext is not implemented")
}

// UpdateItemRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateItemRequest(*dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	panic("UpdateItemRequest is not implemented")
}

// UpdateTable is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTable(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	panic("UpdateTable is not implemented")
}

// UpdateTableWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTableWithContext(aws.Context, *dynamodb.UpdateTableInput, ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	panic("UpdateTableWithContext is not implemented")
}

// UpdateTableRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTableRequest(*dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	panic("UpdateTableRequest is not implemented")
}

// UpdateTableReplicaAutoScaling is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTableReplicaAutoScaling(*dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	panic("UpdateTableReplicaAutoScaling is not implemented")
}

// UpdateTableReplicaAutoScalingWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTableReplicaAutoScalingWithContext(aws.Context, *dynamodb.UpdateTableReplicaAutoScalingInput, ...request.Option) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	panic("UpdateTableReplicaAutoScalingWithContext is not implemented")
}

// UpdateTableReplicaAutoScalingRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTableReplicaAutoScalingRequest(*dynamodb.UpdateTableReplicaAutoScalingInput) (*request.Request, *dynamodb.UpdateTableReplicaAutoScalingOutput) {
	panic("UpdateTableReplicaAutoScalingRequest is not implemented")
}

// UpdateTimeToLive is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTimeToLive(*dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	panic("UpdateTimeToLive is not implemented")
}

// UpdateTimeToLiveWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTimeToLiveWithContext(aws.Context, *dynamodb.UpdateTimeToLiveInput, ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	panic("UpdateTimeToLiveWithContext is not implemented")
}

// UpdateTimeToLiveRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) UpdateTimeToLiveRequest(*dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	panic("UpdateTimeToLiveRequest is not implemented")
}

// WaitUntilTableExists is not implemented. It will panic in all cases.
func (db *DynamoDb) WaitUntilTableExists(*dynamodb.DescribeTableInput) error {
	panic("WaitUntilTableExists is not implemented")
}

// WaitUntilTableExistsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) WaitUntilTableExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error {
	panic("WaitUntilTableExistsWithContext is not implemented")
}

// WaitUntilTableNotExists is not implemented. It will panic in all cases.
func (db *DynamoDb) WaitUntilTableNotExists(*dynamodb.DescribeTableInput) error {
	panic("WaitUntilTableNotExists is not implemented")
}

// WaitUntilTableNotExistsWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) WaitUntilTableNotExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...request.WaiterOption) error {
	panic("WaitUntilTableNotExistsWithContext is not implemented")
}
