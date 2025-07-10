package web

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

type A struct {
}

func (a *A) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello from a HandleFunc #2!\n")
}

func TestA(t *testing.T) {
	//str := []string{"login"}
	split := strings.Split("login", "/")
	split = split[1:]
	fmt.Println(split)
	http.HandleFunc("/u/a", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello from a HandleFunc #2!\n")
	})
	http.ListenAndServe(":80", &A{})
}
