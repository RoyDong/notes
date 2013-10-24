package controller

import (
    "log"
    "github.com/roydong/potato"
)

type Error struct {
    *potato.Controller
}


func (c *Error) ServerError(e *potato.Error) {

    log.Println(e)
    c.Response.SetBody([]byte(e.String()))
}
