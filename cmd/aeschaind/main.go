package main

import (
	"os"

	"github.com/aeternalism/aeschain/app"
	"github.com/aeternalism/aeschain/cmd/aeschaind/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
