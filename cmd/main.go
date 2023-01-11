package cmd

import (
	"flag"
)

type FlagSet struct {
	*flag.FlagSet
	Comment string
}
