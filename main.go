package main


import (
    "log"
_    "time"
    "net/http"
    "github.com/roydong/potato"
)

func init() {


}


func main() {
    router := potato.NewRouter()
    router.InitFile("./router.yml")

    e := http.ListenAndServe(":80", router)
    log.Println(e)
}


