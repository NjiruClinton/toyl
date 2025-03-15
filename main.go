package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"playgroundgo/proc"
	"strings"
	"time"
)

type Account struct {
	Name string `json:"name"`
}

func main() {
	// get pid of this process
	fmt.Println("Process ID:", os.Getpid())
	file, err := os.Open("coa.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var accounts []Account
	err = json.Unmarshal(byteValue, &accounts)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	exeName := filepath.Base(exePath)
	fmt.Println("Executable name:", exeName)

	dirs, err := proc.GetProcDirs()
	if err != nil {
		fmt.Println("Error getting proc directories:", err)
		return
	}
	fmt.Print("Proc directories: ")
	fmt.Println(dirs)

	files, err := proc.GetOpenFiles(26594)
	if err != nil {
		fmt.Println("Error getting open files:", err)
		return
	}
	fmt.Println("Open files:")
	for _, file := range files {
		fmt.Println(file)
	}

	var result []Account
	keyword := "AR"
	for _, account := range accounts {
		if strings.HasPrefix(account.Name, keyword) {
			result = append(result, account)
			if len(result) >= 1000 {
				break
			}
		}
	}
	//for _, account := range result {
	//	fmt.Printf("%+v\n", account)
	//}
	// run forever

	// memory test
	// increase memory usage
	var my_arrr []string
	for {
		time.Sleep(5 * time.Second)
		my_arrr = append(my_arrr, "adsdsd")
		fmt.Println("Running..rrrrrr")
		fmt.Println("Length of array:", len(my_arrr))
	}
}
