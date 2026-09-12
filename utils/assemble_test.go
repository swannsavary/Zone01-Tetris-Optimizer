package utils

import (
	"reflect"
	"testing"
)

func TestToPiece_Square(t *testing.T) {
	got := ToPiece("AA\nAA\n", 'A')
	want := Piece{
		Letter: 'A',
		Cells:  [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		Height: 2,
		Width:  2,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ToPiece() = %+v, want %+v", got, want)
	}
}

func TestToPiece_VerticalLine(t *testing.T) {
	got := ToPiece("B\nB\nB\nB\n", 'B')
	want := Piece{
		Letter: 'B',
		Cells:  [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		Height: 4,
		Width:  1,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ToPiece() = %+v, want %+v", got, want)
	}
}

func TestAssemble_SinglePiece(t *testing.T) {
	square := Piece{
		Letter: 'A',
		Cells:  [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		Height: 2,
		Width:  2,
	}

	got := Assemble([]Piece{square})
	want := [][]byte{
		{'A', 'A'},
		{'A', 'A'},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Assemble() = %v, want %v", toStrings(got), toStrings(want))
	}
}

func TestAssemble_TwoPieces(t *testing.T) {
	square := Piece{
		Letter: 'A',
		Cells:  [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		Height: 2,
		Width:  2,
	}
	line := Piece{
		Letter: 'B',
		Cells:  [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		Height: 4,
		Width:  1,
	}

	got := Assemble([]Piece{square, line})
	want := [][]byte{
		{'A', 'A', 'B', '.'},
		{'A', 'A', 'B', '.'},
		{'.', '.', 'B', '.'},
		{'.', '.', 'B', '.'},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Assemble() = %v, want %v", toStrings(got), toStrings(want))
	}
}

// petit helper pour des messages d'erreur lisibles
func toStrings(grid [][]byte) []string {
	var out []string
	for _, row := range grid {
		out = append(out, string(row))
	}
	return out
}
