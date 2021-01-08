package main

import (
	"flag"
	"log"
	"net/http"
	
)

func main() {
	addr := flag.String("addr", "127.0.0.1:49999", "Address to listen on")
	flag.Parse()
	
	log.Printf("About to listen on %v\n", *addr)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(*addr, nil)
}
