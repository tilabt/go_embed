package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// must be create outside function
//
//go:embed message.txt
var message string

//go:embed tilabt.png
var logo []byte

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(message)

	fmt.Printf("Image size: %d bytesn\n", len(logo))

	// if err := writeFileWithDirs("dist/tilabt_copy.png", logo, 0644); err != nil {
	// 	panic(err)
	// }

	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	dirs, _ := path.ReadDir("files")
	for _, entry := range dirs {
		if !entry.IsDir() {
			dirFile := fmt.Sprintf("files/%s", entry.Name())
			file, _ := path.ReadFile(dirFile)

			content := fmt.Sprintf("Content from %s => %s", dirFile, string(file))
			fmt.Println(content)
		}
	}

}

func writeFileWithDirs(path string, data []byte, perm fs.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, data, perm)
}
