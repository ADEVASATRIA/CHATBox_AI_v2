package Configurations

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"../Utils"
)

func ConnectionDB() (*sql.DB, error) {
	err := godotenv("../.env")
	Utils.CheckError(2, "func ConnectionDB() -> Tidak dapat membaca file '../.env'.", err)

	DB_USER := os.Getenv("DB_USERNAME")
	DB_PASS := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	Utils.CheckError(2, "func ConnectionDB() -> Koneksi ke database gagal.", err)
}