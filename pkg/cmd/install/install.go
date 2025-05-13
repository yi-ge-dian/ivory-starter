package install

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"www.ivorysql.org/ivory-starter/pkg/utils"
)

type Flagpole struct {
	SourcePath  string
	InstallPath string
	Version     string
}

func NewCommand() *cobra.Command {
	flags := &Flagpole{}
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install IvorySQL",
		Long:  "Install IvorySQL",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, args)
		},
	}

	cmd.Flags().StringVarP(&flags.SourcePath,
		"source",
		"s",
		// os.Getenv("HOME")
		os.Getenv("HOME")+"/test1",
		"IvorySQL source path",
	)

	cmd.Flags().StringVarP(&flags.InstallPath,
		"location",
		"l",
		// "/usr/local/ivorysql/ivorysql-4",
		"/test",
		"IvorySQL install path",
	)

	cmd.Flags().StringVarP(&flags.Version,
		"version",
		"v",
		"IVORY_REL_4_STABLE",
		"IvorySQL version",
	)

	return cmd
}

func runE(flags *Flagpole, args []string) error {
	// -------------------------------------Detect OS-------------------------------------
	fmt.Println("[DetectOS...]:")
	system, err := utils.DetectOS()
	if err != nil {
		fmt.Println("[DetectOS Failure]:", err)
		os.Exit(1)
	}
	fmt.Println("[DetectOS Success]:", system)
	fmt.Println("------------------------------------")

	// -------------------------------------Install Dependencies--------------------------
	switch system {
	case "ubuntu":
		fmt.Println("[InstallDependencies...]")
		u := &UbuntuInstaller{}
		if err := u.InstallDependencies(); err != nil {
			fmt.Println("[InstallDependencies Failure]:", err)
			os.Exit(1)
		}
		fmt.Println("[InstallDependencies Success]")
		fmt.Println("------------------------------------")

	case "centos":
		fmt.Println("[InstallDependencies...]:")
		// if err := InstallDependenciesCentOS(); err != nil {
		// 	fmt.Println("[InstallDependencies Failure]:", err)
		// 	os.Exit(1)
		// }
		fmt.Println("[InstallDependencies Success]")
		fmt.Println("------------------------------------")

	default:
		fmt.Println("[InstallDependencies Failure]:", err)
		os.Exit(1)
	}

	// -------------------------------------Install IvorySQL--------------------------
	fmt.Println("[InstallIvorySQL...]")
	if err := InstallIvorySQL(flags); err != nil {
		fmt.Println("[InstallIvorySQL Failure]:", err)
		os.Exit(1)
	}
	fmt.Println("[InstallIvorySQL Success]")
	fmt.Println("------------------------------------")

	// -------------------------------------Configure Env--------------------------
	fmt.Println("[ConfigureEnv...]")
	if err := ConfigureEnv(flags); err != nil {
		fmt.Println("[ConfigureEnv Failure]:", err)
		os.Exit(1)
	}
	fmt.Println("[ConfigureEnv Success]")
	fmt.Println("------------------------------------")

	return nil
}
