package main

import (
	"embed"
	"io/fs"
	"log"

	"github.com/Jiang-Gianni/gosqlx/server"
	"github.com/joho/godotenv"
)

//go:embed all:assets
var assetsFs embed.FS

func main() {
	// loading ".env" file
	godotenv.Load()
	fsys, err := fs.Sub(assetsFs, "assets")
	if err != nil {
		log.Fatal(err)
	}
	server.RegisterAndRun(fsys)
}
