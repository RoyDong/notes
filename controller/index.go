package controller

import (
    "fmt"
    "github.com/roydong/potato"
)

type Index struct {
    *potato.Controller
}

func (c *Index) Show() {

}

func (c *Index) Home() {
    user,_ := c.Request.Session.String("user");
    d := map[string]string{
        "user": user,
    }
    c.Render("home", d)
}


