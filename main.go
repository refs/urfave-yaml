package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/imdario/mergo"

	"gopkg.in/yaml.v2"

	"github.com/urfave/cli/v2"
)

type config struct {
	Test string `yaml:"test"`
	Name string
}

func main() {
	cfg := &config{}

	defaultCfg := &config{
		Test: "DefaultConfigTest",
		Name: "DefaultConfigName",
	}

	d, err := ioutil.ReadFile("configs.yaml")
	if err == nil {
		if err := yaml.Unmarshal(d, &cfg); err != nil {
			panic(err)
		}
	}

	// set default values for undefined values in the config file.
	if err := mergo.Merge(cfg, defaultCfg); err != nil {
		panic(err)
	}

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "test",
			EnvVars:     []string{"APP_TEST"},
			Required:    false,
			Hidden:      false,
			Value:       cfg.Test,
			Destination: &cfg.Test,
		},
		&cli.StringFlag{
			Name:        "name",
			EnvVars:     []string{"APP_NAME"},
			Required:    false,
			Hidden:      false,
			Value:       cfg.Name,
			Destination: &cfg.Name,
		},
	}

	app := cli.App{
		Commands: []*cli.Command{
			{
				Name:      "run",
				UsageText: "runs an example app",
				Action: func(context *cli.Context) error {
					fmt.Printf("Action: %v\n", context.Value("test"))
					d, _ := json.MarshalIndent(cfg, "", "  ")
					fmt.Printf("Action: %+v", string(d))
					return nil
				},
				Flags: flags,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
