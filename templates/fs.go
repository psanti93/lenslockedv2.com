package templates

import "embed"

//go:embed *
var FS embed.FS //embeds the templates into the binary
