package cmd

import "flag"

type CMD interface {
	Usage()
	Run()
}

type baseCMD struct {
	Title string
	FlagSet *flag.FlagSet
}
