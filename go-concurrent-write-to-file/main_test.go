package main

import (
	"log"
	"os"
	"runtime"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Logf("Processes: %v", runtime.GOMAXPROCS(0))

	f, err := tempFile(os.TempDir())
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("File: %v", f.Name())

	t.Run("concurrent write", func(t *testing.T) {
		tc := []struct {
			data []byte
			f    *os.File
		}{
			{data: []byte("test\n"), f: f},
			{data: []byte("test 2\n"), f: f},
			{data: []byte("test\n"), f: f},
			{data: []byte("test 2\n"), f: f},
		}
		for _, tt := range tc {
			tt := tt
			t.Run("", func(st *testing.T) {
				st.Parallel()

				err := write(tt.data, tt.f)
				if err != nil {
					t.Fatal(err)
				}
			})
		}
	})

	os.Remove(f.Name())
}
