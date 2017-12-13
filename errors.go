package pgup

import (
	"github.com/lib/pq"
	"upper.io/db.v3"
)

const pgErrorUniqueViolation = "23505"

func IsErrNoMoreRows(err error) bool {
	return err == db.ErrNoMoreRows
}

func IsPgDup(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		if err.Code == pgErrorUniqueViolation {
			return true
		}
	}
	return false
}
