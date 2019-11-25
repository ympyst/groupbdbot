package main

import (
    "encoding/json"
    "fmt"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    contract "github.com/ympyst/groupbirthday/contract"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    tlg "groupbdbot/telegram"
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

    var replyKeyboard tlg.ReplyKeyboardMarkup

    switch upd.Message.Text {
    case "/start":
        responseMessageText = "Welcome to Group Birthday Bot!\nUse /show_groups to list your groups"
        break
    case "/show_groups":
        groupsReply := getGroupsByUsername(ctx, upd.Message.UserFrom.Username)
        for _, groupName := range groupsReply.Groups {
           responseMessageText += fmt.Sprintf("%s\n", groupName)
        }
    case "/select_group":
        responseMessageText = "Select group:"
        groupsReply := getGroupsByUsername(ctx, upd.Message.UserFrom.Username)
        replyKeyboard.Keyboard = make([][]tlg.KeyboardButton, len(groupsReply.Groups))
        for i := 0; i < len(groupsReply.Groups); i++ {
            replyKeyboard.Keyboard[i] = make([]tlg.KeyboardButton, 1)
            replyKeyboard.Keyboard[i][0].Text = groupsReply.Groups[i]
        }
        replyKeyboard.OneTimeKeyboard = true
        replyKeyboard.Selective = true
        replyKeyboard.ResizeKeyboard = true
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
    
    response := tlg.SendMessageResponse{
        Method:      "sendMessage",
        ChatId:      upd.Message.Chat.Id,
        Text:        responseMessageText,
        ReplyMarkup: replyKeyboard,
    }
    resBody, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), 500)
        fmt.Println("Error marshaling response JSON: " + err.Error())
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(resBody)
}

func getGroupsByUsername(ctx context.Context, username string) *contract.GetGroupsReply  {
    memberIdReply, err := groupBirthdayClient.GetMemberId(ctx, &contract.GetMemberIdRequest{TelegramUsername: username})
    if err != nil {
        panic(err)
    }
    groupsReply, err := groupBirthdayClient.GetGroups(ctx, &contract.GetGroupsRequest{
        MemberId: memberIdReply.MemberId,
    })
    if err != nil {
        panic(err)
    }
    return groupsReply
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

