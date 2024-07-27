# FabricPrompt

FabricPrompt is a simple and effective command-line tool written in Go that serves as a wrapper for [Fabric](https://github.com/danielmiessler/fabric), an open-source framework for augmenting human capabilities using AI.

## Demo
![Demo GIF](https://github.com/ndit-dev/FabricPrompt/blob/main/fabricP%20demo.gif)

## Description

FabricPrompt simplifies interaction with Fabric by offering a user-friendly interface to select and run Fabric patterns directly from the terminal. It supports input from both standard input and clipboard, making it flexible for various use cases. Most importantly, it allows you to choose patterns with a fuzzy-finder interface, eliminating the need to memorize pattern names and enhancing your workflow efficiency.

## Features

- Run Fabric patterns from a command line interface
- Support for input from stdin or clipboard
- Interactive pattern selector with fuzzy search
- Ability to add some optional fabic command lince switches, such as stream, context.md and save

## Installation

### Prerequisites

- [Fabric](https://github.com/danielmiessler/fabric)
- Go 1.22.2 or later
- Git
- xclip (linux only)

#### Mac users
```
brew update
brew upgrade
brew install go git
```

#### Ubuntu/Debian:
```bash
sudo apt-get update
sudo apt-get install golang-go xclip git
```

### Installing FabricPrompt from source

1. Clone the repository:
```bash
git clone https://github.com/ndit-dev/FabricPrompt.git
```
2. Change to the project directory:
```bash
cd FabricPrompt
```

3. Build the project:
```bash
go build -o fabricp
```

4. (Optional) Move the binary to a directory in your PATH to use it from anywhere:
```bash
sudo mv fabricp /usr/local/bin/
```

## Usage

run `fabricp` in your console after installation and follow the on screen instructions (see demo above)

## Contributing

Contributions to FabricPrompt are welcome! However, at this stage, I am not sure if I will actively develop this further. I created this tool to make my everyday interaction with Fabric smoother. Ideally, this serves as inspiration for the original Fabric project, where similar functionality could be incorporated as a switch or similar feature.

But please follow these steps if you would like to contribute:

1. Fork the repository
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## Future Features

Here are some additional features I'm considering for future updates:

- Use `tcell` or `termbox` to enhance input, making it more like a text editor.
- Add an option for users to save output directly to the clipboard.
- Provide an option to display output using `batcat`, `mdcat`, or natively within this Go program for better markdown readability.
- Extend support for more of Fabric's switched options.
- Implement a configuration file or environment variables to allow users to define default options and the default path for Fabric patterns.
- Avoid hardcoding the path to Fabric patterns; instead, dynamically determine their installation path and use it.
- add dependency checks for xclip
- test and add compability for older IntelMacs, should be a easy fix, just don't have one to test it on

## License

This project is licensed under the MIT License.

## Acknowledgments

This project is built upon the amazing [Fabric](https://github.com/danielmiessler/fabric) by Daniel Miessler. If you ended up here without knowing what Fabric is, you should start by looking at that project. This is merely a wrapper to make everyday usage more convenient.
