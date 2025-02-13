package main

import (
	"fmt"
	"net/http"
	"os"
)

func serveFile(w http.ResponseWriter, req *http.Request) {
    path := req.PathValue("path")

    if (path == "") {
        path = "index.html"
    }

    dat, err := os.ReadFile(fmt.Sprintf("/var/www/%s", path))

    if (err != nil) {
        w.WriteHeader(404)
        fmt.Fprintf(w, "404 Not Found")
        return
    }

    fmt.Fprintf(w, "%s", string(dat))
}

func main() {
    http.HandleFunc("GET /{path...}", serveFile)
    http.ListenAndServe(":80", nil)
}
