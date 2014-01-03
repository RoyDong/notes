package main


import (
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
    _"github.com/go-sql-driver/mysql"
    "github.com/roydong/notes/controller"
    "github.com/roydong/notes/controller/admin"
)

func init() {
    potato.Init()
    orm.InitDefault()
}

func main() {
    //define template funcs
    potato.T.SetFuncs(map[string]interface{} {})

    //the map keys here must corresponds to
    //controller's configure in routes.yml
    potato.R.SetControllers(map[string]interface{} {
        "error": new(controller.Error),
        "main": new(controller.Main),
        "topic": new(controller.Topic),
        "comment": new(controller.Comment),

        "admin_user": new(admin.User),
        "admin_topic": new(admin.Topic),
    })

    potato.Serve()
}

