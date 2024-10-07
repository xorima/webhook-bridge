package main

import "github.com/xorima/webhook-bridge/cmd"

// @title           Webhook Bridge API
// @description     This is a bridge to receive various webhook events and publish them to a channel.

// @contact.name   Jason Field
// @contact.url    https://github.com/xorima

// @license.name  MIT
// @license.url  https://github.com/xorima/webhook-bridge/blob/main/LICENSE

// @host      localhost:3000
// @BasePath  /

// @externalDocs.description  GitHub
// @externalDocs.url          https://github.com/xorima/webhook-bridge
func main() {
	cmd.Execute()
}
