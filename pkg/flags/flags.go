package flags

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
)

// CID is a command identifier.
type CID string

const CreateCluster CID = "create.cluster"

type Flag struct {
	Name    string
	Usage   string
	Default string
}

type FlagSet []Flag

func (d *FlagSet) Append(f Flag) {
	*d = append(*d, f)
}

func FromFlagSet(flags *flag.FlagSet) FlagSet {
	var s FlagSet
	flags.VisitAll(func(f *flag.Flag) {
		s.Append(Flag{
			Name:    f.Name,
			Usage:   f.Usage,
			Default: f.DefValue,
		})
	})
	return s
}

func FromPFlagSet(flags *pflag.FlagSet) FlagSet {
	var s FlagSet
	flags.VisitAll(func(f *pflag.Flag) {
		s.Append(Flag{
			Name:    f.Name,
			Usage:   f.Usage,
			Default: f.DefValue,
		})
	})
	return s
}

type FlagParser interface {
	Parse(args []string) error
}

type flagDefinition struct {
	Set    FlagSet
	Parser FlagParser
}

// Registrar is used by plugin authors to register arbitrary flag sets that are to be
// bound to commands.
type Registrar struct {
	flags map[CID]flagDefinition
}

func NewRegistrar() *Registrar {
	return &Registrar{flags: make(map[CID]flagDefinition)}
}

func (f *Registrar) RegisterFlags(cmd CID, set FlagSet, parser FlagParser) {
	f.flags[cmd] = flagDefinition{
		Set:    set,
		Parser: parser,
	}
}

func (f *Registrar) hasFlags(cmd CID) bool {
	_, ok := f.flags[cmd]
	return ok
}

// getFlagSet retrieves the FlagSet associated with cmd. If there is not flag set associated with
// cmd a nil FlagSet is returned.
func (f *Registrar) getFlagSet(cmd CID) FlagSet {
	if flags, ok := f.flags[cmd]; ok {
		return flags.Set
	}
	return nil
}

// parse flags using the FlagParser associated with cmd.
func (f *Registrar) parse(cmd CID, flags []string) error {
	def, ok := f.flags[cmd]
	if !ok {
		return fmt.Errorf("")
	}
	return def.Parser.Parse(flags)
}
