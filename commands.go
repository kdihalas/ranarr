package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/spf13/cobra"
	tb "gopkg.in/tucnak/telebot.v1"
)

func mainCmd(cmd *cobra.Command, args []string) {
	sonarr := &Sonarr{
		ApiKey: Sonarr_token,
		Url:    Sonarr_url,
	}

	radarr := &Radarr{
		ApiKey: Radarr_token,
		Url:    Radarr_url,
	}
	bot, err := tb.NewBot(Bot_token)
	if err != nil {
		log.Fatalln(err)
	}

	messages := make(chan tb.Message, 100)
	bot.Listen(messages, 10*time.Second)

	for message := range messages {

		findM, err := regexp.MatchString("/find movie .*", message.Text)
		if err != nil {
			fmt.Errorf("err: %s", err)
		}

		if message.Text == "/series" {
			text := "*List with series* \n"
			series := sonarr.GetSeries()
			for _, serie := range series {
				text = text + fmt.Sprintf("➸ %s \n", serie.Title)
			}
			bot.SendMessage(message.Chat, text, &tb.SendOptions{ReplyTo: message, ParseMode: tb.ModeMarkdown})
		}

		if message.Text == "/movies" {
			text := "*List with movies* \n"
			movies := radarr.GetMovies()
			for _, movie := range movies {
				text = text + fmt.Sprintf("➸ %s \n", movie.Title)
			}
			bot.SendMessage(message.Chat, text, &tb.SendOptions{ReplyTo: message, ParseMode: tb.ModeMarkdown})
		}

		if findM {
			fmt.Println("Got find request")
		}

	}
}
