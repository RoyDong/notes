package controller

import (
    "github.com/roydong/potato"
)

type User struct {
    *potato.Controller
}

func (c *User) Show() {
    id,_ := c.Request.Int("id")

    c.Render("aa", id)
}

func (c *User) Topic() {


    c.Response.Write([]byte("..."))
}
