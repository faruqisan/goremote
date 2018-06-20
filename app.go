package main

import (
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const lockScreenScript = "gnome-screensaver-command"
const setVolumeScript = "amixer set 'Master'"

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	screenAPI := router.Group("/screen")
	{
		screenAPI.GET("/lock", lockScreen)
	}

	volumeAPI := router.Group("/volume")
	{
		volumeAPI.GET("/set/:value", setVolume)
	}

	router.Run(":8080")

}

func lockScreen(c *gin.Context) {

	err := exec.Command(lockScreenScript, "-l").Start()

	if err != nil {
		log.Println("Error On Exec : ", lockScreenScript)
	}

}

func setVolume(c *gin.Context) {

	v := c.Param("value")

	script := setVolumeScript + " " + v + "%"

	err := exec.Command(script).Start()
	if err != nil {
		log.Println("Error On Exec : ", script)
	}

}
