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

    log.Println("user show", id)
    potato.Panic(11, "aa")
}

func (c *User) Topic() {


    c.Response.Write([]byte("..."))
}
