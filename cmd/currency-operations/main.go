package main

import "currency-operations/internal/app"

const ConfigPath = "config/config.yaml"

func main() {
	app.Run(ConfigPath)
}
