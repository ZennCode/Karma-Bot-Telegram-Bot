package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	errorino "zenncode/tgbot/error"
	funcs "zenncode/tgbot/funcs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

type config struct {
	Token       string      `json:"Token"`
}

func main() {
	content, err := ioutil.ReadFile("config.json")
	errorino.FatalError(err)
	config := config{}
	err = json.Unmarshal(content, &config)
	errorino.FatalError(err)
	bot, err := tgbotapi.NewBotAPI(config.Token)
	errorino.PanicError(err)
    bot.Debug = false
	log.Printf("Authorized with bot account: %s with id: %d", bot.Self.UserName, bot.Self.ID)
	u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
            continue
        }
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// Extract the command from the Message.
		if strings.HasPrefix(update.Message.Text, "++") {
			uID := update.SentFrom().ID
			srnm := update.SentFrom().UserName
			name := update.Message.From.FirstName +" "+update.Message.From.LastName
			fromName := update.SentFrom().FirstName +" "+ update.SentFrom().LastName
			serverID := update.FromChat().ID
 			if update.Message.ReplyToMessage == nil { // ignore any non-Message updates
				continue
			} 
 			if update.Message.ReplyToMessage.From.IsBot { // ignore any non-ReplyToMessage updates
				msg.Text = "While I appreciate that, I don't need any Karma. Because maschines do not reincarnate."
				if _, err := bot.Send(msg); err != nil {
					errorino.PanicError(err)
				}
				continue
			}  
			real_uID := update.Message.ReplyToMessage.From.ID
			real_srnm := update.Message.ReplyToMessage.From.UserName
			real_name := update.Message.ReplyToMessage.From.FirstName +" "+update.Message.ReplyToMessage.From.LastName
			var messageID int64 = int64(update.Message.ReplyToMessage.MessageID)
 			if real_uID == uID {
				msg.Text = "You can't give Karma to yourself"
				if _, err := bot.Send(msg); err != nil {
					errorino.PanicError(err)
				}
				continue
			}
			c1 := funcs.CheckMention(uID, srnm, messageID)
			if c1 == messageID {
				reply := "❌ You have already rated this post."
				msg.Text = reply
				if _, err := bot.Send(msg); err != nil {
					errorino.PanicError(err)
				}
				continue
			}
			funcs.CheckIfUserExists(uID, srnm, messageID,name,serverID)
			funcs.CheckIfUserExists(real_uID, real_srnm, messageID,real_name,serverID)
			funcs.UpdateUserName(real_uID,messageID, real_srnm,serverID)
			funcs.UpdateName(uID,messageID, real_srnm,name,serverID)
			funcs.PlusEins(real_uID, real_srnm, messageID,serverID)
			funcs.UpdateLastMention(uID, messageID, srnm)
			reply := funcs.RandomHallo()+" "+real_name+" you received +1 Karma from "+fromName+". You now have "+strconv.Itoa(funcs.GetPoints(real_uID, real_srnm, messageID,serverID))+" Karma."
			msg.Text = reply
			if _, err := bot.Send(msg); err != nil {
				errorino.PanicError(err)
			}
		}
		if strings.HasPrefix(update.Message.Text, "—") || strings.HasPrefix(update.Message.Text, "--") {
			uID := update.SentFrom().ID
			srnm := update.SentFrom().UserName
			name := update.Message.From.FirstName +" "+update.Message.From.LastName
			fromName := update.SentFrom().FirstName +" "+ update.SentFrom().LastName
			serverID := update.FromChat().ID
 			if update.Message.ReplyToMessage == nil { // ignore any non-Message updates
				continue
			} 
			if update.Message.ReplyToMessage.From.IsBot  { // ignore any non-ReplyToMessage updates
				msg.Text = "HA, nice try! You can't take Karma away from me"
				if _, err := bot.Send(msg); err != nil {
					errorino.PanicError(err)
				}
				continue
			}  
			real_uID := update.Message.ReplyToMessage.From.ID
			real_srnm := update.Message.ReplyToMessage.From.UserName
			real_name := update.Message.ReplyToMessage.From.FirstName +" "+update.Message.ReplyToMessage.From.LastName
			var messageID int64 = int64(update.Message.ReplyToMessage.MessageID)
			if real_uID == uID {
				msg.Text = "You can't take away Karma from yourself"
				if _, err := bot.Send(msg); err != nil {
					errorino.PanicError(err)
				}
				continue
			}
			c1 := funcs.CheckMention(uID, srnm, messageID)
			if c1 == messageID {
				reply := "❌ You have already rated this post."
				msg.Text = reply
				if _, err := bot.Send(msg); err != nil {
					errorino.PanicError(err)
				}
				continue
			}
			funcs.CheckIfUserExists(uID, srnm, messageID,name,serverID)
			funcs.CheckIfUserExists(real_uID, real_srnm, messageID,real_name,serverID)
			funcs.UpdateUserName(real_uID,messageID, real_srnm,serverID)
			funcs.UpdateName(real_uID,messageID, real_srnm,name,serverID)
			funcs.MinusEins(real_uID, real_srnm, messageID,serverID)
			funcs.UpdateLastMention(uID, messageID, srnm)
			reply := funcs.RandomHallo()+" "+real_name+ ", " +fromName+" has substracted 1 Karma from you. You have "+strconv.Itoa(funcs.GetPoints(real_uID, real_srnm, messageID,serverID))+" Karma left."

			msg.Text = reply
			if _, err := bot.Send(msg); err != nil {
				errorino.PanicError(err)
			}
		}
		switch update.Message.Command() {
			case "help":
				msg.Text = "Use /add to register your chat do the database \nI understand the following commands: \n++ or -- as reply to a post to up- or downvote it \n/leaderboard or /lb for the Leaderboard \n/status to check if the bot is running"
			case "leaderboard", "lb":
				serverID := update.FromChat().ID
				uID := update.SentFrom().ID
				srnm := update.SentFrom().UserName
				lb := funcs.LeaderBoard(uID,srnm,serverID)
				// fmt.Println(lb)
				msg.Text = lb
			case "status":
				msg.Text = "I'm ok."
			case "add":
				funcs.NewTable(update.FromChat().ID)
				
				msg.Text = "This chat has successfully been registered to the database."
			case "test":
				bg:= funcs.RandomHallo()
				msg.Text = bg
			default:
				continue
		}
		if _, err := bot.Send(msg); err != nil {
			errorino.PanicError(err)
		}
	}
}