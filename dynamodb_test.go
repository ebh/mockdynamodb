package mockdynamodb_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/ebh/mockdynamodb"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Implements(t, (*dynamodbiface.DynamoDBAPI)(nil), mockdynamodb.New())
}

func TestNewWithTables(t *testing.T) {
	t.Run("CreatedWithTableReturnsTable", func(t *testing.T) {
		db := mockdynamodb.NewWithTables([]string{"mockTable"})

		table := db.GetTable("mockTable")

		assert.NotNil(t, table)
	})
}
