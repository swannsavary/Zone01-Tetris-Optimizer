package utils

import "strings"

type Piece struct {
	Letter byte     //la lettre est par quoi on reconnait le tetriomino au lieu des # il y aurra la lettre
	Cells  [][2]int //position des lettres dans la piece sous la forme (ligne,collonne)
	Height int
	Width  int
}

// transform un string en fonction de la lettre qui lui est atribuer dans l'ordre des tetriomino en une Piece
// parcourt le string et à chaque fois qu'il tombe sur un des composant du tetriomino il ajout la position à cells
func ToPiece(s string, letter byte) Piece {
	lines := strings.Split(strings.Trim(s, "\n"), "\n")
	var cells [][2]int
	for r, line := range lines {
		for c, ch := range line {
			if byte(ch) == letter {
				cells = append(cells, [2]int{r, c})
			}
		}
	}
	return Piece{Letter: letter, Cells: cells, Height: len(lines), Width: len(lines[0])}
}

// resoud par recurence le placement de chaque tetriomino dans la grille
func solve(grid [][]byte, pieces []Piece, index int) bool {
	if index == len(pieces) { //initialisation si tout les pieces sont dans la grille return true
		return true
	}
	p := pieces[index]
	n := len(grid)
	//teste toutes les positions possible en fonction de la taille du tetriomino
	for r := 0; r <= n-p.Height; r++ {
		for c := 0; c <= n-p.Width; c++ {
			if canPlace(grid, p, r, c) {
				place(grid, p, r, c)
				if solve(grid, pieces, index+1) {
					return true
				}
				remove(grid, p, r, c) // backtrack
			}
		}
	}
	return false
}

// return true si on peut placer la piece dans la grille et false si elle est deja occuper
func canPlace(grid [][]byte, p Piece, r, c int) bool {
	for _, cell := range p.Cells {
		if grid[r+cell[0]][c+cell[1]] != '.' {
			return false
		}
	}
	return true
}

// place la piece dans la grille
func place(grid [][]byte, p Piece, r, c int) {
	for _, cell := range p.Cells {
		grid[r+cell[0]][c+cell[1]] = p.Letter
	}
}

// enleve la piece de la grille
func remove(grid [][]byte, p Piece, r, c int) {
	for _, cell := range p.Cells {
		grid[r+cell[0]][c+cell[1]] = '.'
	}
}

// testes pour chaques taille de grille si la grille est possible et renvoie la plus petite grille possible
func Assemble(pieces []Piece) [][]byte {
	n := 1
	for n*n < 4*len(pieces) {
		n++
	}
	for {
		grid := make([][]byte, n)
		for i := range grid {
			grid[i] = make([]byte, n)
			for j := range grid[i] {
				grid[i][j] = '.'
			}
		}
		if solve(grid, pieces, 0) {
			return grid
		}
		n++
	}
}
