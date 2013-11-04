package controller

import (
    "github.com/roydong/potato"
)

type Error struct {
    Base
}


func (c *Error) ServerError(e *potato.Error) {

    c.Render("error/notfound", e)
}

func (c *Error) NotFound(e *potato.Error) {
    c.Render("error/notfound", e)
}
