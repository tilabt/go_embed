package main

import (
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

func main() {
	fmt.Println(message)

	fmt.Printf("Image size: %d bytesn\n", len(logo))

	if err := writeFileWithDirs("dist/tilabt_copy.png", logo, 0644); err != nil {
		panic(err)
	}

}

func writeFileWithDirs(path string, data []byte, perm fs.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, data, perm)
}
