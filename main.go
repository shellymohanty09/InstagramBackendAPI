package main

import "fmt"

type User struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type Posts struct {
	ID              string `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption         string `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL        string `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	PostedTimestamp string `json:"postedTimestamp,omitempty" bson:"postedTimestamp,omitempty"`
}

func main() {
	fmt.Println("Hello World")
}
