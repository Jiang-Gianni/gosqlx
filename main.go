package main

import (
	"embed"
	"os"

	"github.com/Jiang-Gianni/gosqlx/server"
)

//go:embed all:assets
var assetsFs embed.FS

func main() {
	fsys := os.DirFS("assets")
	// fsys, err := fs.Sub(assetsFs, "assets")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	server.RegisterAndRun(fsys)
}
