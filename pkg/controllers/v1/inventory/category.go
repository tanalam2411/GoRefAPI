package inventory

import (
	"github.com/gin-gonic/gin"
	"ms-inventory/global"
	"ms-inventory/pkg/models"
	modelsv1 "ms-inventory/pkg/models/inventory/v1"
	"net/http"
)

/*
Inventory Category Controller
*/

// Creates Inventory Categories
// Expects Category in request body
func CreateInventoryCategory(c *gin.Context) {
	var category modelsv1.InventoryCategory

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "ERROR",
			"message": "Validation Errors",
			"status":  400,
			"error":   err.Error(),
		})
		return
	}

	sqb := models.SimpleQueryBuilder{DB: global.GDB}
	if err := sqb.Create(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "ERROR",
			"message": "Record Error",
			"status":  400,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "OK",
		"id":      category.ID,
		"message": "Record Created",
		"status":  200,
	})
}

// Get Inventory Category
// Expects Category's Id to fetch Category from DB
func GetInventoryCategory(c *gin.Context) {
	var category modelsv1.InventoryCategory

	id := c.Params.ByName("id")

	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "id",
		Operator:   models.EQUALS,
		Value:      id,
	})

	sqb := models.SimpleQueryBuilder{DB: global.GDB}
	err := sqb.Get(&category, conditions)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "ERROR",
			"message": "Record Error",
			"status":  400,
			"error":   err.Error(),
		})
		return
	} else {
		apidata := modelsv1.InventoryCategoryView{
			ID:          category.ID,
			ParentID:    category.ParentID,
			Name:        category.Name,
			Description: category.Description,
		}
		c.JSON(http.StatusOK, gin.H{
			"code":        "OK",
			"page":        c.Query("page"),
			"page_size":   c.Query("page_size"),
			"results_max": 1,
			"status":      200,
			"data":        apidata,
		})
	}
}

// Update Inventory Categories
// Expects Category's Id in URI and Category fields to be updated in request body
func UpdateInventoryCategory(c *gin.Context) {
	var category modelsv1.InventoryCategory

	id := c.Params.ByName("id")

	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "id",
		Operator:   models.EQUALS,
		Value:      id,
	})
	sqb := models.SimpleQueryBuilder{DB: global.GDB}
	err := sqb.Get(&category, conditions)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "ERROR",
			"message": "Record Not found",
			"status":  404,
			"error":   err.Error(),
		})
		return
	}

	var updateCategory modelsv1.InventoryCategory
	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "ERROR",
			"message": "Validation Errors",
			"status":  400,
			"error":   err.Error(),
		})
		return
	}

	err = sqb.Update(&category, conditions, &updateCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "ERROR",
			"message": "Record Error",
			"status":  400,
			"error":   err.Error(),
		})
		return
	} else {
		apidata := modelsv1.InventoryCategoryView{
			ID:          category.ID,
			ParentID:    category.ParentID,
			Name:        category.Name,
			Description: category.Description,
		}
		c.JSON(http.StatusOK, gin.H{
			"code":   "OK",
			"status": 200,
			"data":   apidata,
		})
	}
}

// Delete Inventory Categories
// Expects Category's Id in URI to fetch and delete Category in DB
func DeleteInventoryCategory(c *gin.Context) {
	var category modelsv1.InventoryCategory
	id := c.Params.ByName("id")
	var conditions []interface{}
	conditions = append(conditions, models.Condition{
		ColumnName: "id",
		Operator:   models.EQUALS,
		Value:      id,
	})
	sqb := models.SimpleQueryBuilder{DB: global.GDB}
	err := sqb.Delete(&category, conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "ERROR",
			"message": "Record Error",
			"status":  400,
			"error":   err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    "OK",
			"id":      id,
			"message": "Record Deleted",
			"status":  200,
		})
	}
}
