package main

import (
	"embed"
	"fmt"
	"strings"
)

var (
	//go:embed all:files
	files embed.FS
)

type file struct {
}

// CMD FOLDER
func (f *file) MainGo() string {
	bytes, _ := files.ReadFile("files/maingo.txt")
	return string(bytes)
}

func (f *file) HandlersGo(module string) string {
	bytes, _ := files.ReadFile("files/handlersgo.txt")
	content := string(bytes)
	return strings.Replace(content, "{module}", fmt.Sprintf("%s", module), -1)
}

func (f *file) DatabaseGo() string {
	bytes, _ := files.ReadFile("files/databasego.txt")
	return string(bytes)
}

// CMD FOLDER

// VIEWS FOLDER
func (f *file) EmbedGo() string {
	bytes, _ := files.ReadFile("files/embedgo.txt")
	return string(bytes)
}

// xxxx// VIEWS/PAGES FOLDER
func (f *file) HomeTempl() string {
	bytes, _ := files.ReadFile("files/hometempl.txt")
	return string(bytes)
}
func (f *file) HomeTemplGo() string {
	bytes, _ := files.ReadFile("files/hometemplgo.txt")
	return string(bytes)
}

// xxxx// VIEWS/PAGES FOLDER
// xxxx// VIEWS/STATIC/CSS FOLDER
func (f *file) Styles() string {
	return ""
}

// xxxx// VIEWS/STATIC/CSS FOLDER
// VIEWS FOLDER

// SRC
func (f *file) ModFile(module string) string {
	bytes, _ := files.ReadFile("files/gomod.txt")
	content := string(bytes)
	return strings.Replace(content, "{module}", fmt.Sprintf("module %s", module), -1)
}
func (f *file) SumFile() string {
	bytes, _ := files.ReadFile("files/gosum.txt")
	return string(bytes)
}

func (f *file) Makefile() string {
	bytes, _ := files.ReadFile("files/Makefile.txt")
	return string(bytes)
}

// SRC
