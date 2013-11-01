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
    c.Render("user/signin", form)
}

func (c *User) Signup() {
    form := new(model.UserForm)

    if c.Request.Method == "POST" {
        form.LoadData(c.Request)
        if form.Valid() {
            m := model.UserModel
            if m.Exists(form.Email) {
                form.Message = "Email exists"
                goto L
            }

            user := new(model.User)
            user.Email = form.Email
            user.SetPasswd(form.Passwd)
            if m.Save(user) {
                c.Redirect("/")
                return
            }

            form.Message = "server error, could not save data"
        }
    }

    L:
    c.Title = "Admin - Sign up"
    c.Render("user/signup", form)
}
