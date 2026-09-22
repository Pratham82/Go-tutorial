package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of boolst
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	fileToIterate := cb[file]

	for _, v := range fileToIterate {
		if v {
			count++
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	files := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	// for _, v := range files {
	// 	for i, v := range cb[v] {
	// 		if rank == i+1 && v {
	// 			count++
	// 		}
	// 	}
	// }
	//

	if rank < 1 || rank > 8 {
		return 0
	}

	for _, v := range files {
		if cb[v][rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

	for _, file := range cb {
		count += len(file)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	// count := 0
	// files := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	// for _, v := range files {
	// 	for _, v := range cb[v] {
	// 		if v {
	// 			count++
	// 		}
	// 	}
	// }

	// return count

	count := 0

	for _, file := range cb {
		for _, v := range file {
			if v {
				count++
			}
		}
	}

	return count

}
