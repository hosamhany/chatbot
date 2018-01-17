package main
//
import (
	"fmt"
	"strings"
  "github.com/joho/godotenv"
  "os"
  "log"
	"github.com/nlopes/slack"
)

func main() {
  err := godotenv.Load()
 if err != nil {
   log.Fatal("Error loading .env file")
 }
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
  // var input = []string
  // var RandType = ""
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)
				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
          fmt.Println(ev.User)
          var input = addIntoArray(ev.Text)
					rtm.SendMessage(rtm.NewOutgoingMessage(erza3(input), ev.Channel))
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}
  func addIntoArray(text string) []string{
    return strings.Split(text, " ")
  }


  func erza3(data []string)string{
    // RandType :=  data[1]
    return data[0]
  }
