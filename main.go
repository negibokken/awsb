package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.HideHelp = true
	app.Version = Version
	app.Author = "bokken"
	app.Usage = "browse aws console page easily"
	cli.AppHelpTemplate = AppHelp

	app.Flags = []cli.Flag{
		cli.HelpFlag,
		cli.BoolFlag{
			Name:  "service-list, S",
			Usage: "show service list",
		},
		cli.BoolFlag{
			Name:  "region-list, R",
			Usage: "show region list",
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = func(ctx *cli.Context) {
		if ctx.IsSet("service-list") {
			fmt.Println("service-list")
			os.Exit(0)
		} else if ctx.IsSet("region-list") {
			fmt.Println("region-list")
			os.Exit(0)
		}

		if ctx.NArg() < 2 {
			cli.ShowAppHelp(ctx)
			os.Exit(0)
		}
		service := ctx.Args().Get(0)
		region := ctx.Args().Get(1)

		fmt.Println(service, region)
	}

	app.Run(os.Args)

	// if len(os.Args) != 3 {
	// 	fmt.Println("")
	// 	os.Exit(1)
	// }
	// service := os.Args[1]
	// region := os.Args[2]
	// url := fmt.Sprintf("https://console.aws.amazon.com/%s/home?region=%s#/home", service, region)
	// switch runtime.GOOS {
	// case "linux":
	// 	exec.Command("xdg-open", url).Start()
	// case "windows":
	// 	exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	// case "darwin":
	// 	exec.Command("open", url).Start()
	// default:
	// 	fmt.Println("Your PC is not supported.")
	// }
}
