package mockdynamodb

import (
	"reflect"

	"github.com/google/uuid"
)

func checkRequiredFields(fields map[string]interface{}) error {
	for field, value := range fields {
		if reflect.ValueOf(value).IsNil() {
			return errorMissingField(field)
		}
	}

	return nil
}

func newRequestID() string {
	return uuid.New().String()
}
