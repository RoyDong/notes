package controller

import (
    "github.com/roydong/potato"
)

type Index struct {
    *potato.Controller
}

func (c *Index) Show() {
    c.Render("layout", map[string]string{"title": "nihao", "content": "欢迎"})
}


