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
	CodeIdCannotBeEmpty = "ID_CANNOT_BE_EMPTY"
	CodeScoreCannotBeEmpty = "SCORE_CANNOT_BE_EMPTY"
	CodeMultiplicatorCannotBeEmpty = "MULTIPLICATOR_CANNOT_BE_EMPTY"
)
