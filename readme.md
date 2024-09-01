# FabricPrompt

FabricPrompt is a simple and effective command-line tool written in Go that serves as a wrapper for [Fabric](https://github.com/danielmiessler/fabric), an open-source framework for augmenting human capabilities using AI.

## Description

FabricPrompt simplifies interaction with Fabric by offering a user-friendly interface to select and run Fabric patterns directly from the terminal. It supports input from both standard input and clipboard, making it flexible for various use cases. And most important, allows you to choose pattern with a fuzzyfinder instead of forcing you to know them by heart.

## Features

- Run Fabric patterns directly from the command line
- Support for input from stdin or clipboard
- Interactive pattern selector with fuzzy search
- Ability to add extra context from a file
- Option to save output to a file
- Real-time streaming of results
- Easy integration with other command-line tools

## Installation

To install FabricPrompt, ensure you have Go installed on your system, then run:
```
go install github.com/ndit-dev/FabricPrompt@latest

```

## Usage

type `fabricp` in your console after installation
<placehold for video or gif>

You can pipe text directly to FabricPrompt, and it will use the piped text as input instead of prompting for user input or using the clipboard, just as you do with fabric normally. For example:
```
echo "your text here" | fabricp
```
## onfiguration

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

- Use `tcell` or `termbox` to enhance input, making it more like a text editor.
- Add an option for users to save output directly to the clipboard.
- Provide an option to display output using `batcat`, `mdcat`, or natively within this Go program for better markdown readability.
- Extend support for more of Fabric's switched options.
- Implement a configuration file or environment variables to allow users to define default options and the default path for Fabric patterns.
- Avoid hardcoding the path to Fabric patterns; instead, dynamically determine their installation path and use it.

## License

This project is licensed under the MIT License.

## Acknowledgments

This project is built upon the amazing [Fabric](https://github.com/danielmiessler/fabric) by Daniel Miessler. If you ended up here without knowing what Fabric is, you should start by looking at that project. This is merely a wrapper to make everyday usage more convenient.