package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
)

var requests = 0

func protected(w http.ResponseWriter, req *http.Request) {
	outputHeaders(req, w)
	requests++
	_, _ = fmt.Fprintf(w, "protected location %d!", requests)
}

func main() {
	http.HandleFunc("/protected/", protected)

	_ = http.ListenAndServe(":8080", nil)
}

func outputHeaders(req *http.Request, writer io.Writer) {
	var headers []string

	for header := range req.Header {
		headers = append(headers, header)
	}

	sort.Strings(headers)

	for _, header := range headers {
		for _, value := range req.Header[header] {
			_, _ = fmt.Fprintf(writer, "%s: %s\n", header, value)
		}
	}
}
