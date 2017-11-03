package main

import (
	"github.com/spf13/cobra"
)

var (
	Password string
	Bot_token    string
	Sonarr_url   string
	Sonarr_token string
	Radarr_url   string
	Radarr_token string
	rootCmd      = &cobra.Command{
		Use:   "ranarr",
		Short: "ranarr telegram bot for sonarr and radarr",
		Run:   mainCmd,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&Password, "password", "", "telegram password for authentication")
	rootCmd.PersistentFlags().StringVar(&Bot_token, "token", "", "telegram token")
	rootCmd.PersistentFlags().StringVar(&Sonarr_url, "surl", "", "sonarr url")
	rootCmd.PersistentFlags().StringVar(&Sonarr_token, "stoken", "", "sonarr token")
	rootCmd.PersistentFlags().StringVar(&Radarr_url, "rurl", "", "radarr url")
	rootCmd.PersistentFlags().StringVar(&Radarr_token, "rtoken", "", "radarr token")
}

func main() {
	rootCmd.Execute()
}
