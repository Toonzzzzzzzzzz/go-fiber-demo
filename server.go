package main

import (
	"github.com/Toonzzzzzzzzzz/go-fiber-demo/database"
	"github.com/Toonzzzzzzzzzz/go-fiber-demo/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() //สร้าง fiber

	database.ConnectMongoDB() //ต่อ database

	routes.SetupRoutes(app) //ตั้งค่าเส้น api

	app.Listen(":3001") //รันบน port 3001
}
