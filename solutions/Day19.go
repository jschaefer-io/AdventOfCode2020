package solutions

import (
	"fmt"
	"regexp"
	"strings"
)

type Day19 struct {
	rules        map[string][]ruleGroup
	messages     []string
	maxMsgLength int
	dynResolver  map[string][]string
}

type rule interface {
	resolveToList(d *Day19) []string
}

type ruleGroup struct {
	rules        []rule
	resolveCount int
	count        int
}

func (r *ruleGroup) resolveToList(d *Day19) []string {
	list := make([][]string, 0)
	for _, r := range r.rules {
		list = append(list, r.resolveToList(d))
	}
	if len(list) == 1 {
		return list[0]
	}
	r.count++
	return buildStringVariations(0, list, d)
}

// literal rules
type literalRule struct {
	literal string
}

func (l literalRule) resolveToList(d *Day19) []string {
	return []string{l.literal}
}

// reference rules
type referenceRule struct {
	reference string
}

func (l referenceRule) resolveToList(d *Day19) []string {
	list := make([]string, 0)
	for _, g := range d.rules[l.reference] {
		rules := g.resolveToList(d)
		for _, r := range rules {
			if len(r) > d.maxMsgLength {
				continue
			}
			list = append(list, r)
		}

	}
	d.dynResolver[l.reference] = list
	return list
}

// Solution
func (d *Day19) init(s string) {
	d.dynResolver = make(map[string][]string)
	d.rules = make(map[string][]ruleGroup, 0)
	d.messages = make([]string, 0)
	splitG := regexp.MustCompile("((|\\r)\\n){2}")
	isLiteral := regexp.MustCompile("\"(\\w+)\"")
	groups := splitG.Split(s, -1)

	// rules
	for _, rule := range strings.Split(groups[0], "\n") {
		parts := strings.Split(strings.TrimSpace(rule), ":")
		rule := strings.TrimSpace(parts[0])
		groups := make([]ruleGroup, 0)

		if isLiteral.MatchString(parts[1]) {
			literal := strings.TrimSpace(strings.ReplaceAll(parts[1], "\"", ""))
			g := ruleGroup{}
			g.rules = append(g.rules, literalRule{literal})
			groups = append(groups, g)
		} else {
			for _, gs := range strings.Split(parts[1], "|") {
				g := ruleGroup{}
				for _, ref := range strings.Split(strings.TrimSpace(gs), " ") {
					g.rules = append(g.rules, referenceRule{ref})
				}
				groups = append(groups, g)
			}
		}
		d.rules[rule] = groups
	}

	// messages
	for _, msg := range strings.Split(groups[1], "\n") {
		msg := strings.TrimSpace(msg)
		length := len(msg)
		if d.maxMsgLength < length {
			d.maxMsgLength = length
		}
		d.messages = append(d.messages, msg)
	}
}

func buildStringVariations(index int, list [][]string, d *Day19) []string {
	res := make([]string, 0)
	for _, c := range list[index] {
		if index+1 >= len(list) {
			res = append(res, c)
		} else {
			for _, t := range buildStringVariations(index+1, list, d) {
				sum := c + t
				if len(sum) > d.maxMsgLength {
					continue
				}
				res = append(res, c+t)
			}
		}
	}
	return res
}

func (d *Day19) executeA() int {
	checkList := make(map[string]struct{})
	r := referenceRule{reference: "0"}
	for _, list := range r.resolveToList(d) {
		checkList[list] = struct{}{}
	}

	count := 0
	for _, msg := range d.messages {
		if _, ok := checkList[msg]; ok {
			count++
		}
	}
	return count
}

func (d *Day19) Handle(s string) ([]string, error) {
	d.init(s)

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	//results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
