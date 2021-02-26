package global

import "gorm.io/gorm"

// Global DB object, that can accessed across all packages
var (
	GDB *gorm.DB // DB Shared across the app
)
