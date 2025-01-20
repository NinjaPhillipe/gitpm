package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("No args use help")
		return
	}

	command := args[1]

	fmt.Println("Debug ", args)

	path := buildConfigPath()

	if !configFolderExist(path) {
		fmt.Println(path + " does not exist")
		createConfigFolder(path)
		fmt.Println("Dir has been created")
	}

	profilesPath := filepath.Join(path, "profiles.yaml")

	switch command {
	case "help":
		displayHelp()
	case "add":
		addProfiles(profilesPath)
	case "list":
		listProfiles(profilesPath)
	case "delete":
		deleteProfile(profilesPath, 0)
	default:
		fmt.Println("ERROR Unknow command: ", command)
	}

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

func buildConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR getting home dir %s\n", err)
		os.Exit(1)
		return ""
	}
	return filepath.Join(home, ".config/gitpm")
}

func displayHelp() {
	// TODO see sub menu help
	fmt.Println("TODO HELP MENU")
	fmt.Println("command: help list add")
}

func addProfiles(profilesPath string) {

	profiles := []Profile{
		{Name: "Alice", Email: "alice@mail.com"},
		{Name: "Bob", Email: "bob@mail.com"},
	}

	data, err := yaml.Marshal(profiles)
	if err != nil {
		fmt.Println("Error serializing to YAML", err)
		return
	}

	err = os.WriteFile(profilesPath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
}

func readFile(profilesPath string) []Profile {
	fmt.Println("TODO list profiles")
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

func listProfiles(profilesPath string) {

	profiles := readFile(profilesPath)

	for id, profile := range profiles {
		fmt.Printf("Id : %d, Name: %s, email %s\n", id, profile.Name, profile.Email)
	}
}

func deleteProfile(profilePath string, id int) {

}
