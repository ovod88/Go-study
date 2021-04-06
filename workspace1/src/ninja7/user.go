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
	fmt.Println("Notified", u.Name)
}

func SendNotification(notify Notifier) {
    notify.notify()
}

func main() {
	// // Value of type User can be used to call the method
	// // with a value receiver.
	// bill := User{"Bill", "bill@email.com"}
	// bill.Notify()

	// // Pointer of type User can also be used to call a method
	// // with a value receiver.
	// jill := &User{"Jill", "jill@email.com"}
	// jill.Notify()

	// user := User{
    //     Name:  "janet jones",
    //     Email: "janet@email.com",
    // }

	// SendNotification(&user)

	admin := Admin{
        User: User{
            Name:  "john smith",
            Email: "john@email.com",
        },
        Level: "super",
    }

    // SendNotification(admin)
	admin.User.notify()
	admin.notify()
}