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
	dist, err := fs.Sub(reactContent, "dist")
	if err != nil {
		panic(err)
	}
	var distFS = http.FS(dist)
	web.Serve(distFS, ":4200")
}
