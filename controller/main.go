package controller

import (
    "github.com/roydong/topic/model"
    "github.com/roydong/potato"
)

type Main struct {
    *potato.Controller
}

func (c *Main) Signin() {
    form := new(model.UserForm)

    if c.Request.Method == "POST" {
        form.Email,_  = c.Request.String("email")
        form.Passwd,_ = c.Request.String("passwd")
    }

    c.Render("signin", form)
}
