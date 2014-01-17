package admin

import (
    "github.com/roydong/potato"
    "github.com/roydong/notes/model"
)

type Base struct {
    potato.Controller
    user *model.User
}

func (c *Base) User() *model.User {
    if c.user == nil {
        c.user,_ = c.Request.Session.Value("user").(*model.User)
    }

    return c.user
}

func (c *Base) Init() {
    c.Layout = "admin/layout"

    if c.User() == nil {
        if potato.Env == "dev" {
            if user := model.UserModel.FindByEmail("i@roydong.com"); user != nil {
                c.Request.Session.Set("user", user, true)
                return
            }
        }

        c.Redirect("/admin/signin", 302)
    }
}
