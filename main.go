package main


import (
    _"github.com/go-sql-driver/mysql"
    "github.com/roydong/potato"
    "./controller"
    "./controller/admin"
)

func init() {
    potato.Init()
}

func main() {
    //define template funcs
    potato.T.SetFuncs(map[string]interface{} {})

    //the map keys here must corresponds to
    //controller's configure in routes.yml
    potato.R.SetControllers(map[string]interface{} {
        "main": new(controller.Main),
        "topic": new(controller.Topic),
        "comment": new(controller.Comment),

        "admin_user": new(admin.User),
        "admin_topic": new(admin.Topic),
    })

    potato.Serve()
}

