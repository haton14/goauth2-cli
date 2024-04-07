/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/haton14/goauth2"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// getTokenCmd represents the getToken command
var getTokenCmd = &cobra.Command{
	Use:   "get-token",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getToken called")
		if fileExists("tokens.json") {
			tokensFile, err := os.Open("tokens.json")
			if err != nil {
				log.Fatal(err)
			}
			defer tokensFile.Close()
			tokens := oauth2.Token{}
			if err := json.NewDecoder(tokensFile).Decode(&tokens); err != nil {
				log.Print(err)
			}
			fmt.Println(tokens.AccessToken)
			return
		}
		ctx := context.Background()
		oauth2Conf, err := goauth2.NewOAuth2Config()
		if err != nil {
			log.Fatal(err)
		}
		authorizationURL, err := goauth2.BuildAutorizationRequestURL(oauth2Conf)
		if err != nil {
			log.Fatal(err)
		}
		authzCode, err := goauth2.GetAuthorizationCode(ctx, authorizationURL, oauth2Conf.RedirectURL)
		if err != nil {
			log.Fatal(err)
		}
		accessToken, err := goauth2.GetAccessToken(ctx, oauth2Conf, authzCode)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create("tokens.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := json.NewEncoder(f).Encode(accessToken); err != nil {
			log.Fatal(err)
		}
		fmt.Println(accessToken.AccessToken)
	},
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
func init() {
	rootCmd.AddCommand(getTokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
