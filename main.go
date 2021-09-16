package main

import (
	"fmt"
	"net/http"

	"github.com/tao12345666333/argo-cd-demo/pkg/hello"
	"github.com/tao12345666333/argo-cd-demo/pkg/version"
)

func info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("%s\n%s\n", hello.Greet(), version.Long()))
}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe(":8888", nil)
}
