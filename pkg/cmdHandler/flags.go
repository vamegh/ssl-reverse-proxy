package cmdHandler

import (
	"flag"
	"log"
)

type Flags struct {
	ConfigFile string
	defaultMsg string `default:"Specify the location of the config file""`
	defaultVal string `default:""`
}

func ArgParser() Flags{
	var args Flags
	flag.StringVar(&args.ConfigFile, "config", args.defaultVal, args.defaultMsg)
	flag.StringVar(&args.ConfigFile, "c", args.defaultVal, args.defaultMsg)
	flag.Parse()

	if args.ConfigFile == "" {
		flag.PrintDefaults()
		log.Fatalf("Please provide the location/path of the config file")
	}
	return args
}