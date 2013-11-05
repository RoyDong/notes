package main


import (
    "github.com/roydong/potato"
    "github.com/roydong/topic/controller"
)

func init() {
    potato.Init()

    potato.H.Funcs(map[string]interface{} {
        "creator": func() string {return "Roy Dong"},
    })

    potato.R.Controllers([]interface{} {
        new(controller.Error),
        new(controller.User),
        new(controller.Main),
        new(controller.Topic),
    })
}

func main() {
    potato.Serve()
}



