# Pegasus-Script
Pegasus Script (PGS) is a scripting language specifically designed for the Go-based hacking shell, Pegasus.

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation
To install pegasus script, follow these steps:

1. Ensure you have Go installed on your system. You can download it from [here](https://go.dev/dl)

2. Clone the Pegasus Script repository:
`git clone https://github.com/Codezz-ops/Pegasus-Script.git 
`

3. Navigate to the cloned directory:
`cd pegasus-script`

4. Build the project:
`go build`

5. The executable will be located in the current directory. You can move it to a directory in your PATH for easier access.

## Getting Started
To start using Pegasus Script, you need to create a script file with the `.peg` extension. Here's a simple example to get you started:
```
# This is a comment in Pegasus Script
puts("Hello, world");
```
To run your script, use the following command:
`./main main.peg`

## Usage
Pegaus Script is designed to be easy to learn and use. Here are some basic commands and their usage:

- `puts("Hello, world!");`: Prints a message to the standard output.
- `let x: int = 10;` Declares a variable `x` and assigns it the value of `int 10`.
- `if (condition) { ... }`: Executes a block of code if the condition is true.

For a more indepth tutorial, refer to the [official documentation](#)

## Examples
Here are some examples to help you get started with Pegasus Script:

- **Hello, World**:
    ```
    puts("Hello, World!");
    ```

- **Variable Declaration and Printing**:
    ```
    let name: string = "Alice";
    puts("Hello, " + name + "!");
    ```

- **Conditional Statements**:
    ```
    let age: int = 25;
    if (age >= 18) {
        puts("You are an adult.");
    } else {
        puts("You are not an adult.");
    }
    ```

## Contributing

Contributions to Pegasus Script are welcome! If you find a bug or have a feature request, please open an issue on our [GitHub repository](https://github.com/Codezz-ops/Pegasus-Script/issues). If you'd like to contribute code, please fork the repository and submit a pull request.

## License

Pegasus Script is licensed under the GNUv2 License. See the [LICENSE](LICENSE) file for more details.
