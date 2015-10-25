package main

import (
	"flag"
	"fmt"

	"github.com/waltzofpearls/tetris-go/libs"
	"github.com/waltzofpearls/tetris-go/routes"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "c", "config.json", "Path to config file")
	flag.Parse()
	flag.Visit(func(v *flag.Flag) {
		fmt.Printf("%s - %s: %s\n", v.Usage, v.Name, v.Value)
	})

	config := libs.NewConfigFile(configPath)

	app := libs.NewApp(config)
	app.UseStaticRouter("./static/")
	app.UseRouter("/", &routes.Index{})
	app.UseRouter("/api", &routes.Api{})
	app.Run()
}
