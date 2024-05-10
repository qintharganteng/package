package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	cek "github.com/qintharganteng/backn/module"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetAllAnggotaPerpustakaan(c *fiber.Ctx) error {
	anggota := cek.GetAllAnggotaPerpustakaan()
	return c.JSON(anggota)
}