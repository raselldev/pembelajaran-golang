package main

import "sort"

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"raka", 1},
		{"rakaa", 190},
		{"rakaa", 2123},
		{"rakaaa", 123},
		{"rakaaa", 11},
	}
	sort.Sort(UserSlice(users))
}
