package main


import (
    "github.com/roydong/potato"
    "github.com/roydong/topic/controller"
)

func init() {
    potato.Init()
    potato.R.RegControllers([]interface{}{
        new(controller.Index),
        new(controller.User),
        new(controller.Error),
    })
}

func main() {
    potato.Serve()
}



