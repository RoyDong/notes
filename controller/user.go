package controller

import (
    "log"
    "github.com/roydong/potato"
)

type User struct {
    *potato.Controller
}

func (c *User) Show() {
    id,_ := c.Request.GetInt("id")

    c.Response.Write([]byte("nihao"))
    log.Println("user show", id)
}

func (c *User) Topic() {


    c.Response.Write([]byte("..."))
}
