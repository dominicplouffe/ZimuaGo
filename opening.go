package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
)

// Opening Move definition
type Opening struct {
	code  string
	title string
	moves []string
}

var openings []Opening

func buildOpenings() []Opening {

	var opens []Opening
	r := csv.NewReader(bytes.NewBuffer(csvData))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, row := range records {
		if i == 0 {
			continue
		}

		code := row[0]
		title := row[1]
		pgn := strings.Split(row[2], " ")
		var moves []string

		for _, m := range pgn {
			matched, _ := regexp.MatchString(`\.`, m)
			if matched {
				if m[2:6][0] == '.' {
					m = m[3:7]
				} else {
					m = m[2:6]
				}
			}

			moves = append(moves, m)
		}

		open := Opening{
			code:  code,
			title: title,
			moves: moves,
		}
		opens = append(opens, open)
	}

	return opens
}

func findMove(moves []string) (string, bool) {
	var nextMoves []string

	if len(moves) == 0 {
		for _, open := range openings {
			nextMoves = append(nextMoves, open.moves[0])
		}
	} else {
		for _, open := range openings {
			if len(open.moves) <= len(moves) {
				continue
			}

			good := true
			idx := 0
			for i, m := range moves {
				if open.moves[i] != m {
					good = false
					break
				}
				idx++
			}

			if good {
				nextMoves = append(nextMoves, open.moves[idx])
			}
		}
	}

	if len(nextMoves) == 0 {
		return "", false
	}

	min := 0
	max := len(nextMoves)
	idx := rand.Intn(max-min) + min

	return nextMoves[idx], true
}

func init() {
	fmt.Println("init opening")
	openings = buildOpenings()

}
