package models

import (
	"fmt"
	"gorm.io/gorm"
	v1 "ms-inventory/pkg/models/inventory/v1"
	"reflect"
	"strings"
)

const (
	AND = "and"
	OR  = "or"
)

const (
	EQUALS       = "="
	GREATER_THAN = ">"
	LESS_THAN    = ">"
	IN           = "in"
)

type Condition struct {
	ColumnName string
	Operator   string
	Value      interface{}
}

func (c *Condition) ToQuery() (string, interface{}) {
	return fmt.Sprintf("%s %s ?", c.ColumnName, c.Operator), c.Value
}

// QueryBuilder is an interface, to support multiple query builders,
// BuildQuery is one of them, if it doesn't suites certain requirement
// we can extend by creating multiple implementation of QueryBuilders
type QueryBuilder interface {
	Create(modelObject interface{}) error
	Get(modelObject interface{}, conditions []Condition) (interface{}, error)
	GetAll(modelObject interface{}, conditions []interface{}) (interface{}, error)
	Update(modelObject interface{}, conditions []interface{}, updateParams []interface{}) error
	Delete(modelObject interface{}, conditions []interface{}) error
}

func BuildQuery(conditions []interface{}) (string, []interface{}) {
	var args []interface{}
	var queries []string
	for _, val := range conditions {
		if typ := reflect.TypeOf(val).Kind(); typ == reflect.String {
			conc, _ := val.(string)
			queries = append(queries, conc)
		} else {
			cond, _ := val.(Condition)
			query, arg := cond.ToQuery()
			args = append(args, arg)
			queries = append(queries, query)
		}

	}

	finalQuery := strings.Join(queries, " ")
	fmt.Printf("Final query: %v\n", finalQuery)
	return finalQuery, args
}

type SimpleQueryBuilder struct {
	DB *gorm.DB
}

func (sqb *SimpleQueryBuilder) Create(modelObject v1.BaseModel) error {
	err := sqb.DB.Create(modelObject).Error
	return err
}

func (sqb *SimpleQueryBuilder) Get(modelObject v1.BaseModel, conditions []interface{}) error {
	query, args := BuildQuery(conditions)
	err := sqb.DB.Where(query, args...).First(modelObject).Error
	return err
}

func (sqb *SimpleQueryBuilder) Update(modelObject v1.BaseModel, conditions []interface{}, values map[string]interface{}) error {
	query, args := BuildQuery(conditions)
	err := sqb.DB.Where(query, args...).First(modelObject).Updates(values).Error
	return err
}

func (sqb *SimpleQueryBuilder) Delete(modelObject v1.BaseModel, conditions []interface{}) error {
	query, args := BuildQuery(conditions)
	err := sqb.DB.Where(query, args...).Delete(modelObject).Error
	return err
}
