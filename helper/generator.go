package helper

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateNoRekening(nik string) string {
	now := time.Now()
	mmyy := now.Format("0106") // MMYY

	// Ambil 4 digit terakhir dari NIK
	last4 := "0000"
	if len(nik) >= 4 {
		last4 = nik[len(nik)-4:]
	}

	// Generate 8 digit angka random
	rand.Seed(uint64(time.Now().UnixNano()))
	random8 := rand.Intn(1e8) // nilai antara 0 - 99999999

	return fmt.Sprintf("%s%s%08d", mmyy, last4, random8)
}
