package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/malifrh/template-postgres-go/config"
	"github.com/malifrh/template-postgres-go/services"
	"log"
	"os"
)

type app struct {
	AlbumService services.AlbumService
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	application := app{AlbumService: services.NewPostgresService(db)}

	albums, err := application.AlbumService.GetAllAlbum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all album : %v\n", albums)

	albumNo1, err := application.AlbumService.Get(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album number 1 : %v\n", albumNo1)

	err = application.AlbumService.BatchCreate([]services.Album{
		{Title: "Hari Yang Cerah", Artist: "Peterpan", Price: 50000},
		{Title: "Sebuah Nama Sebuah Cerita", Artist: "Peterpan", Price: 50000},
		{Title: "Bintang Di surga", Artist: "Peterpan", Price: 60000},
	})
	if err != nil {
		log.Fatal(err)
	}

	albumNo1.Price = 70000
	err = application.AlbumService.Update(*albumNo1)
	if err != nil {
		log.Fatal(err)
	}

	err = application.AlbumService.Delete(albumNo1.ID)
	if err != nil {
		log.Fatal(err)
	}
}
