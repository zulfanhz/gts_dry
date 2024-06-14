package util

import (
	"gts-dry/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateJWT(id string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	})
	return token.SignedString(secret)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseNullableTime(nt string) time.Time {

	t, err := time.Parse("2006-01-02 15:04:05.999999", nt)
	if err != nil {
		return time.Time{}
	}
	return t

}

func CalculateStock(stock float64, satuan []model.BarangSatuanModel) []model.BarangSatuanModel {
	// Inisialisasi stok untuk satuan pertama
	satuan[0].Stock = stock

	// Loop untuk menghitung stok untuk satuan berikutnya
	for i := 1; i < len(satuan); i++ {
		satuan[i].Stock = satuan[i-1].Stock / satuan[i].Qty
	}

	return satuan
}

func CalculateRemainingExp(exp string) (int, string) {
	expDate, err := time.Parse("2006-01-02", exp)
	if err != nil {
		return 0, "Invalid date format"
	}

	now := time.Now()
	nowDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	daysRemaining := int(expDate.Sub(nowDate).Hours() / 24)

	var status string
	switch {
	case daysRemaining >= 14:
		status = "AMAN"
	case daysRemaining >= 7:
		status = "SEGERA KIRIM"
	case daysRemaining == 0:
		status = "DILARANG KIRIM"
	case daysRemaining < 0:
		status = "DILARANG KIRIM (EXPIRED)"
	default:
		status = "TIDAK AMAN"
	}

	return daysRemaining, status
}
