# Pegasus Script Documentation

## Table of Contents

1. [Introduction](#introduction)
2. [Basic Concepts](#basic-concepts)
3. [Control Structures](#control-structures)
4. [Functions and Modules](#functions-and-modules)
5. [Object-Oriented Programming](#object-oriented-programming)
6. [Error Handling](#error-handling)
7. [Standard Library](#standard-library)
8. [Advanced Topics](#advanced-topics)
9. [Best Practices](#best-practices)
10. [Appendices](#appendices)
11. [Examples and Tutorials](#examples-and-tutorials)

## Introduction

### Overview of Pegasus Script

- Pegasus Script is a custom scripting language designed to simplify the development and deployment of tooling originally coded in Go for the Pegasus hacking shell. Pegasus, a Docker-based hacking shell, was conceived in 2023 with the primary goal of streamlining and enhancing the security testing and penetration testing process. The language was developed to leverage the robustness and efficiency of Go, while offering a more accessible and user-friendly interface for developers and security professionals.
- The primary target audience for Pegasus Script is cybersecurity professionals, including penetration testers, security researchers, and ethical hackers. These individuals are often tasked with developing and deploying custom tools to identify vulnerabilities, test the security of systems, and ensure the integrity and safety of digital infrastructure. Pegasus Script is designed to cater to this audience by providing a simplified yet powerful platform for creating and customizing security tools.

## Basic Concepts

### Syntax

- Basic structure of a Pegasus Script program.
    ```
    BEGIN

    # Setting up variables
    VARIABLES
        x INTEGER;
        y INTEGER;
        result INTEGER;
    END

    # Assigning values to variables
    x = 10;
    y = 100;

    # Performing calculations
    result = x + y * y;

    # Conditional logic
    IF result <= 100 THEN
        # Output a string
        PUTS "random"
    ELSE
        # Output the result
        PUTS result
    END

    END
    ```
- Comments and documentation.

### Data Types

- Primitive data types (e.g., integers, strings, booleans).
    ```
    int # interger

    float # floating-point interger

    string # string

    bool # boolean

    char # character
    ```
- Complex data types (e.g., arrays).
    ```
    int[] # interger array

    float[] # floating-point interger array

    string[] # string array

    char[] # character array
    ```
### Variables and Constants

- Declaration and assignment.
    - Primitive data types (e.g., integers, strings, booleans).
        ```
        let x: int = 10; # interger

        let y: float = 10.0; # floating-point interger

        let z: string = "10"; # string

        let a: bool = false; # boolean

        let b: char = 'a'; # character
        ```
    - Complex data types (e.g., arrays).
        ```
        let x: int[] = {10, 11, 12}; # interger array

        let y: float[] = {10.0, 11.0, 12.0}; # floating-point interger array

        let z: string[] = {"Hello", "world"}; # string array

        let a: char[] = {'A', 'a'}; # character array
        ```
- Scope rules.
    - Global Scope: Variables declared outside of all functions or classes are in the global scope. They can be accessed from anywhere in the code.
        ```
        let x: int = 10; # global

        function foo() { 
            puts(x); # Accessible here
        }
        
        function bar() { 
            puts(x + 1); # Also Accessible here
        }
        ```
    - Local Scope: Variables declared within a function or a block are in the local scope of that function or block. They can only be accessed within the function or block where they are declared.
        ```
        function foo() { 
            let x: int = 10; # local to foo
            puts(x); 
        } 
        
        function bar() { 
            puts(x + 1); # This would be error because x is not accessible here
        }
        ```
## Control Structures

### Conditional Statements

- `if`, `else if`, `else`.
- Ternary operator.

### Loops

- `for`, `while`, `do-while`.
- Break and continue statements.

### Switch Statements

- Syntax and usage.
- Fall-through behavior.

## Functions and Modules

### Functions

- Defining and calling functions.
- Parameters and return values.

### Modules

- Importing and exporting modules.
- Namespace and package management.

## Object-Oriented Programming

### Classes and Objects

- Defining classes and objects.
- Constructors and destructors.

### Inheritance and Polymorphism

- Inheritance hierarchy.
- Overriding and overloading methods.

### Interfaces and Abstract Classes

- Defining interfaces.
- Implementing abstract classes.

## Error Handling

### Exceptions

- Throwing and catching exceptions.
- Custom exception types.

### Error Handling Techniques

- Try-catch blocks.
- Finally blocks.

## Standard Library

### Introduction to the Standard Library

- Overview of available modules.

### File I/O

- Reading from and writing to files.

### String Manipulation

- String concatenation, splitting, and searching.

### Mathematical Operations

- Basic arithmetic, trigonometric functions, etc.

### Date and Time

- Working with dates and times.

## Advanced Topics

### Concurrency and Parallelism

- Threads and tasks.
- Synchronization primitives.

### Networking

- Socket programming.
- HTTP requests.

### Security

- Encryption and hashing.
- Secure coding practices.

## Best Practices

### Code Style and Conventions

- Naming conventions.
- Code formatting.

### Performance Optimization

- Profiling and optimization techniques.

### Security Practices

- Input validation and sanitization.
- Secure coding guidelines.

## Appendices

### Glossary

- Definitions of key terms.

### References

- External resources and further reading.

### Change Log

- Version history and updates.

## Examples and Tutorials

### Hello World

- A simple "Hello, World!" program.

### Basic Data Types

- Demonstrating the use of different data types.

### Control Structures

- Examples of conditional statements and loops.

### Functions and Modules

- Creating and using custom functions and modules.
