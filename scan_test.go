package mockdynamodb_test

import (
	"context"
	"testing"

	"github.com/ebh/mockdynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
