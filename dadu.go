package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func check(s map[string]int) bool {
	player_active := 0
	for i := 1; i <= len(s); i++ {
		iStr := strconv.Itoa(i)
		if s["dice_players"+iStr] > 0 {
			player_active++
		}
	}
	fmt.Println("")
	fmt.Println("ini player aktif = ", player_active)
	if player_active > 1 {
		return true
	}
	return false
}

func main() {
	player := 3                      // total players
	round := 4                       // round
	value := []int{1, 2, 3, 4, 5, 6} // value of dice

	// make player
	var players = map[string]int{}
	for i := 1; i <= player; i++ {
		iStr := strconv.Itoa(i)
		players["player"+iStr] = 0
	}

	var dice = map[string]int{}
	for i := 1; i <= player; i++ {
		iStr := strconv.Itoa(i)
		dice["dice_players"+iStr] = round
	}

	rand.Seed(time.Now().UnixMicro()) // make random source
	for i := 1; check(dice); i++ {
		fmt.Println(strings.Repeat("===", 15))
		fmt.Println(strings.Repeat(" ", 18), "Round", i)
		fmt.Println(strings.Repeat("===", 15))
		for j := 1; j <= len(players); j++ {
			random := rand.Int() % len(value) // rand.Int need source. if no source = default
			appear := value[random]
			fmt.Println("angka yang keluar adalah : ", appear)
			fmt.Println("player"+strconv.Itoa(j), "mendapatkan nilai", appear)
			fmt.Println("")
			if appear == 6 {
				players["player"+strconv.Itoa(j)]++
				dice["dice_players"+strconv.Itoa(j)]--
			} else if appear == 1 {
				dice["dice_players"+strconv.Itoa(j)]--
				if j == player {
					dice["dice_players"+strconv.Itoa(j-1)]++
				} else {
					dice["dice_players"+strconv.Itoa(j+1)]++
				}
			}
		}
		fmt.Print("sisa dadu = ")
		for i := 1; i <= len(dice); i++ {

			if dice["dice_players"+strconv.Itoa(i)] < 0 {
				fmt.Print("player"+strconv.Itoa(i), ": 0, ")
			} else {
				fmt.Print("player"+strconv.Itoa(i), "  : ", dice["dice_players"+strconv.Itoa(i)], ", ")
			}
		}
	}
	fmt.Print("hasil = ")
	for i := 1; i <= len(players); i++ {
		fmt.Print("player"+strconv.Itoa(i), " : ", players["player"+strconv.Itoa(i)], ",")
	}
	winner := 0
	winner_name := ""
	for k, v := range players {
		if v > winner {
			winner = v
			winner_name = k
		}
	}
	fmt.Println("")
	fmt.Println("pemenangnya adalah : ", winner_name, "dengan score", winner)
}
