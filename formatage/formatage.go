package formatage

import "tetris-optimizer/utils"

func GoodFormat(str []string) bool {
	for _, s := range str {
		if len(s) != 5*4 {
			return false
		}
	}
	return true
}

func GoodTetriomino(piece utils.Piece) bool {
	if len(piece.Cells) != 4 {
		return false
	}
	if piece.Height+piece.Width > 5 {
		return false
	}
	return true
}
