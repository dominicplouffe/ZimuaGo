package main

import (
	"math"
)

func getBishopMobilitySquares(pieces []uint64, colorSquares uint64, oppSquares uint64) int {

	mobility := 0

	if len(pieces) == 0 {
		return mobility
	}

	for _, p := range pieces {
		newp := p
		if newp == 0 {
			break
		}
		for getFile(newp) != bbFILEA && getRank(newp) != bbRANK1 {
			newp = newp >> 9
			if newp&colorSquares != 0 {
				break
			}
			if newp&oppSquares != 0 {
				mobility++
				break
			}

			mobility++
		}
		newp = p
		for getFile(newp) != bbFILEH && getRank(newp) != bbRANK8 {
			newp = newp << 9
			if newp == 0 {
				break
			}
			if newp&colorSquares != 0 {
				break
			}
			if newp&oppSquares != 0 {
				mobility++
				break
			}

			mobility++
		}
		newp = p
		for getFile(newp) != bbFILEH && getRank(newp) != bbRANK1 {
			newp = newp >> 7
			if newp == 0 {
				break
			}
			if newp&colorSquares != 0 {
				break
			}
			if newp&oppSquares != 0 {
				mobility++
				break
			}

			mobility++
		}
		newp = p
		for getFile(newp) != bbFILEA && getRank(newp) != bbRANK8 {
			newp = newp << 7
			if newp == 0 {
				break
			}
			if newp&colorSquares != 0 {
				break
			}
			if newp&oppSquares != 0 {
				mobility++
				break
			}

			mobility++
		}
	}

	return mobility
}

func getRookMobilitySquares(pieces []uint64, colorRooks uint64, colorSquares uint64, oppSquares uint64) (int, bool) {

	mobility := 0
	connected := false

	if len(pieces) == 0 {
		return mobility, false
	}

	for _, p := range pieces {
		file := getFile(p)
		rank := getRank(p)

		newp := p
		for i := 0; i < 7; i++ {
			newp = newp << 1
			if newp == 0 {
				break
			}
			if getRank(newp) != rank {
				break
			}
			//Chech for connected rooks
			if newp&colorRooks != 0 {
				connected = true
				break
			}
			if newp&colorSquares != 0 {
				break
			}
			if newp&oppSquares != 0 {
				mobility++
				break
			}

			mobility++
		}
		newp = p
		for i := 0; i < 7; i++ {
			newp = newp >> 1
			if newp == 0 {
				break
			}
			if getRank(newp) != rank {
				break
			}

			//Chech for connected rooks
			if newp&colorRooks != 0 {
				connected = true
				break
			}

			if newp&colorSquares != 0 {
				break
			}
			if newp&oppSquares != 0 {
				mobility++
				break
			}
			mobility++
		}
		newp = p
		for i := 0; i < 7; i++ {
			newp = newp << 8
			if newp == 0 {
				break
			}
			if getFile(newp) != file {
				break
			}

			//Chech for connected rooks
			if newp&colorRooks != 0 {
				connected = true
				break
			}
			if newp&colorSquares != 0 {
				break
			}
			if newp&oppSquares != 0 {
				mobility++
				break
			}
			mobility++
		}
		newp = p
		for i := 0; i < 7; i++ {
			newp = newp >> 8
			if newp == 0 {
				break
			}
			if getFile(newp) != file {
				break
			}
			//Chech for connected rooks
			if newp&colorRooks != 0 {
				connected = true
				break
			}
			if newp&colorSquares != 0 {
				break
			}
			if newp&oppSquares != 0 {
				mobility++
				break
			}
			mobility++
		}
	}

	return mobility, connected
}

func getKnightMobilitySquares(pieces []uint64, colorSquares uint64) int {

	mobility := 0

	if len(pieces) == 0 {
		return mobility
	}

	moves := [8]int8{15, 17, 6, 10, -6, -10, -15, -17}

	for _, p := range pieces {
		file := getFile(p)
		rank := getRank(p)

		for _, m := range moves {
			if file == bbFILEA || file == bbFILEB {
				if m == 6 || m == -10 {
					continue
				}

			}
			if file == bbFILEA {
				if m == 15 || m == -17 {
					continue
				}

			}
			if file == bbFILEG || file == bbFILEH {
				if m == 10 || m == -10 {
					continue
				}
			}
			if file == bbFILEH {
				if m == 17 || m == -17 {
					continue
				}

			}

			if rank == bbRANK7 || rank == bbRANK8 {
				if m == 15 || m == 17 {
					continue
				}
			}
			if rank == bbRANK8 {
				if m == 6 || m == 10 {
					continue
				}
			}
			if rank == bbRANK1 || rank == bbRANK2 {
				if m == -15 || m == -17 {
					continue
				}
			}
			if rank == bbRANK1 {
				if m == -6 || m == -10 {
					continue
				}
			}

			var pos uint64 = 0
			if m > 0 {
				pos = p << m
			} else {
				pos = p >> int(math.Abs(float64(m)))
			}

			if pos&colorSquares == 0 {
				mobility++
			}
		}

	}

	return mobility
}

func getQueenMobilitySquares(pieces []uint64, colorSquares uint64, oppSquares uint64) int {

	mobilityBishop := getBishopMobilitySquares(pieces, colorSquares, oppSquares)
	mobilityRook, _ := getRookMobilitySquares(pieces, 0, colorSquares, oppSquares)

	mobility := mobilityBishop + mobilityRook

	return mobility
}
