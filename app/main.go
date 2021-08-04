package main

import (
	"embed"
	"io/fs"
	"main/web"
	"net/http"
)

//go:embed dist/*
var reactContent embed.FS

func main() {
	dist, _ := fs.Sub(reactContent, "dist")
	var distFS = http.FS(dist)
	web.Serve(distFS, ":4200")
}
