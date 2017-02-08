package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	f, err := tempFile(os.TempDir())
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("test\n")
	err = write(data, f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("content written to", f.Name())
}

// see BUGS here http://man7.org/linux/man-pages/man2/write.2.html
// seems that we don't need to do sync on app level, because kernel did it for us
//
// But POSIX says:
//		This volume of POSIX.1-2008 does not specify behavior of concurrent writes to a file from multiple processes. Applications should use some form of concurrency control.
// ...which means that you're on your own - different UNIX-likes will give different guarantees.
//
// But.. see https://www.cs.helsinki.fi/linux/linux-kernel/2002-30/1396.html
// ...which means if you need an _order_ in concurrent write you _should_ use O_APPEND
//
// But Linux wrote
//		I still consider any program relying on this behaviour buggy. Your
//		"atomic" write is an illusion, the os can certainly try, but in the end
//		it's the applications responsibility to get this right.
// ... which means you might want to do sync for guarantees
func write(data []byte, f *os.File) error {
	_, err := f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func tempFile(dir string) (*os.File, error) {
	f, err := os.OpenFile(filepath.Join(dir, "test"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}
