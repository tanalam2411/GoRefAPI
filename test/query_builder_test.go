package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ms-inventory/config"
	"ms-inventory/global"
	"ms-inventory/initialize"
	"ms-inventory/pkg/models"
	v1 "ms-inventory/pkg/models/inventory/v1"
	"net/http"
	"net/url"
	"testing"
)

func init() {
	global.GDB = initialize.GormMysql()
}

func TestCreateCategory(t *testing.T) {

	c := v1.InventoryCategory{
		ParentID:    3,
		Name:        "Category 1",
		Description: "Cat1 desc",
	}

	global.GDB = initialize.GormMysql()
	sqb := models.SimpleQueryBuilder{DB: global.GDB}

	err := sqb.Create(&c)

	if err != nil {
		t.Errorf("Failed to create Category: %v", err)
	}
}

func TestGetCategory(t *testing.T) {

	var c v1.InventoryCategory
	sqb := models.SimpleQueryBuilder{DB: global.GDB}

	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "name",
		Operator:   models.EQUALS,
		Value:      "Category 1",
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
		ColumnName: "name",
		Operator:   models.EQUALS,
		Value:      "Category 1",
	})
	conditions = append(conditions, models.AND)
	conditions = append(conditions, models.Condition{
		ColumnName: "parent_id",
		Operator:   models.EQUALS,
		Value:      "3",
	})

	var updateCategory = v1.InventoryCategory{Description: "updated cat1 desc "}
	err := sqb.Update(&c, conditions, &updateCategory)

	if err != nil {
		t.Errorf("Failed to update Cateogry: %v", err)
	}

}

func TestGetAllCategories(t *testing.T) {

	var c v1.InventoryCategory
	sqb := models.SimpleQueryBuilder{DB: global.GDB}

	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "name",
		Operator:   models.EQUALS,
		Value:      "Category 1",
	})
	conditions = append(conditions, models.AND)
	conditions = append(conditions, models.Condition{
		ColumnName: "parent_id",
		Operator:   models.EQUALS,
		Value:      "3",
	})

	res := sqb.GetAll(&c, conditions)
	var inventorycategory []v1.InventoryCategory

	gc := &gin.Context{Request: &http.Request{URL: &url.URL{}}}
	result := res.Limit(1000).Scopes(config.Paginate(gc)).Find(&inventorycategory)

	if result.Error != nil {
		t.Errorf("Failed to get all Cateogry: %v", result.Error)
	}

}

func TestDeleteCategory(t *testing.T) {

	var c v1.InventoryCategory
	sqb := models.SimpleQueryBuilder{DB: global.GDB}

	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "name",
		Operator:   models.EQUALS,
		Value:      "Category 1",
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
