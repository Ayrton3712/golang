package main

import "fmt"

type Email struct {
	Address string
}

type SMS struct {
	Number string
}

type Notifier interface {
	Notify() string
}

func (e Email) Notify() string {
	return "Email sent to " + e.Address
}

func (s SMS) Notify() string {
	return "SMS sent to " + s.Number
}

func Send(n Notifier) {
	fmt.Println(n.Notify())
}

func main() {
	e := Email{Address: "123@nvidia.tw"}
	s := SMS{Number: "0912676767"}

	Send(e)
	Send(s)
}
