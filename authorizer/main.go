package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

var requests = 0

func authorize(w http.ResponseWriter, req *http.Request) {
	outputHeaders(req, log.Writer())

	xUsers, ok := req.Header["X-Subject"]
	if !ok {
		log.Print("not authenticated")
		w.WriteHeader(401)
		return
	}

	if len(xUsers) != 1 {
		log.Print("too many users")
		w.WriteHeader(401)
		return
	}

	xUser := xUsers[0]

	if xUser != "user1" {
		log.Print("unknown user")
		w.WriteHeader(401)
	}

	requests++
	if requests % 2 == 0 {
		log.Print("allowed")
		w.WriteHeader(200)
	} else {
		log.Print("forbidden")
		w.WriteHeader(403)
	}
}

func main() {
	http.HandleFunc("/authorize", authorize)

	_ = http.ListenAndServe(":8081", nil)
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
