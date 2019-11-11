package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type Chat struct {
    Id int `json:"id"`
    Type string `json:"type"`
}

type Message struct {
    Id   int  `json:"message_id"`
    Timestamp int `json:"date"`
    Chat Chat `json:"chat"`
}

type Update struct {
    Id   int  `json:"update_id"`
    Message Message `json:"message"`
}

type SendMessageResponse struct {
    Method string `json:"method"`
    ChatId int `json:"chat_id"`
    Text string `json:"text"`
}

func processUpdate(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()

    fmt.Println("Received new request")
    fmt.Println(req.Header)

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        http.Error(w, err.Error(), 500)
        fmt.Println("Error reading request body: " + err.Error())
        return
    }

    var upd Update
    err = json.Unmarshal(body, &upd)
    if err != nil {
        http.Error(w, err.Error(), 500)
        fmt.Println("Error unmarshaling request JSON: " + err.Error())
        return
    }
    fmt.Println(upd)

    response := SendMessageResponse{"sendMessage", upd.Message.Chat.Id, "Hello"}
    resBody, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), 500)
        fmt.Println("Error marshaling response JSON: " + err.Error())
        return
    }

    fmt.Fprintf(w, string(resBody))
}

func main() {
    token := os.Getenv("TOKEN")
    http.HandleFunc("/" + token, processUpdate)

    port := os.Getenv("PORT")
    fmt.Printf("🔛 Now listening port %v\n", port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        panic(err)
    }
}

