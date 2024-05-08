package core

// Code - fail code.
type Code string

const (
	CodeOK = "OK"
	CodeDBFail = "DB_FAIL"
	CodeNoMultiplicator = "NO_SUCH_MULTIPLICATOR"
	CodeInternal = "INTERNAL_ERROR"
	CodeBadRequest = "BAD_REQUEST"
	CodeForbidden = "FORBIDDEN"
)
