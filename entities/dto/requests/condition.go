package requests

import (
	"fmt"
)

type Operator string

const (
	Equal        Operator = "="
	NotEqual     Operator = "!="
	Ilike        Operator = "ILIKE"
	Greater      Operator = ">"
	GreaterEqual Operator = ">="
	Less         Operator = "<"
)

type Condition struct {
	Field     string
	Operation Operator
	Value     any
}

func NewCondition(field string, operation Operator, value any) *Condition {
	if value == "" {
		return nil
	}
	if operation == Ilike {
		value = "%" + fmt.Sprintf("%v", value) + "%"
	}
	return &Condition{
		Field:     field,
		Operation: operation,
		Value:     value,
	}
}
