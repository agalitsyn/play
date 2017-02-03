package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	t.Run("group", func(t *testing.T) {
		tc := []struct {
			dir string
		}{
			{tmpDir},
			{tmpDir},
		}
		for _, tt := range tc {
			tt := tt
			t.Run("", func(st *testing.T) {
				st.Parallel()

				_, err := createFile(tt.dir)
				if err != nil {
					t.Fatal(err)
				}
			})
		}

	})
	os.RemoveAll(tmpDir)
}
