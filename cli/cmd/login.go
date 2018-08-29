package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

type tokenWrap struct {
	Token string `json:"token"`
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to SPaaS server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter username:")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")
		fmt.Println("Enter password:")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		password := string(bytePassword)
		fmt.Println()
		if !viper.InConfig("url") || !viper.InConfig("port") {
			fmt.Println("URL and Port not set in config. Run \"paas setup\"")
			return
		}
		v := url.Values{}
		v.Add("username", username)
		v.Add("password", password)
		url := "http://" + viper.GetString("url") + ":" + viper.GetString("port") + "/login"
		resp, err := http.PostForm(url, v)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if resp.StatusCode == http.StatusUnauthorized {
			fmt.Println("Wrong login credentials")
			return
		}
		defer resp.Body.Close()
		token := &tokenWrap{}
		err = json.NewDecoder(resp.Body).Decode(token)
		if err != nil {
			fmt.Println(err.Error())
		}
		viper.Set("token", token.Token)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
