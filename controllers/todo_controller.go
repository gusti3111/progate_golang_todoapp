package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"progate_crud_golang/models"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

//buat struct database
type TodoController struct{}

// controller index

func (controller *TodoController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// connect to databse

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/index.html",
	}
	// parsing files

	htmltemp, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	}
	var todos []models.Todo
	db.Find(&todos)

	datas := map[string]interface{}{
		"Todos": todos,
	}

	err = htmltemp.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())

	}

}

// controller create

func (controller *TodoController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println()
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// jika methodnya POST
	if r.Method == "POST" {
		todo := models.Todo{
			Name:    r.FormValue("name"),
			Content: r.FormValue("content"),
			NIK:     r.FormValue("nik"),
			Date:    r.FormValue("date"),
		}
		// create db

		result := db.Create(&todo)
		if result.Error != nil {
			log.Println(result.Error)
			fmt.Println(result.Error)

			return

		}
		http.Redirect(w, r, "/", http.StatusFound)

		// IF GET METHOD
	} else {
		files := []string{
			"./views/base.html",
			"./views/create.html",
		}
		htmltemp, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return

		}
		err = htmltemp.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())

		}

	}

}

// controller Edit
func (controller *TodoController) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// connect to database sqlite
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
		noteId, _ := strconv.Atoi(params.ByName("id"))
		var note models.Todo
		db.Where("ID = ?", noteId).First(&note)
		note.Content = r.FormValue("content")
		note.Date = r.FormValue("deadline")
		note.NIK = r.FormValue("nik")
		note.Name = r.FormValue("name")

		db.Save(&note)
		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		files := []string{
			"./views/base.html",
			"./views/edit.html",
		}
		// parsing files
		htmltemp, err := template.ParseFiles(files...)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return

		}
		// GET by ID
		var todo models.Todo
		db.Find(&todo, params.ByName("id"))

		datas := map[string]interface{}{
			"Todo": todo,
			"ID":   params.ByName("id"),
		}

		// execute tamplate
		err = htmltemp.ExecuteTemplate(w, "base", datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
			return

		}

	}
}

// controller Update

func (controller *TodoController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	noteId, _ := strconv.Atoi(params.ByName("id"))

	var note models.Todo
	// perintah Update
	db.Where("ID = ?", noteId).First(&note)
	note.Name = r.FormValue("name")
	note.Date = r.FormValue("date")
	note.NIK = r.FormValue("nik")
	note.Content = r.FormValue("content")

	db.Save(&note)

	http.Redirect(w, r, "/", http.StatusFound)

}

// controller Done

func (controller *TodoController) Done(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())

	}

	var todo models.Todo
	// Perintah Done

	db.Find(&todo, params.ByName("id"))
	todo.IsDone = !todo.IsDone
	// save todo yg merupakan model database
	db.Save(&todo)
	// redirect http
	http.Redirect(w, r, "/", http.StatusFound)

}

func (controller *TodoController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// koneksikan ke database

	db, err := gorm.Open(sqlite.Open("data,db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())

	}

	var todos models.Todo
	// perintah Delete

	db.Unscoped().Delete(&todos, params.ByName("id"))

	http.Redirect(w, r, "/", http.StatusFound)

}
