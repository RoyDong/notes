package main


import (
    "github.com/roydong/potato"
    "github.com/roydong/notes/controller"
    "github.com/roydong/notes/controller/admin"
)

func init() {
    potato.Init()

    //define template funcs
    potato.T.Funcs(map[string]interface{} {
    })

    //the map keys here must corresponds with 
    //the controller configured in routes.yml
    potato.R.Controllers(map[string]interface{} {
        "error": new(controller.Error),
        "main": new(controller.Main),
        "topic": new(controller.Topic),
        "comment": new(controller.Comment),

        "admin_user": new(admin.User),
        "admin_topic": new(admin.Topic),
    })
}

func main() {
    potato.Serve()
}



