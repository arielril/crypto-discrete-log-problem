package main

import (
	"github.com/arielril/crypto-discrete-log-problem/internal/discrete"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
)

var (
	p string 
	g string
	h string
)

func init() {
	set := goflags.NewFlagSet()
	
	set.SetDescription("Discrete Log - Verifier")

	set.StringVar(&p, "p", "", "p value")
	set.StringVar(&g, "g", "", "g value")
	set.StringVar(&h, "h", "", "h value")

	_ = set.Parse()
	
}

func main() {
	gologger.Info().Msg("Discrete Logarithm Problem Verifier")

	if p == "" || g == "" || h == "" {
		gologger.Fatal().Msg("invalid value for p, g, or h")
	}

	res := discrete.Compute(p, g, h)

	gologger.Info().Msgf("Computation result: %v\n", res)
}

func setGroup(set *goflags.FlagSet, groupName, description string, flags ...*goflags.FlagData) {
	set.SetGroup(groupName, description)
	for _, currentFlag := range flags {
		currentFlag.Group(groupName)
	}
}