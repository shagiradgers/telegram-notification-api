package errors

import (
	"database/sql"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WrapToNetwork wraps sql error to network error
func WrapToNetwork(err error) *NetworkError {
	return wrapSQLErrorToNetwork(err)
}

func wrapSQLErrorToNetwork(err error) *NetworkError {
	n := &NetworkError{
		err: err,
	}

	if errors.Is(err, sql.ErrNoRows) {
		n.code = codes.NotFound
	}

	if n.code == 0 && n.err != nil {
		n.code = codes.Internal
	}
	return n
}

type NetworkError struct {
	err  error
	code codes.Code
}

func (n *NetworkError) Error() string {
	if n.err != nil {
		return n.err.Error()
	}
	return ""
}

func NewNetworkError(code codes.Code, message string) *NetworkError {
	n := &NetworkError{
		err:  errors.New(message),
		code: code,
	}
	return n
}

func (n *NetworkError) ToGRPCError() error {
	return status.Error(n.code, n.Error())
}
