# N-P-R-A Language 

This project is a simple interpreter for a custom programming language, written in Go. The language aims to have Python-like functionality but uses C++-style syntax. The goal is to tokenize and parse mathematical expressions as an initial step toward building a complete programming language.

## Features

- **Lexical Analysis (Lexer):** Tokenizes input expressions, recognizing identifiers, numbers, and operators.
- **Parsing (Parser):** Parses the tokenized expressions to evaluate basic mathematical operations (addition, subtraction, multiplication, division).
- **Expression Evaluation:** Supports simple mathematical expressions with integer values.

## Getting Started

### Prerequisites

- **Go**: Make sure Go is installed. You can download it [here](https://go.dev/dl/).
- **Git** (optional): For managing version control if you're collaborating.

### Installation

1. Clone the repository or download the code:

   ```bash
   git clone https://github.com/yourusername/custom-lang-go.git
   cd custom-lang-go
