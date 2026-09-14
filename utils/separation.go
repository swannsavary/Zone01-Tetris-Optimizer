package utils

// change un slice de byte en un slice de string tout les 4 \n
func SliceByteToString(bytes []byte) []string {
	var result []string
	var tetromino string
	nbsautdeligne := 0
	for i := 0; i < len(bytes); i++ {
		tetromino += string(bytes[i])
		if bytes[i] == 10 {
			nbsautdeligne++
		}
		if nbsautdeligne == 4 {
			result = append(result, tetromino)
			tetromino = ""
			nbsautdeligne = 0
			i++
		}
	}
	return result
}

// supprime toutes les string ou il n'y a que des "." des un []string et remplace tout les # par la lettre dans l'ordre du []string
func SupRowOfString(str []string) []string {
	var result []string
	var tetriomino []rune
	var tetriominos string
	isTetriomino := false
	nb := 0
	for _, s := range str {
		isTetriomino = false
		for _, cara := range s {
			if cara != '.' && cara != '\n' {
				isTetriomino = true
			}
			if cara == '#' {
				tetriomino = append(tetriomino, rune(int('A')+nb))
			} else {
				tetriomino = append(tetriomino, cara)
			}
			if cara == '\n' {
				if isTetriomino {
					tetriominos += string(tetriomino)
				}
				tetriomino = []rune{}
				isTetriomino = false
			}
		}
		if len(tetriominos) != 0 {
			result = append(result, tetriominos)
			tetriominos = ""
		}
		nb++
	}
	return result
}

// note la position du debut du tetriomino et de sa fin pour supprimer les collonnes ou il n'y a que des points
func SupColumnOfString(str []string) [][]int {
	ajoutColumn := false
	var result [][]int
	debut := 0
	isdebut := true
	fin := 0
	for _, s := range str {
		isdebut = true
		for index, cara := range s {
			ajoutColumn = false
			if cara == '\n' {
				break
			}
			for i := 0; i < 4; i++ {
				if s[index+i*5] != '.' && s[index+i*5] != '\n' {
					ajoutColumn = true
				}
			}
			if ajoutColumn {
				if isdebut {
					debut = index
					isdebut = false
				}
				fin = index
			}
		}
		result = append(result, []int{debut, fin})
		debut = 0
		fin = 0
	}
	return result
}

// supprime la collonne ou il n'y a que des "." nbs = SupColumnOfString()
func SupColumn(str []string, nbs [][]int) []string {
	var result []string
	var tetriomino []rune
	for index, s := range str {
		column := nbsautdeligne(s)
		row := nbCaraByRow(s)
		for c := 0; c < column; c++ { //column
			for r := 0; r < row; r++ { //row
				if nbs[index][0] <= r && r <= nbs[index][1] {
					tetriomino = append(tetriomino, rune(s[c*row+r]))
				}
			}
			tetriomino = append(tetriomino, '\n')
		}
		result = append(result, string(tetriomino))
		tetriomino = []rune{}
	}
	return result
}

// return the number of \n in a string
func nbsautdeligne(s string) int {
	var count int
	for _, cara := range s {
		if cara == '\n' {
			count++
		}
	}
	return count
}

// return the number of cara before a \n
func nbCaraByRow(s string) int {
	var count int
	for _, cara := range s {
		if cara == '\n' {
			break
		}
		count++
	}
	return count + 1
}
