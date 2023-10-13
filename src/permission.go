package main

import (
	"fmt"
	"go_starter/database"
	"go_starter/logs"
	"go_starter/models"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a permission")
		return
	}
	permission := strings.ToLower(os.Args[1])
	postgresConnection, err := database.PostgresConnection()
	if err != nil {
		logs.Error(err)
		return
	}
	if permission == "all" {
		var permissions []models.Permission
		postgresConnection.Find(&permissions)
		fmt.Println("--------------------------------------")
		for _, p := range permissions {
			fmt.Println("name: " + p.Name + " - " + "id: " + strconv.Itoa(int(p.ID)))
		}
		fmt.Println("--------------------------------------")
	} else {
		createPermission := models.Permission{
			Name: permission,
		}
		err = postgresConnection.Create(&createPermission).Error
		if err != nil {
			logs.Error(err)
			return
		}
		fmt.Println("Create permission " + permission + " with id " + strconv.Itoa(int(createPermission.ID)) + " success")
		var permissions []models.Permission
		postgresConnection.Find(&permissions)
		fmt.Println("--------------------------------------")
		for _, p := range permissions {
			fmt.Println("name: " + p.Name + " - " + "id: " + strconv.Itoa(int(p.ID)))
		}
		fmt.Println("--------------------------------------")
	}

}
