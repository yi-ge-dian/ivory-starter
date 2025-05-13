package uninstall

// import (
// 	"fmt"

// 	"github.com/spf13/cobra"
// )

// type Flagpole struct {
// 	Location string
// 	Version  string
// }

// func NewCommand() *cobra.Command {
// 	flags := &Flagpole{}

// 	cmd := &cobra.Command{
// 		Use:   "uninstall",
// 		Short: "uninstall IvorySQL",
// 		Long:  "uninstall IvorySQL",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			return runE(flags, args)
// 		},
// 	}

// 	cmd.Flags().StringVarP(&flags.Location,
// 		"location",
// 		"l",
// 		"/usr/local/ivorysql/ivorysql-4",
// 		"IvorySQL source install path",
// 	)

// 	cmd.Flags().StringVarP(&flags.Version,
// 		"version",
// 		"v",
// 		"IVORY_REL_4_STABLE",
// 		"IvorySQL version",
// 	)

// 	return cmd
// }

// func runE(flags *Flagpole, args []string) error {
// 	fmt.Println("Installing IvorySQL...")
// 	fmt.Printf("Installing Ivory %s to %s\n", flags.Version, flags.Location)

// 	fmt.Println(args)
// 	return nil
// }
