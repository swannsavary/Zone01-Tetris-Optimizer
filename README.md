# Tetris Optimizer

A Go program that reads a list of tetrominoes from a text file and assembles them into the **smallest possible square**.

## Features

- Reads tetrominoes from a text file (`sample.txt`)
- Validates the file format and each tetromino
- Finds the smallest square grid that can hold all the pieces (backtracking)
- Prints the result with one letter per piece (`A`, `B`, `C`, ...) and `.` for empty cells

## Requirements

- Go 1.22.2 or later

## Usage

Put your tetrominoes in `sample.txt` at the root of the project, then run:

```bash
go run .
```

## Input format

- Each tetromino is 4 lines of 4 characters
- `#` is a block, `.` is an empty cell
- Tetrominoes are separated by one empty line
- The file must end with a newline

Example (`$` marks the end of each line):

```
...#$
...#$
...#$
...#$
$
....$
....$
..##$
..##$
```

## Output

```
ABB.
ABB.
A...
A...
```

Pieces are named `A`, `B`, `C`... in the order they appear in the file.

## How it works

1. The file is split into one block per tetromino and the format is checked
2. Empty rows and columns are removed from each tetromino
3. Each tetromino is converted into a `Piece` (letter, cell positions, height, width)
4. The smallest possible grid size is computed from the number of pieces (n² ≥ 4 × pieces)
5. Pieces are placed one by one with backtracking
6. If no solution exists, the grid size is increased by 1 and the process restarts

## Errors

If the file cannot be read, or if the format or a tetromino is invalid, the program prints `Error` and stops.

## Tests

```bash
go test ./...
```

## Project structure

```
.
├── formatage
│   ├── formatage.go
│   └── formatage_test.go
├── go.mod
├── main.go
├── README.md
├── sample.txt
└── utils
    ├── assemble.go
    ├── assemble_test.go
    ├── separation.go
    └── separation_test.go
```

## Author

swannsavary