package migrate

import (
	"github.com/jenniekibiri/crud-golang/configs"
	"github.com/jenniekibiri/crud-golang/models"
)

func Migrate() {
	configs.DB.AutoMigrate(&models.Person{})
}
