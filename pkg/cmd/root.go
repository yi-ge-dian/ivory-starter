package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"www.ivorysql.org/ivory-starter/pkg/cmd/install"
)

var rootCmd = &cobra.Command{
	Use:   "ivoryctl",
	Short: "ivoryctl is a tool for ivorySQL",
	Long:  "ivoryctl is a tool for installing IvorySQL from the source code",
}

func init() {
	rootCmd.AddCommand(install.NewCommand())
	// rootCmd.AddCommand(uninstall.NewCommand())

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
