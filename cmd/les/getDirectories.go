package main

import (
	"os"
	"os/user"
	"path/filepath"
)

type espdirectories struct {
	esproot  string
	build    string
	template string
}

func getDirectories() (espdirectories, error) {
	const espHomeDirectory = ".les"
	const buildDirectory = "api"
	const templateDirectory = "template"
	directories := espdirectories{}

	usr, err := user.Current()
	if err != nil {
		return directories, err
	}
	espDirectory := filepath.Join(usr.HomeDir, espHomeDirectory)
	fullTemplateDirectory := filepath.Join(espDirectory, templateDirectory)
	fullBuildDirectory := buildDirectory
	_ = os.MkdirAll(fullTemplateDirectory, os.ModePerm)
	directories.esproot = espDirectory
	directories.build = fullBuildDirectory
	directories.template = fullTemplateDirectory
	return directories, nil
}
