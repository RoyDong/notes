package main


import (
    "github.com/roydong/potato"
    "github.com/roydong/topic/controller"
)

func init() {
    potato.Init()
    controller.Init()
}

func main() {
    potato.Serve()
}



