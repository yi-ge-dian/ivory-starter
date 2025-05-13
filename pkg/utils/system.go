package utils

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func getOSReleaseInfo() (map[string]string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"`)
		info[key] = value
	}

	return info, scanner.Err()
}

func DetectOS() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	osInfo, err := getOSReleaseInfo()
	if err != nil {
		return "", fmt.Errorf("error reading OS information: %v", err)
	}

	osID, ok := osInfo["ID"]
	if !ok {
		return "", fmt.Errorf("could not determine OS ID")
	}

	osID = strings.ToLower(osID)

	switch osID {
	case "ubuntu":
		return "ubuntu", nil
	case "centos":
		return "centos", nil
	case "rhel":
		return "rhel", nil
	case "rocky":
		return "rocky", nil
	default:
		return "", fmt.Errorf("unsupported Linux distribution: %s", osID)
	}
}
