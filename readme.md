# FabricPrompt

FabricPrompt is a simple and effective command-line tool written in Go that serves as a wrapper for [Fabric](https://github.com/danielmiessler/fabric), an open-source framework for augmenting human capabilities using AI.

## Description

FabricPrompt simplifies interaction with Fabric by offering a user-friendly interface to select and run Fabric patterns directly from the terminal. It supports input from both standard input and clipboard, making it flexible for various use cases. And most important, allows you to choose pattern with a fuzzyfinder instead of forcing you to know them by heart.

## Features

- Run Fabric patterns directly from the command line
- Support for input from stdin or directly inputting a prompt
- Interactive pattern selector with fuzzy search
- Ability to add extra context from a file
- Option to save output to a file
- Real-time streaming of results
- Easy integration with other command-line tools

## Install

### Download latest build
You can find tha latest built versions here: https://github.com/ndit-dev/FabricPrompt/actions/workflows/build.yml

### Or build from source

To build FabricPrompt from source, follow these steps:

1. **Clone the repository**:
    ```sh
    git clone https://github.com/ndit-dev/FabricPrompt.git
    cd FabricPrompt
    ```

2. **Build the binary**:
    Make sure you have Go installed on your system. You can download and install it from the [official Go website](https://golang.org/dl/).

    ```sh
    go build -o fabricp
    ```

3. **Move the binary to your PATH**:
    Move the compiled binary to a directory that is in your system's PATH. For example, you can move it to `/usr/local/bin`:

    ```sh
    sudo mv fabricp /usr/local/bin/
    ```

4. **Verify the installation**:
    Ensure that the binary is accessible from your terminal by running:

    ```sh
    fabricp --help
    ```

    You should see the help message for FabricPrompt.

### Prerequisites

- Ensure you have [Fabric](https://github.com/danielmiessler/fabric) installed and properly configured on your system.

## Usage

type `fabricp` in your console after installation
[![Demo](https://github.com/ndit-dev/FabricPrompt/blob/main/fabricPrompt_demo.gif)]

You can pipe text directly to FabricPrompt, and it will use the piped text as input instead of prompting for user input or using the clipboard, just as you do with fabric normally. For example:
```
echo "your text here" | fabricp
```

This is the best way to pipe your clipboard to fabricp since it's inputfield is very simple and limitied. 
```
pbpaste | fabricp
```
By doing this you are still able to benefin from the fuzzy finder in fabricp as well as the the rest of the interface

## Configuration

FabricPrompt uses configuration files from Fabric. Make sure you have Fabric properly configured on your system.

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

- Add an option for users to save output directly to the clipboard.
- Provide an option to display output using `batcat`, `mdcat`, or natively within this Go program for better markdown readability.
- Extend support for more of Fabric's switched options.
- Implement a configuration file or environment variables to allow users to define default options and the default path for Fabric patterns.
- Avoid hardcoding the path to Fabric patterns; instead, dynamically determine their installation path and use it.

## License

This project is licensed under the MIT License.

## Acknowledgments

This project is built upon the amazing [Fabric](https://github.com/danielmiessler/fabric) by Daniel Miessler. If you ended up here without knowing what Fabric is, you should start by looking at that project. This is merely a wrapper to make everyday usage more convenient.
