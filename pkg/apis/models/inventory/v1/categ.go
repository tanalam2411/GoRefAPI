package v1

import (
	"gorm.io/gorm"
	"time"
)

// TableName - no need to add version as prefix for table name,
// as we migration will update the table and we will never have
// more than one table  by a single given name
func (b *InventoryCateg) TableName() string {
	return "inventory_categories"
}


type InventoryCateg struct {
	gorm.Model
	ID          uint   `json:"ID" gorm:"primaryKey;autoIncrement:true"`
	ParentID    uint   `json:"parent_id" gorm:"index:idx_parent;default:1"`
	Name        string `json:"name" binding:"required" gorm:"index:idx_name"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
