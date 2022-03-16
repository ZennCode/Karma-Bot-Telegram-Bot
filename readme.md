# Karma-Bot by ZennCode [ Telegram Bot ]

The Karma-Bot is programmed in  GO and runs with an SQlite 3 database

## Install 

You need to install Go on your server to compile the binary

- https://go.dev/doc/install

run the following command.
```
go build .
```

## Config

First rename the following two files: 

`example_foo.db to foo.db`

`example_config.json to config.json`

You have now a config.json file in the main directory that looks like:

```
{
	"Token":"YOUR_TELEGRAM_TOKEN"	
}
```
Put your "YOUR_TELEGRAM_TOKEN" there.

If you don't know about Telegram Tokens, visit the official Telegram dokumentation:

- https://core.telegram.org/bots#6-botfather


## Usage

### First step

After adding the bot to your group chat, you need to register that chat to the database.

To do so, write the following command in the chat:
```
/add
```
### How to start the bot

On Windows:
```
./tgbot.exe {No Param here}
or double-click the tgbot.exe
```
On Linux:
```
./tgbot {No Param here}
```

### Commands
To upvote a post, reply to it with the command:
```
++
```
To downvote a post, reply to it with the command:
```
--
```
To view the leaderboard use the command:
```
/leaderboard 
or 
/lb
```
To check if the bot is running use the command:
```
/status
```
To list all commands in Telegram use the command:
```
/help
```
## Credits

- Hank for translating to English
- The GO Telegram API library https://github.com/go-telegram-bot-api/telegram-bot-api/v5
- mattn for the GO SQLite3 library https://github.com/mattn/go-sqlite3
