package mockdynamodb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const uuidRegexp = `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`

func assertRegexpError(t *testing.T, err error, regexp string) {
	require.Error(t, err)
	assert.Regexp(t, regexp, err.Error())
}
