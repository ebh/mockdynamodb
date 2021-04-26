package mockdynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

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
