package controller

import (
    "github.com/roydong/notes/model"
    "github.com/roydong/potato"
)

type Base struct {
    potato.Controller
    user *model.User
}

func (c *Base) User() *model.User {
    if c.user == nil {
        c.user, _ = c.Request.Session.Value("user").(*model.User)
    }

    return c.user
}
