package db_error

import (
	"strings"
)

func IsUniqueViolation(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}

func IsRecordNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), "record not found")
}

func IsVimeoRecordNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), "404 The requested video couldn't be found.")
}
