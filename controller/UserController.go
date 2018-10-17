package controller

import "fmt"

type UserController struct {
	Controller
}

func (this *UserController) Login() {
	fmt.Fprintf(this.W, "Welcome to login web site")
}
