package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama  string
	Kelas string
}

type Users []User

func main() {
	var user User
	var users Users
	users = append(users, User{Nama: "NooB", Kelas: "A"})
	users = append(users, User{Nama: "Java", Kelas: "B"})
	users = append(users, User{Nama: "Docker", Kelas: "C"})
	fmt.Println("Data pengguna :", users)

	user.saveToFile(users)
	user.loadFromFile(users)
}

func (u User) loadFromFile(users []User) {
	data := []byte(`users`)
	err := json.Unmarshal(data, &u)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("yang dibaca dari file JSON :", users)
}

func (u User) saveToFile(users []User) {
	usersJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Data pengguna berhasil disimpan dalam file JSON. Bentuk JSON :", string(usersJSON))
}
