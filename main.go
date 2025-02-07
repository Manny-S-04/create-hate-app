package main

import (
	"flag"
	"fmt"
	"os"
)

type FileWritingStrategy interface {
	Execute(directory string, filename string, content string) error
}
type FileWriter struct {
	strategy FileWritingStrategy
}

func (fw *FileWriter) SetStrategy(strategy FileWritingStrategy) {
	fw.strategy = strategy
}

func (fw *FileWriter) ExecuteStrategy(directory, filename, content string) error {
	return fw.strategy.Execute(directory, filename, content)
}

type CreateGoFileStrategy struct{}

func (s *CreateGoFileStrategy) Execute(directory string, filename string, content string) error {
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return err
	}
	filePath := fmt.Sprintf("%s/%s", directory, filename)
	return os.WriteFile(filePath, []byte(content), 0644)
}
func main() {
	baseDir := flag.String("dir", ".", "Base directory to create the project structure")
	moduleName := flag.String("module", "hatestack", "Module name")
	flag.Parse()
	f := &file{}
	writer := &FileWriter{}
	commands := []struct {
		dir      string
		filename string
		content  string
	}{
		{"src/cmd", "main.go", f.MainGo()},
		{"src/cmd", "handlers.go", f.HandlersGo(*moduleName)},
		{"src/cmd", "database.go", f.DatabaseGo()},
		{"src/views", "embed.go", f.EmbedGo()},
		{"src/views/pages", "home.templ", f.HomeTempl()},
		{"src/views/pages", "home.templ.go", f.HomeTemplGo()},
		{"src/views/static/css", "styles.css", f.Styles()},
		{"src", "go.mod", f.ModFile(*moduleName)},
		{"src", "go.sum", f.SumFile()},
		{"src", "Makefile", f.Makefile()},
		{"src", ".env", "DB_URL=\nDB_PATH="},
	}

	for _, cmd := range commands {
		fullDir := fmt.Sprintf("%s/%s", *baseDir, cmd.dir)
		err := os.MkdirAll(fullDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", fullDir, err)
			continue
		}
		writer.SetStrategy(&CreateGoFileStrategy{})
		err = writer.ExecuteStrategy(fullDir, cmd.filename, cmd.content)
		if err != nil {
			fmt.Printf("Error creating file %s/%s: %v\n", fullDir, cmd.filename, err)
		} else {
			fmt.Printf("Successfully created %s/%s\n", fullDir, cmd.filename)
		}
	}
}
