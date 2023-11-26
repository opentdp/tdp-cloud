package subset

import (
	"flag"
)

type FlagSet struct {
	*flag.FlagSet
	Comment string
	Execute func()
}

func NewFlagSets() map[string]*FlagSet {

	server := serverFlag()
	worker := workerFlag()
	update := updateFlag()

	return map[string]*FlagSet{
		server.Name(): server,
		worker.Name(): worker,
		update.Name(): update,
	}

}
