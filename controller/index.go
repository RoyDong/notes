package controller

import (
    "github.com/roydong/potato"
)

type Index struct {
    *potato.Controller
}

func (c *Index) Show() {
    d := map[string]string{"name": "Roy", "message": "hello"}
    c.Render("index", d)
}


