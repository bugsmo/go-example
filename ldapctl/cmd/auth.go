package cmd

import (
	"fmt"
	"ldapctl/g"
	"ldapctl/models"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authUserCmd)
	authCmd.AddCommand(authMultiCmd)
}

var authCmd = &cobra.Command{
	Use:       "auth",
	Short:     "Auth Test",
	Long:      "Authentication commands for LDAP Test",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{authUserCmd.Use, authMultiCmd.Use},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		multi Multi Auth Test
		single Single Auth Test
		`)
	},
}

var authUserCmd = &cobra.Command{
	Use:   "single [username] [password]",
	Short: "Single Auth Test",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Auth"

		username := args[0]
		password := args[1]

		startTime := time.Now()
		PrintStart(action)

		_, err := models.SingleAuth(g.Config().Ldap, username, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s auth test failed: %s \n", username, err.Error())
			PrintEnd(action, startTime)
			return
		}
		fmt.Printf("%s auth test success \n", username)
		PrintEnd(action, startTime)
	},
}

var authMultiCmd = &cobra.Command{
	Use:   "multi [filename]",
	Short: "Multi Auth Test",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Multi Auth"

		userList := args[0]
		authUsers, err := g.GetUsers(userList)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read file %s failed: %s \n", userList, err.Error())
			return
		}

		startTime := time.Now()
		PrintStart(action)

		res, err := models.MultiAuth(g.Config().Ldap, authUsers)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Multi Auth  failed: %s \n", err.Error())
			PrintEnd(action, startTime)
			return
		}
		fmt.Printf("Successed count %d \n", res.Successed)
		fmt.Printf("Failed count %d \n", res.Failed)
		fmt.Println("Failed users:")
		for _, failedMessage := range res.FailedMessages {
			fmt.Printf(" -- User: %s , Msg: %s \n", failedMessage.Username, failedMessage.Message)
		}
		PrintEnd(action, startTime)
	},
}
