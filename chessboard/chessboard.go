package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var counter int
	for _, occupied := range cb[file] {
		if occupied {
			counter++
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var count int

	if rank >= 1 && rank <= 8 {
		for _, v := range cb {
			if v[rank-1] {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var counter int
	for _, file := range cb {
		counter += len(file)
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var counter int
	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				counter++
			}
		}
	}
	return counter
}
