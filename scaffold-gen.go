package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type projectConfig struct {
	Name string
	LocalPath string
	RepoUrl string
	StaticAssets bool
}

func main() {
	conf, err := setupParseFlags(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	// error handling
	errs := validateConf(conf)
	if errs != nil {
		for _, e := range errs {
			fmt.Println(e)
		}
		return
	}
	generateScaffold(os.Stdout, conf)
}

func setupParseFlags(w io.Writer, args []string) (projectConfig, error) {
	conf := projectConfig{}
	fs := flag.NewFlagSet("scaffold-gen", flag.ContinueOnError)
	fs.SetOutput(w) // This will send all output to the writer, w.
	fs.StringVar(&conf.Name, "n", "", "Project name")
	fs.StringVar(&conf.LocalPath, "d", "", "Project location on disk")
	fs.StringVar(&conf.RepoUrl, "r", "", "Project remote repository URL")
	fs.BoolVar(&conf.StaticAssets, "s", false, "Project will have static assets or not")

	err := fs.Parse(args)
	if fs.NArg() != 0 {
    return conf, errors.New("no positional parameters expected")
	}
	return conf, err
}


func validateConf(conf projectConfig) []error {
	var validationErrors []error
	if len(conf.Name) == 0 {
		validationErrors = append(validationErrors, errors.New("project name cannot be empty"))
	}
	if len(conf.LocalPath) == 0 {
		validationErrors = append(validationErrors, errors.New("local path cannot be empty"))
	}
	if len(conf.RepoUrl) == 0 {
		validationErrors = append(validationErrors, errors.New("repo url cannot be empty"))
	}
	// TODO: Add other validations
	return validationErrors
}

func generateScaffold(w io.Writer, conf projectConfig) {
	fmt.Printf("Generating scaffold for project %v in %v\n", conf.Name, conf.LocalPath)
}
