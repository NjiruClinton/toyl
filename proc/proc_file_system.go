package proc

// The proc file system

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetProcessID() int {
	return os.Getpid()
}
func GetExecutableName() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return ""
	}
	exeName := filepath.Base(exePath)
	return exeName
}

func GetProcDirs() ([]string, error) {
	cmd := exec.Command("ps", "-e", "-o", "pid=")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing ps command:", err)
		return nil, err
	}

	pids := strings.Fields(string(output))
	return pids, nil
}

func GetOpenFiles(pid int) ([]string, error) {
	cmd := exec.Command("lsof", "-p", fmt.Sprintf("%d", pid))
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var files []string
	for _, line := range lines[1:] { // Skip the header line
		fields := strings.Fields(line)
		if len(fields) > 8 {
			files = append(files, fields[8])
		}
	}
	return files, nil
}
