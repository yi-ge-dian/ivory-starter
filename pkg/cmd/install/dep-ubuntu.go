package install

import (
	"fmt"
	"os/exec"
)

// UbuntuInstaller implements DependencyInstaller for Ubuntu.
type UbuntuInstaller struct {
}

// List of commands to install dependencies

func (u *UbuntuInstaller) InstallDependencies() error {
	dependencies := []string{
		"sudo apt-get install -y flex bison libreadline-dev zlib1g-dev libssl-dev",
		"sudo apt-get install -y build-essential",
		"sudo apt-get install -y libicu-dev pkg-config",
		"sudo apt-get install -y libxml2-dev",
	}

	for _, cmd := range dependencies {
		if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
			return fmt.Errorf("failed to execute command '%s': %w", cmd, err)
		}
	}

	return nil
}
