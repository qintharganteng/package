package main

import (
	"log"

	"github.com/qintharganteng/package/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/qinthargantenf/package/url"

	"github.com/gofiber/fiber/v2"

)


func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}