package mysql_utils

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestErr {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "no rows") {
			return errors.NewNotFoundError(err.Error())
		}
		return errors.InternalServerError("error parsing db error")
	}
	switch mysqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicate key")
	}

	return errors.InternalServerError("error processing request")

}
