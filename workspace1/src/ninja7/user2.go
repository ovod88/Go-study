package main

import "fmt"

type User struct {
	Name string
	Email string
}

type Admin struct {
    User
    Level string
}

type Notifier interface {
    notify()
}

func (u *User) notify() {
	fmt.Println("Notified user ", u.Name)
}

func (u *Admin) notify() {
	fmt.Println("Notified admin ", u.Name)
}

func SendNotification(notify Notifier) {
    notify.notify()
}

func main() {
	admin := &Admin{
        User: User{
            Name:  "john smith",
            Email: "john@email.com",
        },
        Level: "super",
    }

    SendNotification(admin)
	admin.notify()
}