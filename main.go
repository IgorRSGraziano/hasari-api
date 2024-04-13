package main

import (
	"hasari-api/models"
	"hasari-api/router"
)

func main() {
	models.ConnectDatabase()
	router.Init()
}
