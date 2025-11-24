# 👳🏾‍♂️ TauLang – Desi Tau Programming Language

> _"Sun liyo tau… programming bhi ab apne hisaab se hogi."_

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Tests](https://img.shields.io/badge/tests-passing-brightgreen)](https://github.com/ishaan1091/taulang/actions)

TauLang is a **fully-featured programming language interpreter** written in Go, inspired by the wisdom and humor of the great Desi Tau. While it looks like a parody language with its hilarious Hindi-English keywords, TauLang is a **real, working interpreter** that implements a complete programming language with modern features.

## 🌟 Features

-   ✅ **Complete Language Features**

    -   Variable bindings and assignments
    -   Integer, boolean, string, and null types
    -   Arithmetic and logical expressions
    -   Control flow (if/else, while loops)
    -   Break and continue statements
    -   First-class and higher-order functions
    -   Closures with lexical scoping
    -   Arrays and hash maps
    -   Index expressions and assignments

-   ✅ **Built-in Functions**

    -   `len()` - Get length of strings, arrays, or hash maps
    -   `first()` - Get first element of an array
    -   `last()` - Get last element of an array
    -   `push()` - Add element to array

-   ✅ **Developer Experience**
    -   REPL (Read-Eval-Print Loop) for interactive coding
    -   File execution support
    -   Comprehensive test coverage
    -   Clean, modular architecture

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/ishaan1091/taulang.git
cd taulang

# Build the interpreter
make build

# Run the REPL
./bin/taulang

# Or execute a file
./bin/taulang path/to/file.tau
```

### Requirements

-   Go 1.21 or higher
-   Make (optional, for using Makefile commands)

## 📖 Language Documentation

### Keywords

| TauLang         | English    | Description           |
| --------------- | ---------- | --------------------- |
| `sun_liyo_tau`  | `let`      | Variable declaration  |
| `tau_ka_jugaad` | `func`     | Function definition   |
| `agar_maan_lo`  | `if`       | Conditional statement |
| `na_toh`        | `else`     | Else clause           |
| `laadle_ye_le`  | `return`   | Return statement      |
| `jab_tak`       | `while`    | While loop            |
| `rok_diye`      | `break`    | Break statement       |
| `jaan_de`       | `continue` | Continue statement    |
| `ne_bana_diye`  | `=`        | Assignment operator   |
| `saccha`        | `true`     | Boolean true          |
| `jhootha`       | `false`    | Boolean false         |

### Data Types

#### Integers

```tau
sun_liyo_tau x ne_bana_diye 42;
sun_liyo_tau y ne_bana_diye -10;
```

#### Booleans

```tau
sun_liyo_tau isTrue ne_bana_diye saccha;
sun_liyo_tau isFalse ne_bana_diye jhootha;
```

#### Strings

```tau
sun_liyo_tau greeting ne_bana_diye "Namaste Tau!";
sun_liyo_tau message ne_bana_diye "Sun liyo tau";
```

#### Arrays

```tau
sun_liyo_tau numbers ne_bana_diye [1, 2, 3, 4, 5];
sun_liyo_tau mixed ne_bana_diye [1, "hello", saccha, 42];
```

#### Hash Maps

```tau
sun_liyo_tau person ne_bana_diye {
    "name": "Tau",
    "age": 50,
    "city": "Mumbai"
};
```

### Operators

#### Arithmetic

-   `+` Addition
-   `-` Subtraction
-   `*` Multiplication
-   `/` Division

#### Comparison

-   `==` Equal to
-   `!=` Not equal to
-   `>` Greater than
-   `<` Less than
-   `>=` Greater than or equal
-   `<=` Less than or equal

#### Logical

-   `!` Logical NOT

### Control Flow

#### If/Else Statements

```tau
agar_maan_lo (x > 10) {
    laadle_ye_le "Greater than 10";
} na_toh {
    laadle_ye_le "Less than or equal to 10";
}
```

#### While Loops

```tau
sun_liyo_tau i ne_bana_diye 0;
jab_tak (i < 10) {
    // Do something
    i ne_bana_diye i + 1;
}
```

#### Break and Continue

```tau
jab_tak (saccha) {
    agar_maan_lo (condition) {
        rok_diye;  // Break out of loop
    }
    agar_maan_lo (otherCondition) {
        jaan_de;  // Continue to next iteration
    }
}
```

### Functions

#### Function Definition

```tau
sun_liyo_tau add ne_bana_diye tau_ka_jugaad(a, b) {
    laadle_ye_le a + b;
};
```

#### Function Calls

```tau
sun_liyo_tau result ne_bana_diye add(5, 3);
```

#### Higher-Order Functions

```tau
sun_liyo_tau apply ne_bana_diye tau_ka_jugaad(fn, x) {
    laadle_ye_le fn(x);
};

sun_liyo_tau double ne_bana_diye tau_ka_jugaad(n) {
    laadle_ye_le n * 2;
};

sun_liyo_tau result ne_bana_diye apply(double, 5);
```

#### Closures

```tau
sun_liyo_tau makeCounter ne_bana_diye tau_ka_jugaad() {
    sun_liyo_tau count ne_bana_diye 0;
    laadle_ye_le tau_ka_jugaad() {
        count ne_bana_diye count + 1;
        laadle_ye_le count;
    };
};

sun_liyo_tau counter ne_bana_diye makeCounter();
counter();  // Returns 1
counter();  // Returns 2
```

### Arrays

#### Array Literals

```tau
sun_liyo_tau arr ne_bana_diye [1, 2, 3, 4, 5];
```

#### Array Indexing

```tau
sun_liyo_tau first ne_bana_diye arr[0];
sun_liyo_tau last ne_bana_diye arr[4];
```

#### Array Index Assignment

```tau
arr[0] ne_bana_diye 10;  // Modify existing element
arr[10] ne_bana_diye 20;  // Auto-extends array if needed
```

#### Built-in Array Functions

```tau
len(arr);        // Get array length
first(arr);      // Get first element
last(arr);       // Get last element
push(arr, 6);    // Add element to array (returns new array)
```

### Hash Maps

#### Hash Map Literals

```tau
sun_liyo_tau map ne_bana_diye {
    "one": 1,
    "two": 2,
    "three": 3,
    saccha: 4,
    jhootha: 5,
    1: 7,
    2: 8
};
```

#### Hash Map Indexing

```tau
sun_liyo_tau value ne_bana_diye map["one"];
sun_liyo_tau boolValue ne_bana_diye map[saccha];
```

#### Hash Map Index Assignment

```tau
map["newKey"] ne_bana_diye "newValue";  // Add new key
map["one"] ne_bana_diye 100;            // Update existing key
```

#### Hash Map Keys

Hash maps support various key types:

-   Strings: `"key"`
-   Integers: `1`, `2`, `42`
-   Booleans: `saccha`, `jhootha`

### Variable Assignment

#### Regular Assignment

```tau
sun_liyo_tau x ne_bana_diye 5;
x ne_bana_diye 10;  // Reassign variable
```

#### Index Assignment

```tau
sun_liyo_tau arr ne_bana_diye [1, 2, 3];
arr[0] ne_bana_diye 10;  // Modify array element

sun_liyo_tau map ne_bana_diye {"key": "value"};
map["key"] ne_bana_diye "newValue";  // Modify hash map
```

### Built-in Functions

#### `len(value)`

Returns the length of a string, array, or hash map.

```tau
len("hello");           // Returns 5
len([1, 2, 3]);         // Returns 3
len({"a": 1, "b": 2});  // Returns 2
```

#### `first(array)`

Returns the first element of an array, or `null` if empty.

```tau
first([1, 2, 3]);  // Returns 1
first([]);          // Returns null
```

#### `last(array)`

Returns the last element of an array, or `null` if empty.

```tau
last([1, 2, 3]);  // Returns 3
last([]);          // Returns null
```

#### `push(array, element)`

Returns a new array with the element added to the end.

```tau
sun_liyo_tau arr ne_bana_diye [1, 2, 3];
sun_liyo_tau newArr ne_bana_diye push(arr, 4);  // Returns [1, 2, 3, 4]
```

## 💻 Example Programs

### Hello World

```tau
sun_liyo_tau message ne_bana_diye "Namaste Tau!";
message;
```

### Factorial Function

```tau
sun_liyo_tau factorial ne_bana_diye tau_ka_jugaad(n) {
    agar_maan_lo (n <= 1) {
        laadle_ye_le 1;
    } na_toh {
        laadle_ye_le n * factorial(n - 1);
    }
};

factorial(5);  // Returns 120
```

### Fibonacci Sequence

```tau
sun_liyo_tau fibonacci ne_bana_diye tau_ka_jugaad(n) {
    agar_maan_lo (n <= 1) {
        laadle_ye_le n;
    } na_toh {
        laadle_ye_le fibonacci(n - 1) + fibonacci(n - 2);
    }
};

fibonacci(10);  // Returns 55
```

### Working with Arrays

```tau
sun_liyo_tau numbers ne_bana_diye [1, 2, 3, 4, 5];
sun_liyo_tau sum ne_bana_diye 0;
sun_liyo_tau i ne_bana_diye 0;

jab_tak (i < len(numbers)) {
    sum ne_bana_diye sum + numbers[i];
    i ne_bana_diye i + 1;
}

sum;  // Returns 15
```

### Working with Hash Maps

```tau
sun_liyo_tau person ne_bana_diye {
    "name": "Tau",
    "age": 50,
    "city": "Mumbai"
};

person["name"];  // Returns "Tau"
person["age"];   // Returns 50

person["country"] ne_bana_diye "India";  // Add new key
person["age"] ne_bana_diye 51;            // Update existing key
```

## 🏗️ Architecture

TauLang is built with a clean, modular architecture:

```
taulang/
├── ast/          # Abstract Syntax Tree nodes
├── evaluator/    # Expression and statement evaluation
├── lexer/        # Tokenization (lexical analysis)
├── object/       # Runtime objects and environment
├── parser/       # Parsing (syntax analysis)
├── repl/         # Read-Eval-Print Loop
└── token/        # Token definitions
```

### How It Works

1. **Lexer**: Converts source code into tokens
2. **Parser**: Builds an Abstract Syntax Tree (AST) from tokens
3. **Evaluator**: Traverses the AST and executes the program
4. **Object System**: Manages runtime objects and environment

## 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Generate HTML coverage report
make test-coverage-html
```

## 📝 Development

### Building

```bash
# Build the interpreter
make build

# Build and check for errors
make build-check
```

### Project Structure

-   **Lexer**: Tokenizes source code
-   **Parser**: Parses tokens into AST
-   **Evaluator**: Executes AST nodes
-   **Object**: Runtime representation of values
-   **REPL**: Interactive shell for TauLang

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

-   Inspired by the wisdom and humor of Desi Tau
-   Built following the principles from "Writing An Interpreter In Go" by Thorsten Ball
-   Thanks to all contributors and the open-source community

## 🌐 Website

Visit the [TauLang website](https://ishaan1091.github.io/taulang) for an interactive experience with live examples, documentation, and more!

The website is automatically deployed to GitHub Pages whenever changes are pushed to the repository.

## 🔗 Links

-   [🌐 Website](https://ishaan1091.github.io/taulang) - Interactive landing page
-   [Documentation](#-language-documentation)
-   [Examples](#-example-programs)
-   [GitHub Repository](https://github.com/ishaan1091/taulang)
-   [Issues](https://github.com/ishaan1091/taulang/issues)

---

**Made with ❤️ and lots of "Sun liyo tau"**
