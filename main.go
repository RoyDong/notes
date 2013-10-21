package main


import (
    "log"
_    "time"
    "github.com/roydong/potato"
    "github.com/roydong/topic/controller"
)

func init() {
    potato.Init()

    r := potato.R
    r.RegController(new(controller.Index))
    r.RegController(new(controller.User))
}

func main() {
    log.Fatal(potato.S.ListenAndServe())
}



