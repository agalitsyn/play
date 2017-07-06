```
$ fallocate -l 1G file.out
$ tar -czvf archive.tar.gz file.out
$ go run main.go archive.tar.gz
Regular file: file.out
```