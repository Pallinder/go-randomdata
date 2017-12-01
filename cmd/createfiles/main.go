package main

import (
	"flag"
	"fmt"
	"github.com/raff/go-randomdata"
	"math/rand"
	"os"
	"path"
)

func main() {
	nFolders := flag.Uint("folders", 0, "number of unique folders to create per level")
	nFiles := flag.Uint("files", 0, "number of unique files to create")
	depth := flag.Uint("depth", 0, "number of nested folders")
	total := flag.Uint("total", 0, "total number of files")
	create := flag.Bool("create", false, "create files and folders")
	templatedir := flag.String("templates", "", "path to templates folder")
	verbose := flag.Bool("verbose", false, "print created resource names")

	flag.Parse()

	extensions := flag.Args() // list of extensions to use

	if *nFiles == 0 {
		*nFiles = *total
	}

	if *depth == 0 {
		*depth = 1
	}

	templates := map[string]string{}

	if *templatedir != "" {
		d, err := os.Open(*templatedir)
		if err != nil {
			fmt.Println(err)
			return
		}

		files, err := d.Readdirnames(0)
		d.Close()

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, f := range files {
			ext := path.Ext(f)
			templates[ext] = path.Join(*templatedir, f)
		}
	}

	folders := make([]string, *nFolders)
	files := make([]string, *nFiles)

	for i := range folders {
		folders[i] = randomdata.FileName("") + "/"
	}

	for i := range files {
		files[i] = randomdata.FileName("")
	}

	for d := uint(0); d < *depth; d++ {
		for i := uint(0); i < *nFolders; i++ {
			j := d * i

			folders = append(folders, folders[j]+randomdata.FileName("")+"/")
		}
	}

	folders = append([]string{""}, folders...)

	if *create {
		for _, f := range folders[1:] {
			if *verbose {
				fmt.Println("mkdir", f)
			}
			if err := os.Mkdir(f, 0755); err != nil {
				fmt.Println(err)
			}
		}
	}

	for i := uint(0); i < *total; i++ {
		fi := rand.Intn(len(folders))
		fj := rand.Intn(len(files))

		ext := ""

		if len(extensions) > 0 {
			x := rand.Intn(len(extensions))
			ext = "." + extensions[x]
		}

		filename := folders[fi] + files[fj] + ext
		if *create {
			if template, ok := templates[ext]; ok {
				if *verbose {
					fmt.Println("create", filename)
				}
				if err := os.Link(template, filename); err != nil {
					fmt.Println("link", template, "-", err)
				}
			} else if f, err := os.Create(filename); err != nil {
				fmt.Println(err)
			} else {
				/*
					if template, ok := templates[ext]; ok {
						src, err := os.Open(template)
						if err != nil {
							fmt.Println(err)
						} else {
							_, err = io.Copy(f, src)
							if err != nil {
								fmt.Println("copy", template, "-", err)
							}
						}
					}
				*/
				f.Close()
			}
		} else {
			fmt.Println(filename)
		}
	}
}
