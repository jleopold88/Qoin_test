package main

import (
	"fmt"
	"math/rand"
	"time"
)

func max(x []int) int {
	temp := 0
	index := -1
	for i, v := range x {
		if v > temp {
			temp = v
			index = i
		}
	}

	return index
}

func main() {
	scan := fmt.Scan
	print := fmt.Print
	println := fmt.Println
	var pemain, dadu int
	stop := false
	rand.Seed(time.Now().UnixNano())

	print("Pemain = ")
	scan(&pemain)
	print("Dadu = ")
	scan(&dadu)

	score := make([]int, pemain)
	done := make([]int, pemain)

	play := make([][]int, pemain)
	for i := range play {
		play[i] = make([]int, dadu, dadu*pemain)
		for j := range play[i] {
			play[i][j] = rand.Intn(6) + 1
		}
		print("Pemain #", i+1, "(", score[i], ") : ", play[i], "\n")
	}
	println("========================================================================")

	for !stop {
		for i, v := range play {
			length := len(v)
			for j := 0; j < length; j++ {
				if play[i][j] == 6 {
					score[i] += 1
					play[i][j] = play[i][len(play[i])-1]
					play[i] = play[i][:len(play[i])-1]
					length--
					j--
				} else if play[i][j] == 1 {
					play[i][j] = play[i][len(play[i])-1]
					play[i] = play[i][:len(play[i])-1]
					length--
					j--

					if i >= len(play)-1 {
						for z := 0; z < len(play)-1; z++ {
							if done[z] == 0 {
								play[z] = append(play[z], 0)
								break
							}
							if z+1 == len(play)-1 {
								stop = true
							}
						}
					} else {
						for z := i + 1; z != i; z++ {

							if done[z] == 0 {
								play[z] = append(play[z], 0)
								break
							}
							if z+1 == i {
								stop = true
								break
							}
							if z+1 > len(play)-1 {
								z = -1
							}
						}
					}
				}
			}
		}
		temp := 0
		for i, v := range play {
			if len(v) == 0 {
				done[i] = 1
			}
			if done[i] == 1 {
				temp++
			}
		}
		for i := range play {
			if done[i] == 1 {
				print("Pemain #", i+1, "(", score[i], ") : ", "_(Berhenti bermain karena tidak memiliki dadu)\n")
			} else {
				print("Pemain #", i+1, "(", score[i], ") : ")
				for j := range play[i] {
					if play[i][j] == 0 {
						print(1, " ")
					} else {
						print(play[i][j], " ")
					}
				}
				println()
			}
		}
		println("========================================================================")
		if temp == pemain-1 {
			break
		}

		println("===========================NEW ROLL===========================")
		for i := range play {
			for j := range play[i] {
				play[i][j] = rand.Intn(6) + 1
			}
			print("Pemain #", i+1, "(", score[i], ") : ", play[i], "\n")
		}
		println("===========================NEW ROLL===========================")

	}
	println("Pemenang adalah Pemain #", max(score)+1)

}
