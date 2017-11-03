package main

import (
	"fmt"
	"regexp"
	"time"
	"strings"
	"strconv"

	"github.com/spf13/cobra"
	tb "gopkg.in/tucnak/telebot.v1"
	log "github.com/sirupsen/logrus"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func mainCmd(cmd *cobra.Command, args []string) {
	var authUsers []string
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
		log.Infof("Got message: %s", message.Text)

		auth, _ := regexp.MatchString("/auth .*", message.Text)
		findMovie, _ := regexp.MatchString("/find movie .*", message.Text)
		findSerie, _ := regexp.MatchString("/find series .*", message.Text)
		downloadSerie, _ := regexp.MatchString("/download_series_.*", message.Text)
		downloadMovie, _ := regexp.MatchString("/download_movie_.*", message.Text)

		chatId := strconv.Itoa(int(message.Chat.ID))

		if auth {
			password := strings.TrimPrefix(message.Text, "/auth ")
			if password == Password {
				if !stringInSlice(chatId, authUsers) {
					authUsers = append(authUsers, chatId)
					bot.SendMessage(message.Chat, "You are now authenticated!", &tb.SendOptions{ParseMode: tb.ModeHTML})
				}
			}
		}

		if stringInSlice(chatId, authUsers) {
			if message.Text == "/series" {
				log.Info("Matched series list handler")
				text := "<b>List with series</b> \n"
				series := sonarr.GetSeries()
				for _, serie := range series {
					text = text + fmt.Sprintf("➸ %s \n", serie.Title)
				}
				bot.SendMessage(message.Chat, text, &tb.SendOptions{ParseMode: tb.ModeHTML})
			}

			if message.Text == "/movies" {
				log.Info("Matched movies list handler")
				text := "<b>List with movies</b> \n"
				movies := radarr.GetMovies()
				for _, movie := range movies {
					text = text + fmt.Sprintf("➸ %s \n", movie.Title)
				}
				bot.SendMessage(message.Chat, text, &tb.SendOptions{ParseMode: tb.ModeHTML})
			}

			if findMovie {
				log.Info("Matched find movie handler")
				match := strings.TrimPrefix(message.Text, "/find movie ")
				movies := radarr.SearchForMovies(match)
				text := fmt.Sprintf("<b>Movies found %d</b> \n", len(movies))
				for _, movie := range movies {
					text = text + fmt.Sprintf("/download_movie_%d ", movie.TmdbId)
					text = text + fmt.Sprintf("➸ %s (%d) \n", movie.Title, movie.Year)
				}
				bot.SendMessage(message.Chat, text, &tb.SendOptions{ParseMode: tb.ModeHTML})

			}

			if findSerie {
				log.Info("Matched find series handler")
				match := strings.TrimPrefix(message.Text, "/find series ")
				series := sonarr.SearchForSeries(match)
				text := fmt.Sprintf("<b>Series found %d</b> \n", len(series))
				for _, serie := range series {
					text = text + fmt.Sprintf("/download_series_%d ", serie.TvdbId)
					text = text + fmt.Sprintf("➸ %s (%d) \n", serie.Title, serie.Year)
				}
				bot.SendMessage(message.Chat, text, &tb.SendOptions{ParseMode: tb.ModeHTML})
			}

			if downloadSerie {
				log.Info("Matched download serie handler")
				tvdbidm, err := strconv.Atoi(strings.TrimPrefix(message.Text, "/download_series_"))
				if err != nil {
					log.Error(err)
				}
				err = sonarr.DownloadSerie(tvdbidm)

				if err != nil {
					bot.SendMessage(message.Chat, "Oops, there was a problem please try again later!", &tb.SendOptions{ParseMode: tb.ModeHTML})
				}

				bot.SendMessage(message.Chat, "Success!", &tb.SendOptions{ParseMode: tb.ModeHTML})
			}

			if downloadMovie {
				log.Info("Matched download movie handler")
				tmdbid, err := strconv.Atoi(strings.TrimPrefix(message.Text, "/download_movie_"))
				if err != nil {
					log.Error(err)
				}
				err = radarr.DownloadMovie(tmdbid)

				if err != nil {
					bot.SendMessage(message.Chat, "Oops, there was a problem please try again later!", &tb.SendOptions{ParseMode: tb.ModeHTML})
				}

				bot.SendMessage(message.Chat, "Success!", &tb.SendOptions{ParseMode: tb.ModeHTML})
			}
		}
	}
}