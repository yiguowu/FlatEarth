package main

import (
	"FlatEarth/SharedLib"
	"fmt"
	"github.com/ip2location/ip2location-go"
)

func main() {
	db, err := ip2location.OpenDB("./IP2LOCATION-LITE-DB11.BIN")
	SharedLib.PanicOnError(err, SharedLib.WARNING)
	fmt.Println(db.Get_all(SharedLib.GetPublicIP().String()))
}
