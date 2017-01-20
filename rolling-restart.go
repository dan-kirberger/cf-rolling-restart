package main

import (
	"fmt"
	"strconv"
	"time"
	"code.cloudfoundry.org/cli/plugin"
)

type RollingRestart struct{}

func (c *RollingRestart) Run(cliConnection plugin.CliConnection, args []string) {

	if args[0] == "rolling-restart" {
		println("Performing rolling restart of " + args[1])

		app, _ := cliConnection.GetApp(args[1])

		for i := 0; i < app.InstanceCount; i++ {
			cliConnection.CliCommandWithoutTerminalOutput("restart-app-instance", args[1], strconv.Itoa(i))

			app, _ := cliConnection.GetApp(args[1])

			waitForMore := true

			fmt.Printf("\rInstance %v restarting. Waiting for minimum healthy instances before proceeding (Currently %v/%v)", i, app.RunningInstances, app.InstanceCount)
			for waitForMore {
				app, _ := cliConnection.GetApp(args[1])
				fmt.Printf("\rInstance %v restarting. Waiting for minimum healthy instances before proceeding (Currently %v/%v)", i, app.RunningInstances, app.InstanceCount)
				time.Sleep(time.Second * 2)
				waitForMore = app.InstanceCount > app.RunningInstances + 0
			}
			println()
		}

		println("Successfully restarted " + args[1])
	}
}

func (c *RollingRestart) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "rolling-restart",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "rolling-restart",
				HelpText: "Perform a rolling restart of a specified application. Restarting each instance individually",

				UsageDetails: plugin.Usage{
					Usage: "rolling-restart\n   cf rolling-restart my-app",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(RollingRestart))
}
