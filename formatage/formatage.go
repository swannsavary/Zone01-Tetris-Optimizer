package formatage

import "tetris-optimizer/utils"

// renvoie true si le string n est pas au format 4 lines de 5 collonnes
func GoodFormat(str []string) bool {
	for _, s := range str {
		if len(s) != 5*4 {
			return false
		}
	}
	return true
}

// renvoi true si c est bien un tetriomino valide
func GoodTetriomino(piece utils.Piece) bool {
	if len(piece.Cells) != 4 { //composer de 4 #
		return false
	}
	if piece.Height+piece.Width > 5 { // si les # ne sont pas coller
		return false
	}
	return true
}
