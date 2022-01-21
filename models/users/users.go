package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type User struct {
	Id       int64  `json:"id"`
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

func LoadUsers(filepath string) {
	var users []User

	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Cound not read json file")
	}
	if err := json.Unmarshal(byteValue, &users); err != nil {
		fmt.Println("Invalid json format")
	}

	fmt.Printf("Found %d users\n", len(users))
	for i := 0; i < len(users); i++ {
		fmt.Println("id: " + strconv.FormatInt(users[i].Id, 10))
		fmt.Println("User name: " + users[i].FistName)
		fmt.Println("User lastname: " + (users[i].LastName))
		fmt.Println("User email: " + users[i].Email)
		fmt.Println("-------------------------------")
	}
}
