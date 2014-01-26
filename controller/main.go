package controller

import ()

type Main struct {
    Base
}

func (c *Main) Index() {

    c.Render("index", nil)
}

func (c *Main) About() {

    c.Render("about", nil)
}
