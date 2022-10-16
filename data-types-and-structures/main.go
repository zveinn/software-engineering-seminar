package main

import (
	"fmt"

	uuid "github.com/google/uuid"
)

func main() {

	InitializeUserIDQueue()

	User1 := CREATE_USER("Anon", "unknown", "anon@anon.is", 21, "super secret passphrase")

	isPasswordCorrect := User1.ComparePasswords("super not secret passphrase")
	fmt.Println("Was the password correct?:", isPasswordCorrect)

	User2 := CREATE_USER("Anon2", "unknown2", "anon2@anon.is", 22, "super secret passphrase 2")

	// ADD_USER_TO_MAP(User1)
	// ADD_USER_TO_MAP(User2)

	// for i := range USER_MAP {
	// 	fmt.Println(USER_MAP[i])
	// }

	ADD_USER_TO_SLICE(User1)
	ADD_USER_TO_SLICE(User2)

	for i := range USER_SLICE {
		fmt.Println(USER_SLICE[i])
	}
	fmt.Println("THIS IS THE USER CACHE LIST ...")
	for i := range UserCache {
		if UserCache[i] == nil {
			continue
		}
		fmt.Println(UserCache[i])
	}

}

var USER_MAP = make(map[string]*User)
var USER_SLICE = make([]*User, 0)

type User struct {
	ID      string
	CacheID int
	Name    string
	Age     int
	Gender  string

	Email    string
	Password string
}

func (U *User) ComparePasswords(password string) bool {
	return U.Password == password
}

func CREATE_USER(Name, Gender, Email string, Age int, Password string) (NewUser *User) {
	NewUser = new(User)
	NewUser.ID = uuid.NewString()
	NewUser.Name = Name
	NewUser.Age = Age
	NewUser.Email = Email
	NewUser.Password = Password

	select {

	case NewUser.CacheID = <-UserIDQueue:
	default:
		fmt.Println("WE COULD NOT GET A CACHE ID")
	}

	return
}

func ADD_USER_TO_MAP(U *User) {
	USER_MAP[U.Email] = U
}

func ADD_USER_TO_SLICE(U *User) {
	USER_SLICE = append(USER_SLICE, U)
	if U.CacheID != 0 {
		UserCache[U.CacheID] = U
	}
}

/////////////////////////////////////////
/////////////////////////////////////////
/////////////////////////////////////////
/////////////////////////////////////////
/////////////////////////////////////////
/////////////////////////////////////////

var UserIDQueue = make(chan int, 200010)

func InitializeUserIDQueue() {
	for i := 1; i < 200000; i++ {
		UserIDQueue <- i
	}
}

var UserCache [200000]*User
