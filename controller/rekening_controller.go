package controller

import (
	"errors"
	"service-account/helper"
	"service-account/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterNasabah(c *fiber.Ctx) error {
	type Request struct {
		Nama string `json:"nama"`
		NIK  string `json:"nik"`
		NoHP string `json:"no_hp"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		helper.LogError("Failed to parse register request", err)
		return c.Status(400).JSON(fiber.Map{"remark": "Invalid payload"})
	}

	var existing model.Nasabah
	if err := DB.Where("nik = ? OR no_hp = ?", req.NIK, req.NoHP).First(&existing).Error; err == nil {
		helper.LogWarning("NIK or NoHP already registered", map[string]string{
			"NIK":  req.NIK,
			"NoHP": req.NoHP,
		})
		return c.Status(400).JSON(fiber.Map{"remark": "NIK atau No HP sudah terdaftar"})
	}

	nasabah := model.Nasabah{Nama: req.Nama, NIK: req.NIK, NoHP: req.NoHP}
	DB.Create(&nasabah)

	rekening := model.Rekening{
		NoRekening: helper.GenerateNoRekening(req.NIK),
		Saldo:      0,
		NasabahID:  nasabah.ID,
	}
	DB.Create(&rekening)

	helper.LogInfo("Nasabah registered successfully", nasabah)

	return c.JSON(fiber.Map{"no_rekening": rekening.NoRekening})
}

func Tabung(c *fiber.Ctx) error {
	type Request struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		helper.LogError("Invalid tabung payload", err)
		return c.Status(400).JSON(fiber.Map{"remark": "Invalid payload"})
	}

	var rekening model.Rekening
	if err := DB.Where("no_rekening = ?", req.NoRekening).First(&rekening).Error; err != nil {
		helper.LogWarning("Tabung failed: rekening not found", req.NoRekening)
		return c.Status(400).JSON(fiber.Map{"remark": "Nomor rekening tidak ditemukan"})
	}

	rekening.Saldo += req.Nominal
	DB.Save(&rekening)

	mutasi := model.Mutasi{
		RekeningID: rekening.ID,
		Tipe:       "kredit",
		Nominal:    req.Nominal,
		Keterangan: "Setoran tunai",
	}
	DB.Create(&mutasi)

	helper.LogInfo("Tabung success", map[string]any{
		"rekening": rekening,
		"nominal":  req.Nominal,
	})

	return c.JSON(fiber.Map{"saldo": rekening.Saldo})
}

func Tarik(c *fiber.Ctx) error {
	type Request struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		helper.LogError("Invalid tarik payload", err)
		return c.Status(400).JSON(fiber.Map{"remark": "Invalid payload"})
	}

	var rekening model.Rekening
	if err := DB.Where("no_rekening = ?", req.NoRekening).First(&rekening).Error; err != nil {
		helper.LogWarning("Tarik failed: rekening not found", req.NoRekening)
		return c.Status(400).JSON(fiber.Map{"remark": "Nomor rekening tidak ditemukan"})
	}

	if rekening.Saldo < req.Nominal {
		helper.LogWarning("Tarik gagal: saldo tidak cukup", map[string]any{
			"rekening": rekening,
			"diminta":  req.Nominal,
		})
		return c.Status(400).JSON(fiber.Map{"remark": "Saldo tidak cukup"})
	}

	rekening.Saldo -= req.Nominal
	DB.Save(&rekening)

	mutasi := model.Mutasi{
		RekeningID: rekening.ID,
		Tipe:       "debit",
		Nominal:    req.Nominal,
		Keterangan: "Penarikan tunai",
	}
	DB.Create(&mutasi)
	helper.LogInfo("Tarik success", map[string]any{
		"rekening": rekening,
		"nominal":  req.Nominal,
	})

	return c.JSON(fiber.Map{"saldo": rekening.Saldo})
}

func GetSaldo(c *fiber.Ctx) error {
	noRek := c.Params("no_rekening")
	var rekening model.Rekening
	if err := DB.Where("no_rekening = ?", noRek).First(&rekening).Error; err != nil {
		helper.LogWarning("Get saldo gagal: rekening tidak ditemukan", noRek)
		return c.Status(400).JSON(fiber.Map{"remark": "Nomor rekening tidak ditemukan"})
	}
	helper.LogInfo("Get saldo success", rekening)

	return c.JSON(fiber.Map{"saldo": rekening.Saldo})
}

// tambahan
func RiwayatMutasi(c *fiber.Ctx) error {
	noRek := c.Params("no_rekening")

	var rekening model.Rekening
	if err := DB.Where("no_rekening = ?", noRek).First(&rekening).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.LogWarning("No rekening tidak ditemukan", noRek)
			return c.Status(400).JSON(fiber.Map{
				"remark": "Nomor rekening tidak ditemukan",
			})
		}
		helper.LogError("Gagal cek rekening", err)
		return c.Status(500).JSON(fiber.Map{
			"remark": "Terjadi kesalahan saat cek rekening",
		})
	}

	var mutasi []model.Mutasi

	err := DB.
		Preload("Rekening").
		Preload("Rekening.Nasabah").
		Where("rekenings.no_rekening = ?", noRek).
		Joins("INNER JOIN rekenings ON rekenings.id = mutasis.rekening_id").
		Order("mutasis.created_at desc").
		Find(&mutasi).Error

	if err != nil {
		helper.LogError("Gagal ambil mutasi", err)
		return c.Status(500).JSON(fiber.Map{"remark": "Gagal mengambil mutasi"})
	}
	helper.LogInfo("Get mutasi success by noRek", noRek)

	return c.JSON(fiber.Map{"mutasi": mutasi})
}
