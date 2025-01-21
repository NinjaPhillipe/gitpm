package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

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
	case "set":
		id := toInt(args[2], "expect a correct id to set")
		setProfile(profilesPath, id)
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

func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  help              Display this help message")
	fmt.Println("  add NAME EMAIL    Add a profile with the specified name and email")
	fmt.Println("  list              Display all profiles")
	fmt.Println("  rm ID             Remove a profile with the specified ID")

}

func addProfiles(profilesPath string, newProfile Profile) {

	profiles := readFile(profilesPath)

	save(profilesPath, append(profiles, newProfile))
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

func setProfile(profilesPath string, id int) {
	profiles := readFile(profilesPath)

	if id >= len(profiles) {
		fmt.Printf("No profile for id: %d", id)
		return
	}

	profile := profiles[id]

	// Set the username
	cmdName := exec.Command("git", "config", "--local", "user.name", profile.Name)
	if err := cmdName.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR failed to set name %s\n", err)
		os.Exit(1)
	}

	// Set the email
	cmdEmail := exec.Command("git", "config", "--local", "user.email", profile.Email)
	if err := cmdEmail.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR failed to set email %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Profile locally set\n   name:    %s\n   email:   %s", profile.Name, profile.Email)
}
