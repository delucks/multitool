package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func filterFilename(name string) bool {
	parts := strings.Split(name, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}

type filteredFilename struct {
	http.File
}

func (f filteredFilename) Readdir(n int) (fis []os.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files {
		if !filterFilename(file.Name()) {
			fis = append(fis, file)
		}
	}
	return
}

type filteredFileSystem struct {
	http.FileSystem
}

func (fs filteredFileSystem) Open(name string) (http.File, error) {
	if filterFilename(name) { // If dot file, return 403 response
		return nil, os.ErrPermission
	}

	file, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return filteredFilename{file}, err
}

func ServeDirViaHTTP(args []string, _ io.Reader) error {
	var directory string
	if len(args) < 2 {
		directory = "."
	} else {
		directory = args[1]
	}
	fs := filteredFileSystem{http.Dir(directory)}
	fileHandler := http.FileServer(fs)
	http.Handle("/", fileHandler)
	fmt.Printf("Serving %s via port 8080\n", directory)
	return http.ListenAndServe(":8080", RequestLogger(fileHandler))
}
