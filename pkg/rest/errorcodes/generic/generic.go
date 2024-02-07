package generic

import (
	"strconv"
)

type ErrorCode uint

const (
	INTERNAL_ERROR ErrorCode = iota + 1 //Can be no access to database, network issue, missing server...
	BAD_REQUEST                         //Should not be used. Used temporary before the requests are correctly checked
	REQUIRED_FIELD
	NUMERIC_FIELD
	EMAIL_FIELD
	UNKNOWN
	DB_UNIQUE
	DB_RECORD_NOT_FOUND
	INVALID_OID
	UNKNOWN_TENANT_ID
	UNKNOWN_BATCH_ID
	UNKNOWN_DATAPROVIDER
	WORKLOW_COMPONENT
	UNKNOWN_FOLIO_ID
	PARSING_ERROR
	UNKNOWN_DEFAULT_BORROWER
)

func (code ErrorCode) String() string {
	return strconv.Itoa(int(code))
}
