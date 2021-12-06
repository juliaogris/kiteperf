package main

import (
	"fmt"
	"io"
	"log"

	"github.com/alecthomas/kong"
	"github.com/pkg/profile"
)

var version = "v0.0.0"

const helpDescription = `A go CLI template.`

type config struct {
	Arg string `arg:"" help:"Argument"`

	Version kong.VersionFlag `short:"V" help:"Print version information" group:"Other:"`
	Debug   bool             `short:"d" help:"Print debug information" group:"Other:"`

	out io.Writer
}

func main() {
	defer profile.Start(profile.ProfilePath("."), profile.CPUProfile).Stop()
	cfg := &config{}
	opts := []kong.Option{
		kong.Description(helpDescription),
		kong.Vars{"version": version},
		kong.PostBuild(updateHelpFlag),
		kong.ConfigureHelp(kong.HelpOptions{Compact: true, Summary: true}),
	}
	ctx := kong.Parse(cfg, opts...)
	log.Println(cfg)
	err := ctx.Run(cfg)
	ctx.FatalIfErrorf(err)
}

func (cfg *config) AfterApply() error {
	if !cfg.Debug {
		log.SetOutput(io.Discard)
	}
	return nil
}

func (cfg *config) Run() error {
	fmt.Fprintf(cfg.out, "Hello %s\n", cfg.Arg)
	return nil
}

func updateHelpFlag(k *kong.Kong) error {
	// Remove '.' at the end of the help message to keep consistent with others.
	k.Model.HelpFlag.Help = "Show context-sensitive help"
	// Move help flag to the bottom and promote action group to the top
	k.Model.HelpFlag.Group = &kong.Group{Key: "Other:", Title: "Other:"}
	k.Model.Flags = append(k.Model.Flags[1:], k.Model.Flags[0])
	return nil
}
