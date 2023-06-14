package main

import (
	"fmt"

	"github.com/go-fiber-CRM/lead"
	"github.com/go-fiber-CRM/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App){

	app.Get("api/v1/lead",lead.GetLeads)
	app.Get("api/v1/lead/:id",lead.GetLead)
	app.Post("api/v1/lead",lead.NewLead)
	app.Delete("api/v1/lead/:id",lead.DeleteLead)
	
}

func initDatabase(){
	var err error
database.DBConn, err=gorm.Open("sqlite3","test.db")
if err!=nil{
	panic("failed to connect with db!")
}

fmt.Println("connection open to DB")
database.DBConn.AutoMigrate(&lead.Lead{})

}

func main(){
	app:=fiber.New()

	setupRoutes(app)

	initDatabase()

	app.Listen(3000)


	defer database.DBConn.Close()
}
