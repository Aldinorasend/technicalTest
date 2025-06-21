package main

import (
	"fmt"
)

func denseRanking(scores []int, gitsScores []int) []int {

	var ranks []int
	for _, s := range scores {
		if len(ranks) == 0 || s != ranks[len(ranks)-1] {
			ranks = append(ranks, s)
		}
	}

	var result []int

	for _, g := range gitsScores {
		rank := 1
		for _, r := range ranks {
			if g < r {
				rank++
			} else {
				break
			}
		}
		result = append(result, rank)
	}

	return result
}

func main() {
	var n int
	fmt.Print("Jumlah skor leaderboard: ")
	fmt.Scan(&n)

	scores := make([]int, n)
	fmt.Println("Masukkan skor leaderboard (dari besar ke kecil):")
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	var m int
	fmt.Print("Jumlah skor GITS: ")
	fmt.Scan(&m)

	gitsScores := make([]int, m)
	fmt.Println("Masukkan skor GITS:")
	for i := 0; i < m; i++ {
		fmt.Scan(&gitsScores[i])
	}

	result := denseRanking(scores, gitsScores)

	fmt.Println("Peringkat GITS:")
	for _, r := range result {
		fmt.Print(r, " ")
	}
}
