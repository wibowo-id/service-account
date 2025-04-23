package route

import (
	"service-account/controller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	controller.Init(db)

	app.Post("/daftar", controller.RegisterNasabah)
	app.Post("/tabung", controller.Tabung)
	app.Post("/tarik", controller.Tarik)
	app.Get("/saldo/:no_rekening", controller.GetSaldo)
	app.Get("/mutasi/:no_rekening", controller.RiwayatMutasi) // tambahan
}
