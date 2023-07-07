package main

import (
	"github.com/JuanPabloSerna/jps-alert-crud/bootstrap"
	"github.com/JuanPabloSerna/jps-alert-crud/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
