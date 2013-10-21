package controller

import (
    "log"
    "github.com/roydong/potato"
)

type Index struct {
    *potato.Controller
}

func (c *Index) Show() {
    log.Println("index show")
}


