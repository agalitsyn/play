package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("no args was provided")
	}
	filepath := args[0]
	// TODO: check that file exists and regular
	if err := readArch(filepath); err != nil {
		log.Fatalf("could not read archive: %v", err)
	}
}

func readArch(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzf)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			fmt.Println("Directory:", header.Name)
		case tar.TypeReg:
			fmt.Println("Regular file:", header.Name)
		default:
			fmt.Printf("Unable to figure out type: %c in file %s\n", header.Typeflag, header.Name)
		}
	}

	return nil
}
