package domain

import "fmt"

type ErrorKind struct {
	code string
}

func (e ErrorKind) String() string {
	return e.code
}

var (
	ErrorKindInvalid  = ErrorKind{"invalid"}
	ErrorKindConflict = ErrorKind{"conflict"}
	ErrorKindNotFound = ErrorKind{"not found"}
	ErrorKindInternal = ErrorKind{"internal"}
)

type Error struct {
	kind ErrorKind
	msg  string
}

func EInvalid(msg string) error {
	return Error{kind: ErrorKindInvalid, msg: msg}
}

func EConflict(msg string) error {
	return Error{kind: ErrorKindConflict, msg: msg}
}

func ENotFound(msg string) error {
	return Error{kind: ErrorKindNotFound, msg: msg}
}

func EInternal(msg string) error {
	return Error{kind: ErrorKindInternal, msg: msg}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.kind, e.msg)
}
