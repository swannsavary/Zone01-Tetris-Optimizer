package utils

import (
	"reflect"
	"testing"
)

// Deux tetrominoes de référence utilisés dans plusieurs tests :
// blockA = un carré O (2x2) centré dans la grille 4x4
// blockB = une barre I verticale
const blockA = "....\n.##.\n.##.\n....\n"
const blockB = ".#..\n.#..\n.#..\n.#..\n"

func TestSliceByteToString(t *testing.T) {
	input := blockA + "\n" + blockB // les deux blocs séparés par une ligne vide
	got := SliceByteToString([]byte(input))
	want := []string{blockA, blockB}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SliceByteToString() = %#v, want %#v", got, want)
	}
}

func TestSliceByteToString_SingleBlock(t *testing.T) {
	got := SliceByteToString([]byte(blockA))
	want := []string{blockA}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SliceByteToString() = %#v, want %#v", got, want)
	}
}

func TestSupColumnOfString(t *testing.T) {
	got := SupColumnOfString([]string{blockA, blockB})
	want := [][]int{{1, 2}, {1, 1}} // colonnes occupées par le # dans chaque bloc

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SupColumnOfString() = %v, want %v", got, want)
	}
}

func TestSupColumn(t *testing.T) {
	nbs := [][]int{{1, 2}, {1, 1}}
	got := SupColumn([]string{blockA, blockB}, nbs)
	want := []string{
		"..\n##\n##\n..\n", // colonnes 1-2 de blockA conservées
		"#\n#\n#\n#\n",     // colonne 1 de blockB conservée
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SupColumn() = %#v, want %#v", got, want)
	}
}

func TestSupRowOfString(t *testing.T) {
	trimmed := []string{"..\n##\n##\n..\n", "#\n#\n#\n#\n"}
	got := SupRowOfString(trimmed)
	want := []string{"AA\nAA\n", "B\nB\nB\nB\n"} // lignes vides supprimées, # -> lettre

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SupRowOfString() = %#v, want %#v", got, want)
	}
}

func TestNbsautdeligne(t *testing.T) {
	got := nbsautdeligne("A\nB\nC\n")
	if got != 3 {
		t.Fatalf("nbsautdeligne() = %d, want 3", got)
	}
}

func TestNbCaraByRow(t *testing.T) {
	got := nbCaraByRow("ABCD\nXYZ")
	if got != 5 {
		t.Fatalf("nbCaraByRow() = %d, want 5", got)
	}
}
