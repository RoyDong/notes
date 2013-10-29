package controller

import (
    "fmt"
    "github.com/roydong/potato"
)

type Index struct {
    *potato.Controller
}

func (c *Index) Show() {
    d := map[string]string{
        "name": "Roy",
        "time": fmt.Sprintf("%d", c.Request.Session.LastActivity),
    }

    c.Request.Session.Mount("user", "Roy Dong", true)
    c.Render("index", d)
}

func (c *Index) Home() {
    user,_ := c.Request.Session.String("user");
    d := map[string]string{
        "user": user,
    }
    c.Render("home", d)
}


