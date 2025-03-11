package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
