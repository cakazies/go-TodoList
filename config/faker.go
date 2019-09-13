package config

import (
	"github.com/icrowley/fake"
	"github.com/local/TaskListGo/models"
)

func init() {
	limit := 100
	for i := 0; i < limit; i++ {
		todo := &models.Todo{}
		todo.Name = fake.FullName()
		todo.Priority = fake.Year(1900, 2010)
		todo.Priority = 0
		_ = todo.Create()
	}

}
