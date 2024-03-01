package initializers

import (
	"github.com/sathish-30/GopherBlogServer/internal/database"
	"github.com/sathish-30/GopherBlogServer/internal/model"
)

func MigrateModel() {
	if err := database.DB.AutoMigrate(&model.User{}); err != nil {
		panic("Failed to migrate model in database")
	}
}
