package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (u *User) ageupdt(a int) {
	u.Age = a
}

func main() {
	jsonString := `{"name": "Саша", "age": 25, "email": "maria@mail.ru"}`
	var sup User
	err := json.Unmarshal([]byte(jsonString), &sup)
	if err != nil {
		log.Fatalln("Ошибка сериализации", err)
	}
	sup.ageupdt(43)
	fmt.Println(sup)
}
