package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
		cli.StringFlag{
			Name:  "region, r",
			Value: "s3",
			Usage: "specify region to move (default: s3)",
		},
		cli.StringFlag{
			Name:  "service, s",
			Value: "us-west-2",
			Usage: "specify service to move (default: us-west-2)",
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = action

	app.Run(os.Args)

}

func browse(url string) {
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

func action(ctx *cli.Context) {
	var service string
	var region string

	if ctx.IsSet("service-list") {
		fmt.Println(Service)
		os.Exit(0)
	} else if ctx.IsSet("region-list") {
		fmt.Println(Region)
		os.Exit(0)
	}

	if ctx.IsSet("region") || ctx.IsSet("service") {
		service = ctx.String("service")
		region = ctx.String("region")
		url := fmt.Sprintf("https://console.aws.amazon.com/%s/home?region=%s", service, region)
		browse(url)
		os.Exit(0)
	}

	if ctx.NArg() < 1 {
		cli.ShowAppHelp(ctx)
		os.Exit(0)
	}

	service = ctx.Args().Get(0)
	if ctx.Args().Get(1) == "" {
		region = "us-west-2"
	}
	region = ctx.Args().Get(1)

	url := fmt.Sprintf("https://console.aws.amazon.com/%s/home?region=%s", service, region)
	browse(url)
	os.Exit(0)
}
