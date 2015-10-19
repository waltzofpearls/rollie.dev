package main

import (
	"flag"
	"fmt"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "c", "config.json", "Path to config file")
	flag.Parse()
	flag.Visit(func(v *flag.Flag) {
		fmt.Printf("%s - %s: %s\n", v.Usage, v.Name, v.Value)
	})

	config := NewConfigFile(configPath)

	app := NewApp(config)
	app.useStaticRouter("./static/")
	app.useRouter("/", &Index{})
	app.useRouter("/api", &Api{})
	app.Run()
}
