package main

import (
	"os"
	"tetris-optimizer/formatage"
	"tetris-optimizer/utils"
)

func main() {
	file, err := os.ReadFile("sample.txt")
	if err != nil {
		os.Stdout.WriteString("Error\n")
		return
	}

	//prend les donnees du fichier et on separe chaque pieces dans un slice de string alltriomino
	file2 := utils.SliceByteToString(file)
	if !formatage.GoodFormat(file2) {
		os.Stdout.WriteString("Error\n")
		return
	}
	file2 = utils.SupColumn(file2, utils.SupColumnOfString(file2))
	alltriomino := utils.SupRowOfString(file2)

	//on prend chaque triomino et on le transform dans une struct
	var pieces []utils.Piece
	for i, tetriomino := range alltriomino {
		letter := byte('A' + i)
		pi := utils.ToPiece(tetriomino, letter)
		if !formatage.GoodTetriomino(pi) {
			os.Stdout.WriteString("Error\n")
			return

		}
		pieces = append(pieces, pi)
	}

	//on place chaque piece au bonne endroit puis on affiche le resultat
	for _, piece := range utils.Assemble(pieces) {
		for _, p := range piece {
			os.Stdout.WriteString(string(p))
		}
		os.Stdout.WriteString("\n")
	}
}
