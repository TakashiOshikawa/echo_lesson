package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
	age  int
}

func main() {
	fmt.Printf("start\n")

	user1 := User{id: 1, name: "user1", age: 23}

	fmt.Printf("%+v\n", user1)

	ary := []int{1, 2, 3, 4}

	fmt.Printf("%d\n", ary[0])

	for v := range ary {
		fmt.Printf("%d, ", ary[v])
	}
	fmt.Printf("\n")

	users := []User{
		user1,
		{id: 2, name: "user2", age: 33},
		{id: 3, name: "user3", age: 4},
	}

	for i := range users {
		fmt.Printf("%+v\n", users[i].age)
	}

	v := 3
	v2 := &v
	fmt.Println(v)
	fmt.Println(&v)
	fmt.Println(v2)
	fmt.Println(*v2)


	var user10 User = User{10, "user10", 19}
	fmt.Println(user10.name)
	newUser10 := changeUser10Name("new user name", user10)
	fmt.Println(newUser10)

	var newUser1 *User = newUser(11, "user11", 34)
	fmt.Println(newUser1.name)


	fmt.Println(user1.GetName())

	/*	noNameStruct := struct {
			aaa string
			bbb int
			ccc string
		}{
			aaa: "hello",
			bbb: 45,
			ccc: "world",
		}

		fmt.Printf("%+v\n", noNameStruct)

		fmt.Printf("%d\n", user1.age)
		fmt.Printf("%s\n", user1.name)
	*/
}

func (u *User) GetName() string {
	return u.name
}


func newUser(id int, name string, age int) *User {
	u := new(User)
	u.id = id
	u.name = name
	u.age = age

	return u
}

func changeUser10Name(name string, u User) User {
	u.name = name
	return u
}
