package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	authValue := fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
	return authValue
}

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

type user struct {
	number int
	name   string
}

type User struct {
	Name string
	Membership
}

type Membership struct {
	MessageCharLimit int
	Type             string
}

func (User User) SendMessage(message string, messageLength int) (string, bool) {

	if messageLength <= User.MessageCharLimit {
		return message, true
	} else {
		return "", false
	}

}

func newUser(name string, membership string) User {

	var limit int

	if membership == "premium" {
		limit = 1000
	} else {
		limit = 100
	}

	var newUser User

	newUser.Name = name
	newUser.Type = membership
	newUser.MessageCharLimit = limit

	return newUser

}
