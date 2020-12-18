package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day18 struct {
	expressions       []tokenGroup
	properExpressions []tokenGroup
}

type token interface {
	getRaw() string
	getValue() int
}

// group
type tokenGroup struct {
	tokens []token
}

func (t tokenGroup) getValue() int {
	count := len(t.tokens)
	if count == 0 {
		return 0
	}
	res := t.tokens[0].getValue()
	for i := 1; i < count; i++ {
		switch t.tokens[i].getRaw() {
		case "*":
			res *= t.tokens[i+1].getValue()
		case "+":
			res += t.tokens[i+1].getValue()
		}
	}
	return res
}

func (t tokenGroup) getRaw() string {
	return "g"
}

// literals
type tokenLiteral struct {
	raw string
}

func (t tokenLiteral) getValue() int {
	res, _ := strconv.Atoi(t.raw)
	return res
}

func (t tokenLiteral) getRaw() string {
	return t.raw
}

func (d *Day18) buildExpression(start int, length int, tokens []string, preferAddition bool) (tokenGroup, int) {
	g := tokenGroup{tokens: make([]token, 0)}

	var i int
	for i = start; i < length; i++ {
		switch tokens[i] {
		case "(":
			group, offset := d.buildExpression(i+1, length, tokens, preferAddition)
			g.tokens = append(g.tokens, group)
			i = offset
		case ")":
			if preferAddition {
				g = d.preferAddition(g)
			}
			return g, i
		default:
			token := tokenLiteral{tokens[i]}
			g.tokens = append(g.tokens, token)
		}
	}

	if preferAddition {
		g = d.preferAddition(g)
	}

	return g, i
}

func (d *Day18) preferAddition(expression tokenGroup) tokenGroup {
	g := expression.tokens

	changed := true
	for changed {
		changed = false
		newG := make([]token, 0)
		for i := 1; i < len(g); i += 2 {
			if g[i].getRaw() == "+" && !changed {
				newG = append(newG, tokenGroup{tokens: []token{
					g[i-1],
					g[i],
					g[i+1],
				}})
				i++
				changed = true
			} else {
				newG = append(newG, g[i-1], g[i])
			}
		}
		if changed {
			g = newG
		}
	}

	return tokenGroup{tokens: g}
}

func (d *Day18) init(s string) {
	d.expressions = make([]tokenGroup, 0)
	d.properExpressions = make([]tokenGroup, 0)
	for _, expressions := range strings.Split(s, "\n") {
		tokens := strings.Split(strings.TrimSpace(strings.ReplaceAll(expressions, " ", "")), "")
		a, _ := d.buildExpression(0, len(tokens), tokens, false)
		b, _ := d.buildExpression(0, len(tokens), tokens, true)
		d.expressions = append(d.expressions, a)
		d.properExpressions = append(d.properExpressions, b)
	}
}

func (d *Day18) execute(expressions []tokenGroup) int {
	sum := 0
	for _, exp := range expressions {
		sum += exp.getValue()
	}
	return sum
}

func (d *Day18) Handle(s string) ([]string, error) {
	d.init(s)

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.execute(d.expressions)))
	results = append(results, fmt.Sprintf("%d", d.execute(d.properExpressions)))
	return results, nil
}
