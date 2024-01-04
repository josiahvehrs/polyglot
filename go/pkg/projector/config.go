package projector

import (
	"fmt"
	"os"
	"path"
)

type Operation = int

const (
	Print Operation = iota
	Add
	Remove
)

type Config struct {
	Args      []string
	Operation Operation
	Pwd       string
	Config    string
}

func GetPwd(opts *Opts) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}

	return os.Getwd()
}

func getConfig(opts *Opts) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}

	confDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(confDir, "projector", "projector.json"), nil
}

func getOperation(opts *Opts) Operation {
	if len(opts.Args) == 0 {
		return Print
	}

	if opts.Args[0] == "add" {
		return Add
	}

	if opts.Args[0] == "rm" {
		return Remove
	}

	return Print
}

func getArgs(opts *Opts) ([]string, error) {
	if len(opts.Args) == 0 {
		return []string{}, nil
	}

	opertion := getOperation(opts)
	if opertion == Add {
		if len(opts.Args) != 3 {
			return nil, fmt.Errorf("add requires 2 arguments. got=%d", len(opts.Args))
		}
		return opts.Args[1:], nil
	}

	if opertion == Remove {
		if len(opts.Args) != 2 {
			return nil, fmt.Errorf("remove requires 1 argument. got=%d", len(opts.Args))
		}
		return opts.Args[1:], nil
	}

	if len(opts.Args) > 1 {
		return nil, fmt.Errorf("print requires 0 or 1 arguments. got=%d", len(opts.Args))
	}

	return opts.Args, nil
}

func NewConfig(opts *Opts) (*Config, error) {
	pwd, err := GetPwd(opts)
	if err != nil {
		return nil, err
	}

	operation := getOperation(opts)

	config, err := getConfig(opts)
	if err != nil {
		return nil, err
	}

	args, err := getArgs(opts)
	if err != nil {
		return nil, err
	}

	return &Config{
		Pwd:       pwd,
		Config:    config,
		Operation: operation,
		Args:      args,
	}, nil
}
