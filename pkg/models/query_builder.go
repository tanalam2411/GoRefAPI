package models

import (
	"fmt"
	"gorm.io/gorm"
	v1 "ms-inventory/pkg/models/inventory/v1"
	"reflect"
	"strings"
)


// AND and OR operator, that should be used while building the query
const (
	AND = "and"
	OR  = "or"
)

// Comparison operator that can be used while building the where clause
const (
	EQUALS       = "="
	GREATER_THAN = ">"
	LESS_THAN    = ">"
	IN           = "in"
)

// Condition Type which make o=one comparison part of a where clause
// e.g. where `name == "Max"`
// Condition{ColumnName: "name", Operator: EQUALS, Value: "Max"}
type Condition struct {
	ColumnName string
	Operator   string
	Value      interface{}
}

// Formats Condition type to query string
// e.g., Condition{ColumnName: "name", Operator: EQUALS, Value: "Max"}
// `name` = ?
func (c *Condition) ToQuery() (string, interface{}) {
	return fmt.Sprintf("%s %s ?", c.ColumnName, c.Operator), c.Value
}

// QueryBuilder is an interface, to support multiple query builders,
// BuildQuery is one of them, if it doesn't suites certain requirement
// we can extend by creating multiple implementation of QueryBuilders
type QueryBuilder interface {
	Create(modelObject v1.BaseModel) error
	Get(modelObject v1.BaseModel, conditions []interface{}) error
	GetAll(modelObject v1.BaseModel, conditions []interface{}) (interface{}, error)
	Update(modelObject v1.BaseModel, conditions []interface{}, values v1.BaseModel) error
	Delete(modelObject v1.BaseModel, conditions []interface{}) error
}

// Constructs the query and its argument that should become input for DB's where clause
// 	conditions = append(conditions, models.Condition{
//		ColumnName: "name",
//		Operator:   models.EQUALS,
//		Value:      "Category 1",
//	})
//	conditions = append(conditions, models.AND)
//	conditions = append(conditions, models.Condition{
//		ColumnName: "parent_id",
//		Operator:   models.EQUALS,
//		Value:      "3",
//	})
//
//  query: name = ? and parent_id = ?
// args: []interface{"Category 1", "3"}
// 	query, args := BuildQuery(conditions)
//	err := sqb.DB.Where(query, args...).First(modelObject).Error
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

// One of the QueryBuild implementation
type SimpleQueryBuilder struct {
	DB *gorm.DB
}

// insert into operation
func (sqb *SimpleQueryBuilder) Create(modelObject v1.BaseModel) error {
	err := sqb.DB.Create(modelObject).Error
	return err
}

// select * from [table] where [condition]
func (sqb *SimpleQueryBuilder) Get(modelObject v1.BaseModel, conditions []interface{}) error {
	query, args := BuildQuery(conditions)
	err := sqb.DB.Where(query, args...).First(modelObject).Error
	return err
}

// update [table] where [condition] set [values]
func (sqb *SimpleQueryBuilder) Update(modelObject v1.BaseModel, conditions []interface{}, values v1.BaseModel) error {
	query, args := BuildQuery(conditions)
	err := sqb.DB.Where(query, args...).First(modelObject).Updates(values).Error
	return err
}

// delete from [table] where [condition]
func (sqb *SimpleQueryBuilder) Delete(modelObject v1.BaseModel, conditions []interface{}) error {
	query, args := BuildQuery(conditions)
	err := sqb.DB.Where(query, args...).Delete(modelObject).Error
	return err
}
