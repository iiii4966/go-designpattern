package builder

import (
	"log"
	"net"
	"regexp"
	"strings"
)

type UserBuilder struct{}

func (b UserBuilder) setEmail(u *User, email string) UserBuilder {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	ok := func(e string) bool {
		if len(e) < 3 && len(e) > 254 {
			return false
		}
		if !emailRegex.MatchString(e) {
			return false
		}
		parts := strings.Split(e, "@")
		mx, err := net.LookupMX(parts[1])
		if err != nil || len(mx) == 0 {
			return false
		}
		return true
	}(email)

	if !ok {
		log.Fatal("Invalid email form")
	}

	u.Email = email
	return b
}

func (b UserBuilder) setPassword(u *User, password string) UserBuilder {
	if len(password) < 10 {
		log.Fatal("Password length must be least 10")
	}
	u.Password = password
	return b
}

func (b UserBuilder) setAge(u *User, age int) UserBuilder {
	if age < 20 {
		log.Fatal("Age must be under 20")
	}
	u.Age = age
	return b
}

func (b UserBuilder) setName(u *User, name string) UserBuilder {
	u.Name = name
	return b
}

// Essential Initialization Method
func (b UserBuilder) build(email, password, name string, age int) *User {
	user := &User{
		Email:    "",
		Password: "",
		Name:     "",
		Age:      0,
	}
	b.setEmail(user, email).setPassword(user, password)
	if name != "" { // default parameter not support; Why?
		b.setName(user, name)
	}
	if age != 0 {
		b.setAge(user, age)
	}
	return user
}

type User struct {
	Email    string
	Password string
	Name     string
	Age      int
}

func New(email, password, name string, age int) *User {
	builder := UserBuilder{}
	return builder.build(email, password, name, age)
}
