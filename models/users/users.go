package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var UsersDB []User

func (u *User) GetUserById() bool {

	for _, user := range UsersDB {
		if user.Id == u.Id {
			u.FirstName = user.FirstName
			u.LastName = user.LastName
			u.Email = user.Email
			return true
		}
	}
	fmt.Println("Failed to Get User by Id")
	return false
}

func LoadUsers(filepath string) {

	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Cound not read json file")
	}
	if err := json.Unmarshal(byteValue, &UsersDB); err != nil {
		fmt.Println("Invalid json format")
	}

	fmt.Printf("Found %d users\n", len(UsersDB))
	for i := 0; i < len(UsersDB); i++ {
		fmt.Println("id: " + strconv.FormatInt(UsersDB[i].Id, 10))
		fmt.Println("User firstname: " + UsersDB[i].FirstName)
		fmt.Println("User lastname: " + (UsersDB[i].LastName))
		fmt.Println("User email: " + UsersDB[i].Email)
		fmt.Println("-------------------------------")
	}
}
