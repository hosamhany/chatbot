package main
//
import (
	"fmt"
	"strings"
  "github.com/joho/godotenv"
  "os"
  "log"
	"github.com/nlopes/slack"
	"math/rand"
	"time"
	"strconv"
)
func main() {
	rand.Seed(time.Now().Unix())
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
          assignRandomType(input[1], rtm, ev, input)
          // rtm.SendMessage(rtm.NewOutgoingMessage(randType, ev.Channel))
					// rtm.SendMessage(rtm.NewOutgoingMessage(erza3(input), ev.Channel))
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

  func assignRandomType(text string, rtm *slack.RTM, ev *slack.MessageEvent, input []string){
    switch text {
    case "randomPairs":
			output:=shuffleAll(input[2:])
			selectRandomX(output, 2, rtm, ev)
    case "one":
			selectRandomOne(input, rtm, ev)
    case "all":
			output:=shuffleAll(input[2:])
			rtm.SendMessage(rtm.NewOutgoingMessage(strings.Join(output," "), ev.Channel))
		case "goGet":
			output:=shuffleAll(input[3:])
			groupCount, _ := strconv.Atoi(input[2])
			selectRandomX(output, groupCount, rtm, ev)

    default:
      rtm.SendMessage(rtm.NewOutgoingMessage("arza3", ev.Channel))

    }
  }
	func selectRandomOne(input []string,rtm *slack.RTM, ev *slack.MessageEvent){
		input = append(input[:0], input[2:]...)
		rtm.SendMessage(rtm.NewOutgoingMessage(input[rand.Intn(len(input))], ev.Channel))
	}

	func selectRandomX(input []string, groupCount int, rtm *slack.RTM, ev *slack.MessageEvent){
		for i:= range input{
			if(i%groupCount==0 && i< len(input)-groupCount+1){
				j := 0
				messageContent := ""
				for j < groupCount {
					messageContent += input[i+j] + " "
					j++
				}
				rtm.SendMessage(rtm.NewOutgoingMessage(messageContent, ev.Channel))
				if(i == len(input)-groupCount){
					break
				}
			} else if (i >= len(input)-groupCount+1) {
				rtm.SendMessage(rtm.NewOutgoingMessage(input[i], ev.Channel))
				}
		}
	}

	func shuffleAll(input []string) []string{
		for i:=range input{
			j:= rand.Intn(i+1)
			input[i], input[j] = input[j], input[i]
		}
		return input
		// return strings.Join(input, " ")
	}
