package cmd

import "flag"

type CMD interface {
	//Usage()
	//Parse()
}

type baseCMD struct {
	Title string
	FlagSet *flag.FlagSet
}
