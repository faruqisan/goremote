package main

import (
	"log"
	"os/exec"
)

const lockScreenScript = "gnome-screensaver-command"

func main() {

	// Test Lock the screen

	err := exec.Command(lockScreenScript, "-l").Start()

	if err != nil {
		log.Println("Error On Exec : ", lockScreenScript)
	}

}
