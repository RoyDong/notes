package controller

import (
)

type Topic struct {
    Base
}

func (c *Topic) New() {

    c.Render("topic/new", nil)
}
