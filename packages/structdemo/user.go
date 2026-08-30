package structdemo

import "fmt"

type User struct {
	Name  string // exported field — accessible outside package
	Email string // exported field
	age   int    // unexported field — NOT accessible outside package
}

// Constructor pattern — common in Go since unexported fields
// can't be set directly from outside the package.
func NewUser(name, email string, age int) User {
	return User{
		Name:  name,
		Email: email,
		age:   age,
	}
}

// Exported getter — the standard way to expose an unexported field
func (u User) Age() int {
	return u.age
}

func Demo() {
	u := NewUser("Prathmesh", "test@example.com", 21)
	fmt.Println("Name:", u.Name)            // ✅ works, exported
	fmt.Println("Email:", u.Email)          // ✅ works, exported
	fmt.Println("Age via getter:", u.Age()) // ✅ works, using exported method
	// fmt.Println(u.age) // ❌ would fail if called from main — unexported field
}
