package helper

import (
	"archive/zip"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// creating zip using archive/zip package
func CreateZip(files []string, location string) (string, error) {
	cache := location
	var zipname string
	if len(files) == 1 {
		zipname, _ = filepath.Abs(location + "/" + files[0] + ".zip")
	} else {
		// create zip name using half of each file name
		zipname = location + "/"
		for _, file := range files {
			div := math.Abs(float64(len(file) / 2))
			zipname += file[:int(div)]
		}
		zipname, _ = filepath.Abs(zipname + ".zip")
	}
	nzip, err := os.Create(zipname)
	if err != nil {
		return "", err
	}
	defer nzip.Close()
	zipWriter := zip.NewWriter(nzip)
	defer zipWriter.Close()

	// walk the cache directory and skip files whose names are not in the files array
	// return nil if dir else write the file to the zip archive
	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// zip only specified files
		if !StringInArray(files, path, true) {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// path = strings.TrimPrefix(path, "/")
		path = strings.TrimPrefix(path, cache)
		f, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}

	err = filepath.Walk(cache, walker)
	if err != nil {
		return "", err
	}

	// return path to zip file
	abszip, _ := filepath.Abs(zipname)
	return abszip, nil
}
