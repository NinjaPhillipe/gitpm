package main

type Profile struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func remove(slice []Profile, s int) []Profile {
	return append(slice[:s], slice[s+1:]...)
}
