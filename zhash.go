package main

import (
	"github.com/dominicplouffe/chess"
)

//ZimuaHash Hash to store position
type ZimuaHash struct {
	score    int
	pos      chess.Position
	flag     string
	depth    int
	move     chess.Move
	inserted int
	beta     int
	alpha    int
}

func (zg ZimuaGame) store(p chess.Position, alpha int, beta int, score int, move chess.Move, depth int) {
	flag := "E"

	if score <= alpha {
		flag = "U"
	} else if score >= beta {
		flag = "L"
	}

	h := p.Hash()

	item := zg.hashTable[h]

	if item.score != 0 {
		if item.depth > depth {
			return
		}
		if flag != "E" && item.depth == depth {
			return
		}
	}

	zg.hashTable[h] = ZimuaHash{
		score:    score,
		pos:      p,
		flag:     flag,
		depth:    depth,
		move:     move,
		inserted: 1,
		alpha:    alpha,
		beta:     beta,
	}
}

func (zg ZimuaGame) lookup(p chess.Position, alpha int, beta int, depth int, legalMoves []MoveScore) (*ZimuaHash, bool) {
	h := p.Hash()

	item := zg.hashTable[h]

	if item.inserted != 1 {
		return nil, false
	}

	found := false
	for _, mv := range legalMoves {
		if mv.move == item.move {
			found = true
			break
		}
	}

	if found && item.depth == depth {
		flag := "E"

		if item.score <= alpha {
			flag = "U"
		} else if item.score >= beta {
			flag = "L"
		}

		if item.flag == flag {
			return &item, true
		}
	}

	return nil, false
}

// def lookup(self, board, legal_moves):
// h = polyglot.zobrist_hash(board)

// item = self.table.get(h, None)
// if item is None:
// 	return

// if item.move is None or item.move in legal_moves:
// 	return item

// return None
