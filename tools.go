//+build tools

package main

import (
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/jessevdk/go-assets-builder"
	_ "github.com/lestrrat-go/server-starter/cmd/start_server"
	_ "honnef.co/go/tools/cmd/megacheck"
)
