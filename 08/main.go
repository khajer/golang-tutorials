package main

import "fmt"

type Address struct {
	Address string
}

type UserProfile struct {
	Id        int
	FirstName string
	LastName  string
	Address   Address
}

// method
func (u UserProfile) GetUserProfile() {
	fmt.Println(u.FirstName)
	fmt.Println(u.LastName)
}

func (u UserProfile) Show() string {
	fmt.Println("show log string ", u.FirstName)
	return "5555"

}

type ToLogString interface { // interface have method

	Show() string
}

func main() {
	fmt.Println("test")

	// map (key, value)
	user := map[string]string{
		"firstname": "itsara",
		"lastname":  "konsombut",
	}

	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty

	fmt.Println(user["firstname"])
	fmt.Println(user["lastname"])

	fmt.Println(a)
	delete(a, "year")

	v, ok := a["model"]
	if ok {
		fmt.Println(v)
	}
	fmt.Println(a)
	a["year"] = "1964"
	fmt.Println(a)

	// struct
	userProfile := UserProfile{
		Id:        3,
		FirstName: "itsara",
		LastName:  "konsombut",
		Address: Address{
			Address: "home",
		},
	}
	fmt.Println(userProfile)
	fmt.Println(userProfile.Id)
	userProfile.GetUserProfile()

	// interface

	Check(userProfile) // like a bank 3 step struct must be have similar method

}

func Check(a ToLogString) {
	a.Show()

}
