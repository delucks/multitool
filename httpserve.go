package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		referer := r.Referer()
		if referer == "" {
			referer = "-"
		}

		// Apache combined log format, but since we don't have access to content-length it's the request timing instead
		fmt.Printf(
			"%s - %s [%s] \"%s %s %s\" %d %v %s \"%s\"\n",
			r.RemoteAddr,
			"-", // user assumed to be empty because we don't have authentication
			start.Format("02/Jan/2006:15:04:05 -0700"),
			r.Method,
			r.RequestURI,
			r.Proto,
			200, // we don't have access to the Response, so we assume everything is a 200
			time.Since(start),
			referer,
			r.UserAgent(),
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
	port := "8080"
	switch len(args) {
	default:
		directory = "."
	case 2:
		ok, err := pathIsDir(args[1])
		if err != nil {
			return err
		}
		if ok {
			directory = args[1]
		} else {
			return errors.New(args[1] + " is not a valid directory")
		}
	case 3:
		ok, err := pathIsDir(args[1])
		if err != nil {
			return err
		}
		if ok {
			directory = args[1]
		} else {
			return errors.New(args[1] + " is not a valid directory")
		}
		_, err = strconv.Atoi(args[2])
		if err != nil {
			return err
		}
		port = args[2]
	}
	fs := filteredFileSystem{http.Dir(directory)}
	fileHandler := http.FileServer(fs)
	http.Handle("/", fileHandler)
	fmt.Fprintf(os.Stderr, "Serving %s on http://localhost:%s. Combined log format to stdout:\n", directory, port)
	return http.ListenAndServe(":"+port, RequestLogger(fileHandler))
}
