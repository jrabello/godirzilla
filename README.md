# Godirzilla
<div align="center">
  <img src="https://github.com/jrabello/godirzilla/raw/master/assets/godirzilla.png" alt="Godirzilla Logo" width="300"/>
</div>

Godirzilla (gdz) is a powerful CLI tool that enables you to execute commands across multiple directories simultaneously. It provides rich output and intelligent directory management, making project management across multiple repositories effortless and informative.

## Features

- **Directory Management**: Easily manage and organize groups of directories
- **AI-Powered Organization**: Uses AI to intelligently organize your directory groups
- **Parallel Command Execution**: Run commands across multiple repositories simultaneously
- **Rich Command Output**: Get detailed information about command execution status and results
- **User-Friendly Interface**: Intuitive CLI with clear command structure
- **Cross-Platform**: Works seamlessly across different operating systems
- **Configurable**: Customize behavior through configuration files
- **Performance Focused**: Optimized for efficient execution of parallel operations

## Installation

```sh
go get github.com/jrabello/godirzilla
```

## Usage

Godirzilla provides several commands to help you manage and execute operations across multiple directories:

### Basic Commands

```sh
# Run a command across directories
gdz run "git status"

# Manage directory groups
gdz grp add <group-name> <directory-path>
gdz grp list
gdz grp remove <group-name>

# Directory operations
gdz dir add <path>
gdz dir list
gdz dir remove <path>

# AI-powered directory organization
gdz ai organize
```

### Configuration

Godirzilla can be configured through a configuration file. The tool will automatically create a default configuration on first run.

## Contributing

Contributions to Godirzilla are always welcome! Whether it's feature requests, bug fixes, documentation improvements, or any other change, we appreciate all community involvement. 

For more information on contributing to Godirzilla, please see our [Contributor's Guide](CONTRIBUTING.md).

## License

Godirzilla is open-source software licensed under the [MIT License](LICENSE).
