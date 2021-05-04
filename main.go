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
	Name string `yaml:"name"`
	Node Node   `yaml:"node"`
}

type Node struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
}

func (c *config) load() {
	d, err := ioutil.ReadFile("config.yaml")
	if err == nil {
		if err := yaml.Unmarshal(d, &c); err != nil {
			panic(err)
		}
	}
}

func main() {
	cfg := &config{}

	defaultCfg := &config{
		Test: "TestValueFromDefaultStruct",
		Name: "NameValueFromDefaultStruct",
		Node: Node{
			Hostname: "localhost",
			Port:     "8080",
		},
	}

	cfg.load()

	// set default values for undefined values in the config file.
	if err := mergo.Merge(cfg, defaultCfg); err != nil {
		panic(err)
	}

	flags := getFlagset(cfg)

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

func getFlagset(cfg *config) []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "test",
			EnvVars:     []string{"TEST"},
			Required:    false,
			Hidden:      false,
			Value:       cfg.Test,
			Destination: &cfg.Test,
		},
		&cli.StringFlag{
			Name:        "name",
			EnvVars:     []string{"NAME"},
			Required:    false,
			Hidden:      false,
			Value:       cfg.Name,
			Destination: &cfg.Name,
		},
		&cli.StringFlag{
			Name:        "hostname",
			EnvVars:     []string{"NODE_HOSTNAME"},
			Required:    false,
			Hidden:      false,
			Value:       cfg.Node.Hostname,
			Destination: &cfg.Node.Hostname,
		},
		&cli.StringFlag{
			Name:        "port",
			EnvVars:     []string{"NODE_HOSTNAME"},
			Required:    false,
			Hidden:      false,
			Value:       cfg.Node.Port,
			Destination: &cfg.Node.Port,
		},
	}
	return flags
}
