package main

import (
	"log"
	"os"
	"runtime"
	"sync"
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
			{data: []byte("test\n"), f: f},
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

func TestWriteToProtectedFile(t *testing.T) {
	t.Logf("Processes: %v", runtime.GOMAXPROCS(0))

	f, err := tempFile(os.TempDir())
	if err != nil {
		log.Fatal(err)
	}

	mu := new(sync.Mutex)
	pf := protectedFile{
		f:  f,
		mu: mu,
	}

	t.Logf("File: %v", f.Name())

	t.Run("concurrent write with mutex", func(t *testing.T) {
		tc := []struct {
			data []byte
			f    protectedFile
		}{
			{data: []byte("test\n"), f: pf},
			{data: []byte("test\n"), f: pf},
		}
		for _, tt := range tc {
			tt := tt
			t.Run("", func(st *testing.T) {
				st.Parallel()

				err := tt.f.Write(tt.data)
				if err != nil {
					t.Fatal(err)
				}
			})
		}
	})

	os.Remove(f.Name())
}

func TestWriteToFileByChannel(t *testing.T) {
	t.Logf("Processes: %v", runtime.GOMAXPROCS(0))

	f, err := tempFile(os.TempDir())
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("File: %v", f.Name())

	queue := make(chan []byte)
	complete := make(chan bool)
	errors := make(chan error, 1)

	go writeFromChannel(queue, complete, errors, f)

	go func() {
		queue <- []byte("test\n")
		queue <- []byte("test\n")
		close(queue)

	}()

	go func() {
		for {
			t.Error(<-errors)
		}
	}()

	<-complete

	os.Remove(f.Name())
}
