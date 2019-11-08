package score

import "fmt"

type lottoGame struct {
	owner   string
	game    int
	numbers []int
}

func scoreGame(lg lottoGame, w []int) (string, int) {
	m, d := scoreEntery(lg.numbers, w)
	return fmt.Sprintf("%s wins a division %d on game #%d with matches %v in game %v", lg.owner, d, lg.game, m, lg.numbers), d
}

func scoreEntery(g []int, w []int) (m []int, d int) {
	for _, a := range w {
		if contains(g, a) {
			d++
			m = append(m, a)
		}
	}
	return m, 7 - d
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
