package characteristics

import (
	"bufio"
	"fmt"
	"math/rand"
	"mods/structure"
	"os"
	"time"
)

func Get_xp(p1_xp int) int {
	rand.Seed(time.Now().UnixNano())
	min := 15
	max := 50
	xp := (rand.Intn(max+1-min) + min)
	p1_xp += xp
	fmt.Printf("\033[97;3mVous avez gagné %v XP\n",xp)
	return p1_xp
}

func Levelup(p1 *structure.Character, xp int) int {
	scanner := bufio.NewScanner(os.Stdin)

	if xp >= 100 {
		p1.Str += 2
		p1.Int += 2
		xp -= 100
		fmt.Println("LEVEL UP")
		scanner.Scan()
		scanner.Text()
	}
	return xp
}
