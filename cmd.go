package cmd

import "github.com/codegangsta/cli"

const HelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]
VERSION:
   {{.Version}}{{if len .Authors}}
AUTHOR(S): 
   {{range .Authors}}{{ . }}{{end}}{{end}}
{{if .Commands}}
COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .Flags}}OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

func New() *cli.App {
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) {}
	app.HideHelp = true
	cli.AppHelpTemplate = HelpTemplate

	return app
}
