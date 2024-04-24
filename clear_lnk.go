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

func main() {
	log.Println("start clean lnk")

	exts := []string{
		".bin",
		".lnk",
		".qwe",
	}
	files := find("E:\\", exts)

	for _, i := range files {
		err := os.Remove(i)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Remove %s\n", i)
	}

	log.Println("Done!")
	fmt.Scanln()
}
