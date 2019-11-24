package main

import (
    "encoding/json"
    "fmt"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    tlg "groupbdbot/telegram"
    contract "github.com/ympyst/groupbirthday/contract"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

var groupBirthdayClient contract.GroupBirthdayClient

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

    responseMessageText := ""
    ctx := context.Background()

    switch upd.Message.Text {
    case "/start":
        responseMessageText = "Welcome to Group Birthday Bot!\nUse /show_groups to list your groupsReply"
        break
    case "/show_groups":
        memberIdReply, err := groupBirthdayClient.GetMemberId(ctx, &contract.GetMemberIdRequest{TelegramUsername: upd.Message.UserFrom.Username})
        if err != nil {
            panic(err)
        }
        groupsReply, err := groupBirthdayClient.GetGroups(ctx, &contract.GetGroupsRequest{
            MemberId: memberIdReply.MemberId,
        })
        if err != nil {
            panic(err)
        }
        for _, groupName := range groupsReply.Groups {
           responseMessageText += fmt.Sprintf("%s\n", groupName)
        }
    case "/list_birthdays":
        memberBirthdaysReply, err := groupBirthdayClient.GetMemberBirthdays(ctx, &contract.GetMemberBirthdaysRequest{GroupName: "family"})
        if err != nil {
            panic(err)
        }

        for _, birthday := range memberBirthdaysReply.MemberBirthdays {
            responseMessageText += fmt.Sprintf("%s %s %v.%v\n", birthday.FirstName, birthday.LastName, birthday.Day, birthday.Month)
        }
        break
    default:
        responseMessageText = "Unknown command"
    }

    response := tlg.SendMessageResponse{"sendMessage", upd.Message.Chat.Id, responseMessageText}
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

    serverAddr := os.Getenv("SERVER_ADDR")
    fmt.Println(serverAddr)
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithInsecure())
    conn, err := grpc.Dial(serverAddr, opts...)
    if err != nil {
        log.Fatalf("fail to dial: %v", err)
    }
    defer conn.Close()
    groupBirthdayClient = contract.NewGroupBirthdayClient(conn)

    port := os.Getenv("PORT")
    fmt.Printf("🤖 Now listening port %v\n", port)
    err = http.ListenAndServe(":" + port, nil)
    if err != nil {
        panic(err)
    }
}

