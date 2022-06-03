package etxea

import (
	"github.com/chrisdoherty4/etxea/pkg/flags"
	"github.com/hashicorp/go-plugin"
)

// Handshake defines unique properties of plugin host to plugin handshaking. There's nothing special
// about the MagicCookie data, its arbitrarily defined but must remain consistent as it features
// as a check when plugins are launched.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "ETXEA_PLUGIN_MAGIC_COOKIE",
	MagicCookieValue: "d9338fc2c7711733d3bb2e2d6e9bcda122d78fc1aa173afde76900d600fc908b",
}

var Plugins = plugin.PluginSet{
	flags.PluginName:   flags.NewPlugin(nil),
	BindingsPluginName: &BindingsPlugin{},
}

type PluginDefinition struct {
	Flags    plugin.Plugin
	Bindings plugin.Plugin
}

func NewPluginSet(definition PluginDefinition) plugin.PluginSet {
	set := plugin.PluginSet{}

	if definition.Flags != nil {
		set[flags.PluginName] = definition.Flags
	}

	if definition.Bindings != nil {
		set[BindingsPluginName] = definition.Bindings
	}

	return set
}
