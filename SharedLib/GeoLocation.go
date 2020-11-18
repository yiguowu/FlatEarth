package SharedLib

import "fmt"

func GetGeoLocation() {
	ip := GetPublicIP()
	fmt.Println(ip)
}
