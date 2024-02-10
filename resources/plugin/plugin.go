package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Version = "development"
	Team    = "hermanschaaf"
	Kind    = "source"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin("chess-com", Version, Configure,
		plugin.WithTeam(Team),
		plugin.WithKind(Kind))
}
