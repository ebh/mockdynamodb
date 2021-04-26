package mockdynamodb

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Scan is not implemented. It will panic in all cases.
func (db *DynamoDb) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	db.Lock()
	defer db.Unlock()

	return db.scan(input)
}

// ScanWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ScanWithContext(_ aws.Context, input *dynamodb.ScanInput, _ ...request.Option) (*dynamodb.ScanOutput, error) {
	db.Lock()
	defer db.Unlock()

	return db.scan(input)
}

// ScanRequest is not implemented. It will panic in all cases.
func (db *DynamoDb) ScanRequest(*dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	panic("ScanRequest is not implemented")
}

// ScanPages is not implemented. It will panic in all cases.
func (db *DynamoDb) ScanPages(i *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool) error {
	db.Lock()
	defer db.Unlock()

	return db.scanPages(i, fn)
}

// ScanPagesWithContext is not implemented. It will panic in all cases.
func (db *DynamoDb) ScanPagesWithContext(_ aws.Context, i *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool, _ ...request.Option) error {
	db.Lock()
	defer db.Unlock()

	return db.scanPages(i, fn)
}

func (db *DynamoDb) scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	err := checkRequiredFields(map[string]interface{}{
		"ScanInput.TableName": input.TableName,
	})
	if err != nil {
		return &dynamodb.ScanOutput{}, err
	}

	if table := db.GetTable(*input.TableName); table != nil {
		table.receivedScanInputs = append(table.receivedScanInputs, *input)
		if table.moreScanOutputs() {
			o := table.popReturnScanOutput()
			if o != nil {
				return o, nil
			}
			return nil, errors.New("nil ScanOutput")
		}

		return nil, errors.New("no ScanOutputs to return")
	}

	return &dynamodb.ScanOutput{}, errorNonExistentTable()
}

func (db *DynamoDb) scanPages(input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool) error {
	err := checkRequiredFields(map[string]interface{}{
		"ScanInput.TableName": input.TableName,
	})
	if err != nil {
		return err
	}

	if table := db.GetTable(*input.TableName); table != nil {
		table.receivedScanInputs = append(table.receivedScanInputs, *input)
		if table.moreScanOutputs() {
			for ok := true; ok; ok = table.moreScanOutputs() {
				output := table.popReturnScanOutput()

				if output == nil {
					return errors.New("nil ScanOutput")
				}

				if !fn(output, false) {
					return errors.New("output handler function returned false")
				}
			}

			return nil
		}

		return errors.New("no ScanOutputs to return")
	}

	return errorNonExistentTable()
}
