package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

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
		addProfiles(profilesPath, Profile{Name: args[2], Email: args[3]})
	case "list":
		listProfiles(profilesPath)
	case "rm":
		id := toInt(args[2], "expect a correct id to remove")
		rmProfile(profilesPath, id)
	default:
		fmt.Println("ERROR Unknow command: ", command)
	}

}

func toInt(input string, errMsg string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR %s is not an int %s\n", input, errMsg)
		os.Exit(1)
	}
	return num
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

func addProfiles(profilesPath string, newProfile Profile) {

	profiles := readFile(profilesPath)

	save(profilesPath, append(profiles, newProfile))
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

func listProfiles(profilesPath string) {
	profiles := readFile(profilesPath)
	display(profiles)
}

func display(profiles []Profile) {
	for id, profile := range profiles {
		fmt.Printf("Id : %d, Name: %s, email %s\n", id, profile.Name, profile.Email)
	}
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

func remove(slice []Profile, s int) []Profile {
	return append(slice[:s], slice[s+1:]...)
}

func rmProfile(profilesPath string, id int) {
	profiles := readFile(profilesPath)

	if id >= len(profiles) {
		fmt.Printf("No profile for id: %d", id)
		return
	}

	profiles = remove(profiles, id)

	display(profiles)

	save(profilesPath, profiles)
}
