package store

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/l-orlov/simple-todo-list/server/internal/model"
)

// Column names for models
var (
	asteriskTasks, asteriskUsers string
)

func init() {
	// Init column names for models
	asteriskTasks = asterisk(model.Task{})
	asteriskUsers = asterisk(model.User{})
}

type tableNameGetter interface {
	DbTable() string
}

// asterisk replace * in queries select(*) by column names (only for models without nesting)
func asterisk(a tableNameGetter) string {
	modelType := reflect.TypeOf(a)
	var columns []string
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		columnName, ok := field.Tag.Lookup("db")
		if !ok || columnName == "-" {
			continue
		}
		columns = append(columns, fmt.Sprintf("%s.%s", a.DbTable(), columnName))
	}
	return strings.Join(columns, ", ")
}
