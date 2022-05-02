package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const fakeSuffix = "-fake"

type Config struct {
	source string
	target string
}

func ParseConfig() *Config {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		fmt.Println("fz: wrong number of arguments")
		fmt.Println("Usage: fz <sourceFile> [<targetFile>]")
		return nil
	}

	source := arguments[0]
	var target string
	if len(arguments) > 1 {
		target = arguments[1]
	} else {
		name := filepath.Base(source)
		ext := filepath.Ext(source)
		parts := []string{
			strings.TrimSuffix(name, ext),
			fakeSuffix,
			ext,
		}
		target = strings.Join(parts, "")
	}

	return &Config{
		source,
		target,
	}
}

func main() {
	config := ParseConfig()
	if config != nil {
		runApp(config)
	}
	os.Exit(0)
}

func runApp(config *Config) int {
	reader, err := zip.OpenReader(config.source)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	targetFile, err := os.Create(config.target)
	if err != nil {
		log.Fatal(err)
	}
	defer targetFile.Close()

	zipWriter := zip.NewWriter(targetFile)
	defer zipWriter.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		f, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		f.Write([]byte{})
	}

	return 0
}
