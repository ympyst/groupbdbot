package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type Update struct {
    Id   int  `json:"update_id"`
    Message string `json:"message"`
}

func processUpdate(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    var upd Update
    err = json.Unmarshal(body, &upd)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    fmt.Println(upd)
    fmt.Fprintf(w, "Hello\n")
}

func main() {
    token := os.Getenv("TOKEN")
    http.HandleFunc("/" + token, processUpdate)

    port := os.Getenv("PORT")
    fmt.Printf("🔛 Now listening port %v", port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        panic(err)
    }
}

