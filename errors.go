package elastic

import "fmt"

type KVError interface {
	Type() string
	Key() string
	Value() string
	error
}

func NewIndexError(err error) *IndexError {
	return &IndexError{err: err}
}

type IndexError struct {
	err error
}

func (e *IndexError) Type() string {
	return ErrTypeIndex
}

func (e *IndexError) Key() string {
	return emptyStr
}

func (e *IndexError) Value() string {
	return emptyStr
}

func (e *IndexError) Error() string {
	return fmt.Sprintf("index failed:%s", e.err.Error())
}

func NewParseMessageError(err error) *ParseMessageError {
	return &ParseMessageError{err: err}
}

type ParseMessageError struct {
	err error
}

func (e *ParseMessageError) Type() string {
	return ErrTypeParse
}

func (e *ParseMessageError) Key() string {
	return emptyStr
}

func (e *ParseMessageError) Value() string {
	return emptyStr
}

func (e *ParseMessageError) Error() string {
	return fmt.Sprintf("parse message failed:%s", e.err.Error())
}

func NewParseRowDataError(err error) *ParseRowDataError {
	return &ParseRowDataError{err: err}
}

type ParseRowDataError struct {
	err error
}

func (e *ParseRowDataError) Type() string {
	return ErrTypeParse
}

func (e *ParseRowDataError) Key() string {
	return emptyStr
}

func (e *ParseRowDataError) Value() string {
	return emptyStr
}

func (e *ParseRowDataError) Error() string {
	return fmt.Sprintf("parse row data failed:%s", e.err.Error())
}

func NewUnKnowFieldValueError(field, value string) *UnKnowFieldValueError {
	return &UnKnowFieldValueError{field: field, value: value}
}

type UnKnowFieldValueError struct {
	field string
	value string
}

func (e *UnKnowFieldValueError) Type() string {
	return ErrTypeUnKnow
}

func (e *UnKnowFieldValueError) Key() string {
	return e.field
}

func (e *UnKnowFieldValueError) Value() string {
	return e.value
}

func (e *UnKnowFieldValueError) Error() string {
	return fmt.Sprintf("Unknow field %s's value [%s] ", e.field, e.value)
}

func NewTransFieldValueError(field, value string) *TransFieldValueError {
	return &TransFieldValueError{field: field, value: value}
}

type TransFieldValueError struct {
	field string
	value string
}

func (e *TransFieldValueError) Type() string {
	return ErrTypeTrans
}

func (e *TransFieldValueError) Key() string {
	return e.field
}

func (e *TransFieldValueError) Value() string {
	return e.value
}

func (e *TransFieldValueError) Error() string {
	return fmt.Sprintf("Trans field %s's value [%s] failed", e.field, e.value)
}

func NewEmptyFieldValueError(field string) *EmptyFieldValueError {
	return &EmptyFieldValueError{field: field}
}

type EmptyFieldValueError struct {
	field string
}

func (e *EmptyFieldValueError) Type() string {
	return ErrTypeEmpty
}

func (e *EmptyFieldValueError) Key() string {
	return e.field
}

func (e *EmptyFieldValueError) Value() string {
	return emptyStr
}

func (e *EmptyFieldValueError) Error() string {
	return fmt.Sprintf("Required field %s's value is empty", e.field)
}

const (
	ErrTypeIndex  = "index"
	ErrTypeParse  = "parse"
	ErrTypeUnKnow = "unKnow"
	ErrTypeTrans  = "trans"
	ErrTypeEmpty  = "empty"
	ErrTypeNormal = "normal"
)

var emptyStr = ""
