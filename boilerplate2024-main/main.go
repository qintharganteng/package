package main

import (
	"log"

	"github.com/qintharganteng/package/boilerplate2024-main/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/qintharganteng/package/boilerplate2024-main/url"

	"github.com/gofiber/fiber/v2"

)


func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}

