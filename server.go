package main

import (
    "fmt"
    "net/http"
)

func processUpdate(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "Hello\n")
}

func main() {

    http.HandleFunc("/803026579:AAEVPiHB9B3c5V63vvIwkLCFQZC68h5wTZo", processUpdate)

    http.ListenAndServe(":8090", nil)
}

