package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func InstallIvorySQL(flags *Flagpole) error {
	fmt.Println("	[Checking Path...]")
	if _, err := os.Stat(flags.SourcePath); os.IsNotExist(err) {
		return fmt.Errorf("source path does not exist: %s", flags.SourcePath)
	}

	if err := os.Chdir(flags.SourcePath); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", flags.SourcePath, err)
	}

	path := filepath.Join(flags.SourcePath, "IvorySQL")
	// if _, err := os.Stat(path); !os.IsNotExist(err) {
	// 	return fmt.Errorf("source code has been cloned, please delete the directory: %s", path)
	// }
	fmt.Println("	[Checking Path Success]")

	//-----------------------------------------------------------------
	// fmt.Println("	[Cloning Code...]")
	// cmd := exec.Command("git", "clone", "https://github.com/IvorySQL/IvorySQL.git")
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to clone repository: %v", err)
	// }
	// fmt.Println("	[Cloning Code Success]")

	//-----------------------------------------------------------------
	fmt.Println("	[Changing Version...]")
	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", path, err)
	}

	// // checkout the specified version
	// cmd = exec.Command("git", "checkout", "-b", flags.Version, "origin/"+flags.Version)
	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to checkout version %s: %v", flags.Version, err)
	// }
	fmt.Println("	[Changing Version Success]")

	//-----------------------------------------------------------------
	fmt.Println("	[Configuring...]")
	cmd := exec.Command("./configure", "--prefix="+flags.InstallPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run configure: %v", err)
	}
	fmt.Println("	[Configuring Success]")

	//-----------------------------------------------------------------
	fmt.Println("	[Building...]")
	cmd = exec.Command("make")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run make: %v", err)
	}
	fmt.Println("	[Building Success]")

	//-----------------------------------------------------------------
	fmt.Println("	[Installing...]")
	cmd = exec.Command("sudo", "make", "install")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run make install: %v", err)
	}
	fmt.Println("	[Installing Success]")

	//-----------------------------------------------------------------
	return nil
}

func ConfigureEnv(flags *Flagpole) error {
	fmt.Println("	[Configuring Environment...]")
	return nil
}
