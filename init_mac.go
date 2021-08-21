/*
* Author: keronori
* Date: 2021/08/15
 */

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// check whether docker command-line exists
	if existsCommand("docker") {
		fmt.Println("[PASS] docker command-line is installed.")
	} else {
		fmt.Println("[ERROR] Cannot find docker command-line")
		os.Exit(1)
	}
	// check whether pack CLI exists
	if existsCommand("pack") {
		fmt.Println("[PASS] pack CLI is installed.")
	} else {
		fmt.Println("[ERROR] Cannot find pack CLI")
		// fmt.Println("Install pack CLI: ")
		// fmt.Println("		$ brew install buildpacks/tap/pack")
		os.Exit(1)
	}
}

func existsCommand(name string) bool {
	err := exec.Command(name, "--help").Run()
	if err != nil {
		return false
	}
	return true
}
