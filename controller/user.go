package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/topic/model"
)

type User struct {
    *potato.Controller
}

func (c *User) Home() {
    u,_ := c.Request.Session.Value("user").(*model.User)
    c.Render("user/home", u)
}

func (c *User) Signin() {
    form := new(model.UserForm)
    if c.Request.Method == "POST" {
        form.LoadData(c.Request)
        if form.Valid() {
            m := model.UserModel
            if user := m.FindByEmail(form.Email); user != nil {
                c.Request.Session.Set("user", user, true)
                c.Redirect("/home")
                return
            }

            form.Message = "user exists"
        }
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
                goto RENDER
            }

            user := new(model.User)
            user.Email = form.Email
            user.SetPasswd(form.Passwd)
            if m.Save(user) {
                c.Request.Session.Set("user", user, true)
                c.Redirect("/home")
                return
            }

            form.Message = "server error, could not save data"
        }
    }

    RENDER:
        c.Title = "Admin - Sign up"
        c.Render("user/signup", form)
}
