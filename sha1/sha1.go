package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
	sig1, err := sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig1)
}

/*
if file names ends with .gz
	$ cat http.log.gz | gunzip | openssl sha1
else
	$ cat http.log.gz | openssl sha1
*/

// cat http.log.gz | gunzip | openssl sha1
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close() // diferred are called in LIFO stack order
	var r io.Reader = file
	if strings.HasSuffix(fileName, ".gz") {
		file, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer file.Close()
		r = file
	}

	// io.CopyN(os.Stdout, r, 100)
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
