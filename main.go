package main

import (
	"fmt"
	"log"
	"net/http"
	"progate_crud_golang/controllers"
	"progate_crud_golang/models"

	"github.com/julienschmidt/httprouter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// connect ke database data.db
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//untuk otomatis membuat db

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		fmt.Println(err.Error())
	}
	// deklarasi router

	router := httprouter.New()

	todoController := &controllers.TodoController{}

	// http method GET AND POST
	router.GET("/", todoController.Index)
	router.GET("/create", todoController.Create)
	router.POST("/create", todoController.Create)
	router.GET("/edit/:id", todoController.Edit)
	router.POST("/update/:id", todoController.Update)
	router.POST("/done/:id", todoController.Done)
	// router.POST("/delete/:id", todoController.Delete)

	log.Println("server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
