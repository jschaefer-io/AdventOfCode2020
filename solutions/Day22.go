package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day22 struct {
	a queue
	b queue
}

type queue struct {
	cards []int
}

func (q queue) length() int {
	return len(q.cards)
}

func (q *queue) enqueue(cards ...int) {
	q.cards = append(q.cards, cards...)
}

func (q *queue) pop() int {
	res := q.cards[0]
	q.cards = q.cards[1:]
	return res
}

type history struct {
	history [][2][]int
}

func (h *history) addToHistory(a, b queue) {
	history := [2][]int{
		a.cards,
		b.cards,
	}
	h.history = append(h.history, history)
}

func (h *history) compareIntSlice(a, b *[]int) bool {
	if len(*a) != len(*b) {
		return false
	}
	for i, aV := range *a {
		if aV != (*b)[i] {
			return false
		}
	}
	return true
}

func (h *history) checkHistory(a, b queue) bool {
	for _, history := range h.history {
		if !h.compareIntSlice(&a.cards, &history[0]) && !h.compareIntSlice(&b.cards, &history[1]) {
			continue
		}
		return true
	}
	return false
}

func (d *Day22) init(s string) error {
	splitG := regexp.MustCompile("((|\\r)\\n){2}")
	players := make([]queue, 2)
	for n, player := range splitG.Split(s, -1) {
		cards := strings.Split(player, "\n")
		player := queue{
			cards: make([]int, 0),
		}
		for _, card := range cards[1:] {
			num, err := strconv.Atoi(strings.TrimSpace(card))
			if err != nil {
				return err
			}
			player.enqueue(num)
		}
		players[n] = player
	}
	d.a = players[0]
	d.b = players[1]
	return nil
}

func (d Day22) countWinnerScore(deck queue) int {
	res := 0
	count := deck.length()
	for i, n := range deck.cards {
		res += (count - i) * n
	}
	return res
}

func (d Day22) executeA() int {
	for d.a.length() > 0 && d.b.length() > 0 {
		a := d.a.pop()
		b := d.b.pop()
		if a == b {
			panic(fmt.Sprintf("%d  equals %d", a, b))
		} else if a > b {
			d.a.enqueue(a, b)
		} else if a < b {
			d.b.enqueue(b, a)
		}
	}
	if d.a.length() == 0 {
		return d.countWinnerScore(d.b)
	}
	return d.countWinnerScore(d.a)
}

func (d *Day22) playRecursiveGame(a, b queue) (bool, queue) {
	history := history{history: make([][2][]int, 0)}
	for a.length() > 0 && b.length() > 0 {
		if history.checkHistory(a, b) {
			return true, a
		}
		history.addToHistory(a, b)

		cA := a.pop()
		cB := b.pop()
		if a.length() >= cA && b.length() >= cB {
			cpyA := make([]int, cA)
			cpyB := make([]int, cB)
			copy(cpyA, a.cards[:cA])
			copy(cpyB, b.cards[:cB])
			nA := queue{cards: cpyA}
			nB := queue{cards: cpyB}
			winA, _ := d.playRecursiveGame(nA, nB)
			if winA {
				a.enqueue(cA, cB)
			} else {
				b.enqueue(cB, cA)
			}
		} else {
			if cA == cB {
				panic(fmt.Sprintf("%d  equals %d", a, b))
			} else if cA > cB {
				a.enqueue(cA, cB)
			} else if cA < cB {
				b.enqueue(cB, cA)
			}
		}
	}
	if a.length() > 0 {
		return true, a
	}
	return false, b
}

func (d *Day22) executeB() int {
	_, winner := d.playRecursiveGame(d.a, d.b)
	return d.countWinnerScore(winner)
}

func (d *Day22) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
