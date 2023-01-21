package cmd

import (
	"embed"
	"flag"
)

var FrontFS *embed.FS

var SubCommand string

type FlagSet struct {
	*flag.FlagSet
	Comment string
}

type FlagSets map[string]*FlagSet
