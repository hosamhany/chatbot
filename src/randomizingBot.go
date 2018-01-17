package main
//
import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

func main() {

	token := "xoxb-299757536226-OaWKMHM8bncfFQKcuhXspG31"
	api := slack.New(token)
	rtm := api.NewRTM()
  // var input = []string
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
    return data[0]
  }
