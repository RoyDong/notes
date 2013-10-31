package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/topic/model"
)

type User struct {
    *potato.Controller
}


func (c *User) Signin() {
    form := new(model.UserForm)

    if c.Request.Method == "POST" {
        form.Email,_  = c.Request.String("email")
        form.Passwd,_ = c.Request.String("passwd")
    }

    c.Title = "Admin - Sign in"
    c.Render("signin", form)
}
