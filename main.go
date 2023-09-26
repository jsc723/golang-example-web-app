package main

import (
	"fmt"
	"learn_go/todo_app/app/controllers"
	"learn_go/todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
