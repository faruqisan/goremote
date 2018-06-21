package main

import (
	"log"
	"net"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const lockScreenScript = "gnome-screensaver-command"

func main() {

	ip, err := getIPAddress()
	if err != nil {
		log.Println(err)
	}

	log.Println("System IP Address : ", ip)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	screenAPI := router.Group("/screen")
	{
		screenAPI.GET("/lock", lockScreen)
	}

	router.Run(":8080")

}

func getIPAddress() (ipAddr string, err error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Println(err)
			break
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip.IsGlobalUnicast() {
				ipAddr = ip.String()
				break
			}

		}
	}

	return
}

func lockScreen(c *gin.Context) {

	err := exec.Command(lockScreenScript, "-l").Start()

	if err != nil {
		log.Println("Error On Exec : ", lockScreenScript)
	}

}
