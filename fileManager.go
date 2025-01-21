package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func buildConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR getting home dir %s\n", err)
		os.Exit(1)
		return ""
	}
	return filepath.Join(home, ".config/gitpm")
}

func configFolderExist(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func createConfigFolder(path string) {
	err := os.Mkdir(path, os.ModePerm)

	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "Cannot create config folder %s error: %s\n", path, err)

	os.Exit(1)
}

func readFile(profilesPath string) []Profile {
	data, err := os.ReadFile(profilesPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR while reading file: %s error: %s\n", profilesPath, err)
		os.Exit(1)
	}

	var profiles []Profile
	err = yaml.Unmarshal(data, &profiles)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR error while unmarshal error: %s\n", err)
		os.Exit(1)
	}
	return profiles
}

func save(profilesPath string, profiles []Profile) {

	data, err := yaml.Marshal(profiles)
	if err != nil {
		fmt.Println("Error serializing to YAML", err)
		return
	}

	err = os.WriteFile(profilesPath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		os.Exit(1)
	}
}
