package main

import (
    "encoding/json"
    "fmt"
    "time"
    "io/ioutil"
    "net/http"
    "os"
    tlg "groupbdbot/telegram"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "groupbdbot/groupdb"
)

func processUpdate(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()

    fmt.Print("\n➡ ️Received new request: ")

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        http.Error(w, err.Error(), 500)
        fmt.Println("Error reading request body: " + err.Error())
        return
    }

    var upd tlg.Update
    err = json.Unmarshal(body, &upd)
    if err != nil {
        http.Error(w, err.Error(), 500)
        fmt.Println("Error unmarshaling request JSON: " + err.Error())
        return
    }
    fmt.Printf("%+v\n", upd)

    db, err := gorm.Open("postgres", os.Getenv("DB_DSN"))
    defer db.Close()
    if err != nil {
        panic(err)
    }

    var members []groupdb.Member
    db.Where("group_id = ?", 1).Find(&members)
    fmt.Println(members)

    messageText := ""
    for _, member := range members  {
        bd, err := time.Parse(time.RFC3339, member.Birthday)
        if err != nil {
            panic(err)
        }
        month := bd.Month()
        day := bd.Day()
        messageText += fmt.Sprintf("%s %s %v.%v\n", member.FirstName, member.LastName, day, month)
    }

    response := tlg.SendMessageResponse{"sendMessage", upd.Message.Chat.Id, messageText}
    resBody, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), 500)
        fmt.Println("Error marshaling response JSON: " + err.Error())
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(resBody)
}

func main() {
    token := os.Getenv("TOKEN")
    http.HandleFunc("/" + token, processUpdate)

    port := os.Getenv("PORT")
    fmt.Printf("🤖 Now listening port %v\n", port)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        panic(err)
    }
}

