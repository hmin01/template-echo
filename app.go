package main

import (
	"bytes"
	"strconv"
	app "template-echo/router"
)

const PORT = 4000

func main() {
	// Create port
	var buffer bytes.Buffer
	buffer.WriteString(":")
	buffer.WriteString(strconv.FormatInt(PORT, 10))
	strPort := buffer.String()
	// Run server
	server := app.CreateApp()
	server.Logger.Fatal(server.Start(strPort))
}