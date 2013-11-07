package controller

import (
)

type Main struct {
    Base
}

func (c *Main) Index() {

    c.Render("index", nil)
}

