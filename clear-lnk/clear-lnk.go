package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func find(root string, exts []string) []string {
	var result []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		for _, ext := range exts {
			if filepath.Ext(d.Name()) == ext {
				result = append(result, s)
			}
		}
		return nil
	})
	return result
}

func getSupportedExtensions() []string {
	return []string{".bin", ".lnk", ".qwe"}
}

func main() {
	log.Println("start clean lnk")

	files := find("E:\\", getSupportedExtensions())

	for _, f := range files {
		err := os.Remove(f)
		if err != nil {
			log.Printf("Error remove file %s", f)
			log.Fatal(err)
		}
		log.Printf("Remove %s\n", f)
	}

	log.Println("Done!")
	fmt.Scanln()
}
