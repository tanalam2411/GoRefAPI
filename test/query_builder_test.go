package test

import (
	"fmt"
	"ms-inventory/global"
	"ms-inventory/initialize"
	"ms-inventory/pkg/models"
	v1 "ms-inventory/pkg/models/inventory/v1"
	"testing"
)

func init() {
	global.GDB = initialize.GormMysql()
}

func TestCreateCategory(t *testing.T) {

	c := v1.InventoryCategory{
		ParentID:    3,
		Name:        "Cat2",
		Description: "Cat2 desc",
	}

	global.GDB = initialize.GormMysql()
	sqb := models.SimpleQueryBuilder{DB: global.GDB}

	err := sqb.Create(&c)

	if err != nil {
		t.Errorf("Failed to create Cateogry: %v", err)
	}
}

func TestGetCategory(t *testing.T) {

	var c v1.InventoryCategory
	sqb := models.SimpleQueryBuilder{DB: global.GDB}

	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "id",
		Operator:   models.EQUALS,
		Value:      "1",
	})
	conditions = append(conditions, models.AND)
	conditions = append(conditions, models.Condition{
		ColumnName: "parent_id",
		Operator:   models.EQUALS,
		Value:      "3",
	})
	err := sqb.Get(&c, conditions)

	fmt.Println("Found category: ", c)

	if err != nil {
		t.Errorf("Failed to get Cateogry: %v", err)
	}

}

func TestUpdateCategory(t *testing.T) {

	var c v1.InventoryCategory
	sqb := models.SimpleQueryBuilder{DB: global.GDB}

	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "id",
		Operator:   models.EQUALS,
		Value:      "1",
	})
	conditions = append(conditions, models.AND)
	conditions = append(conditions, models.Condition{
		ColumnName: "parent_id",
		Operator:   models.EQUALS,
		Value:      "3",
	})
	err := sqb.Update(&c, conditions, map[string]interface{}{"description": "updated cat1 desc "})

	if err != nil {
		t.Errorf("Failed to update Cateogry: %v", err)
	}

}

func TestDeleteCategory(t *testing.T) {

	var c v1.InventoryCategory
	sqb := models.SimpleQueryBuilder{DB: global.GDB}

	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "id",
		Operator:   models.EQUALS,
		Value:      "1",
	})
	conditions = append(conditions, models.AND)
	conditions = append(conditions, models.Condition{
		ColumnName: "parent_id",
		Operator:   models.EQUALS,
		Value:      "3",
	})
	err := sqb.Delete(&c, conditions)

	if err != nil {
		t.Errorf("Failed to delete Cateogry: %v", err)
	}

}
