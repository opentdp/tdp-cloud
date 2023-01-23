package cmd

import (
	"embed"
	"flag"
)

var FrontFS *embed.FS

var SubCommand string

type FlagSets map[string]*FlagSet

type FlagSet struct {
	*flag.FlagSet
	Comment string
}
