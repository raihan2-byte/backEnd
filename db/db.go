package db

import (
	barang "blog/Barang"
	"blog/artikel"
	"blog/berita"
	endpointcount "blog/endpointCount"
	"blog/home"
	"blog/karyakmpf"
	"blog/merch"
	"blog/phototalk"
	"blog/shortvideo"
	"blog/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbName := os.Getenv("MYSQLDATABASE")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error", err)
	}

	err = db.AutoMigrate(&home.TagLineHome{}, &user.User{}, &endpointcount.Statistics{}, &berita.Berita{}, &barang.Barang{}, &phototalk.PhotoTalk{}, &karyakmpf.KMPF{}, &merch.Merch{}, &barang.Category{}, &berita.TagsBerita{}, &shortvideo.ShortVideo{}, &berita.KaryaBerita{}, &artikel.Artikel{}, &berita.BeritaImage{})
	if err != nil {
		log.Fatal("eror migration")
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}
