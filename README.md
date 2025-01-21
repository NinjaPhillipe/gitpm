# Git Profile Manager

A command-line tool to manage multiple Git profiles locally. Easily add, list, remove, and set Git profiles on a per-project basis.

## Features
- **Add** a new Git profile.
- **List** all saved Git profiles.
- **Remove** an existing Git profile.
- **Set** a Git profile for the current project.

## Installation
## Usage

### Add a Profile
Add a new Git profile with a name, email:

```bash
gitpm add "John Doe" "john.doe@company.com"
```

### List Profiles
View all saved profiles:

```bash
gitpm list
```

Example output:
```
ID   | NAME            | EMAIL                                   
0    | Alice           | alice@gmail.com                         
1    | Bob             | bob@gmail.com       
```

### Remove a Profile
Remove an existing Git profile:

```bash
gitpm rm <ID>
```

Example:
```
ID   | NAME            | EMAIL                                   
0    | Alice           | alice@gmail.com                         
1    | Bob             | bob@gmail.com       
```
```bash
gitpm rm 0
```

Example output:
```
ID   | NAME            | EMAIL                                   
0    | Bob             | bob@gmail.com       
```

### Set a Profile
Set a specific Git profile for the current repository:

```bash
gitpm set <ID>
```

Example:
```bash
gitpm set 1
```

This updates the local `.git/config` file for the repository.

## Configuration
Profiles are stored in a configuration file located at `~/gitpm/profiles.yaml`. Each profile includes:
- `name`: The name associated with the profile.
- `email`: The email address associated with the profile.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.
