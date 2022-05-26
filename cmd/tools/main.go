package main

import (
	"github.com/spf13/cobra"

	"aegis"
	"aegis/cmd/tools/cmd"
	"aegis/internal/config"
	_ "aegis/internal/config/auto"
)

func main() {
	cliName := "tools"
	cli := &cobra.Command{
		Use:     cliName,
		Version: aegis.VersionString(),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	cli.AddCommand(cmd.MustMakeMigrationMetadataCommand(config.MySQL))
	cli.AddCommand(cmd.MustMakeMigrationWarehouseCommand(config.ClickHouse))
	cli.AddCommand(cmd.MustMakeSeedMetadataCommand())
	_ = cli.Execute()
}
