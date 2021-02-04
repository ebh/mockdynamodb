package mockdynamodb

import "fmt"

func errorMissingField(field string) error {
	return fmt.Errorf("InvalidParameter: 1 validation error(s) found.\n- missing required field, %s.\n", field)
}

func errorNonExistentTable() error {
	return errorWithRequestID("AWS.DynamoDB.NonExistentTable: The specified table does not exist for this wsdl version.")
}

func errorWithRequestID(message string) error {
	return fmt.Errorf("%s\n\tstatus code: 400, request id: %s", message, newRequestID())
}
