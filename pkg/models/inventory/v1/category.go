package v1

import (
	"gorm.io/gorm"
)

// InventoryCategory Model
type InventoryCategory struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	ParentID    uint   `json:"parentId" gorm:"index:idx_parent;default:1"`
	Name        string `json:"name" binding:"required" gorm:"index:idx_name"`
	Description string `json:"description"`
}

// TableName - no need to add version as prefix for table name,
// as we migration will update the table and we will never have
// more than one table  by a single given name
func (b *InventoryCategory) TableName() string {
	return "inventory_categories"
}

// InventoryCategory Model View
type InventoryCategoryView struct {
	ID          uint
	ParentID    uint
	Name        string
	Description string
}
