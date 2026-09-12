package formatage

import (
	"testing"

	"tetris-optimizer/utils"
)

func TestGoodFormat_Valid(t *testing.T) {
	block := "....\n.##.\n.##.\n....\n" // 20 caractères (5*4)
	if !GoodFormat([]string{block}) {
		t.Fatalf("GoodFormat() = false, want true for a valid 4x4 block")
	}
}

func TestGoodFormat_InvalidLength(t *testing.T) {
	if GoodFormat([]string{"abc"}) {
		t.Fatalf("GoodFormat() = true, want false for a block that isn't 20 chars long")
	}
}

func TestGoodTetriomino_Valid(t *testing.T) {
	p := utils.Piece{
		Cells:  [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		Height: 2,
		Width:  2,
	}
	if !GoodTetriomino(p) {
		t.Fatalf("GoodTetriomino() = false, want true for a valid tetromino")
	}
}

func TestGoodTetriomino_WrongCellCount(t *testing.T) {
	p := utils.Piece{
		Cells:  [][2]int{{0, 0}, {0, 1}, {1, 0}}, // seulement 3 cellules
		Height: 2,
		Width:  2,
	}
	if GoodTetriomino(p) {
		t.Fatalf("GoodTetriomino() = true, want false when there aren't exactly 4 cells")
	}
}

func TestGoodTetriomino_TooSpreadOut(t *testing.T) {
	p := utils.Piece{
		Cells:  [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		Height: 4,
		Width:  4, // Height+Width = 8 > 5
	}
	if GoodTetriomino(p) {
		t.Fatalf("GoodTetriomino() = true, want false when Height+Width > 5")
	}
}
