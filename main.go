package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) != 3 {
		fmt.Println("")
		os.Exit(1)
	}
	service := os.Args[1]
	region := os.Args[2]
	fmt.Println(service)
	fmt.Println(region)
	url := fmt.Sprintf("https://console.aws.amazon.com/%s/home?region=%s#/home", service, region)
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		fmt.Println("Your PC is not supported.")
	}
}
