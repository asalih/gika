package main

import (
	"fmt"
	"path/filepath"

	"github.com/asalih/gika"
	gkh "github.com/asalih/gika/handlers"
)

func main() {

	// const text = "Text content zlib\n"
	// var buf bytes.Buffer

	// // compress text
	// w := zlib.NewWriter(&buf)
	// // if err != nil {
	// // 	log.Fatalf("xz.NewWriter error %s", err)
	// // }
	// if _, err := io.WriteString(w, text); err != nil {
	// 	log.Fatalf("WriteString error %s", err)
	// }
	// if err := w.Close(); err != nil {
	// 	log.Fatalf("w.Close error %s", err)
	// }

	// fs, err := os.OpenFile("testdata/test.zlib", os.O_CREATE|os.O_RDWR, 0644)
	// if err != nil {
	// 	log.Fatalf("OpenFile error %s", err)
	// }
	// defer fs.Close()
	// fs.Write(buf.Bytes())

	files, err := filepath.Glob("testdata/archives/*.json")
	if err != nil {
		panic(err)
	}

	gk := gika.New(&gkh.AutoDetectContentHandler{})

	for _, file := range files {
		if file == "testdata/.DS_Store" {
			continue
		}

		handle, err := gk.WithPath(file)
		if err != nil {
			panic(err)
		}

		entries, err := handle.Read()
		if err != nil {
			panic(err)
		}

		handle.Close()

		fmt.Println(entries)
	}
}
