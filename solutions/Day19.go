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
	resolveToRegex(d *Day19) string
}

type ruleGroup struct {
	rules []rule
}

func (r *ruleGroup) resolveToRegex(d *Day19) string {
	list := make([]string, 0)
	for _, r := range r.rules {
		list = append(list, r.resolveToRegex(d))
	}
	if len(list) == 1 {
		return list[0]
	}
	return "(" + strings.Join(list, "") + ")"
}

// literal rules
type literalRule struct {
	literal string
}

func (l literalRule) resolveToRegex(d *Day19) string {
	return l.literal
}

// reference rules
type referenceRule struct {
	reference string
}

func (l referenceRule) resolveToRegex(d *Day19) string {
	list := make([]string, 0)
	for _, r := range d.rules[l.reference] {
		list = append(list, r.resolveToRegex(d))
	}
	if len(list) == 1 {
		return list[0]
	}
	return "(" + strings.Join(list, "|") + ")"
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

func (d *Day19) getRegex(num int) regexp.Regexp {
	r := referenceRule{reference: "0"}
	exp := fmt.Sprintf("^%s$", r.resolveToRegex(d))
	regex := regexp.MustCompile(exp)
	return *regex
}

func (d *Day19) execute(rule int) int {
	exp := d.getRegex(rule)
	count := 0
	for _, msg := range d.messages {
		if exp.MatchString(msg) {
			count++
		}
	}
	return count
}

func (d *Day19) executeA() int {
	return d.execute(0)
}

func (d *Day19) executeB() int {
	count := 5

	// Update rule 8 to a finite recursion
	group8 := make([]ruleGroup, 0)
	for i := 1; i <= count; i++ {
		r := ruleGroup{}
		for u := 0; u < i; u++ {
			r.rules = append(r.rules, referenceRule{"42"})
		}
		group8 = append(group8, r)
	}
	d.rules["8"] = group8

	// Update rule 11 to a finite recursion
	group11 := make([]ruleGroup, 0)
	for i := 1; i <= count; i++ {
		r := ruleGroup{}
		for u := 0; u < i; u++ {
			r.rules = append(r.rules, referenceRule{"42"})
		}
		for u := 0; u < i; u++ {
			r.rules = append(r.rules, referenceRule{"31"})
		}
		group11 = append(group11, r)
	}
	d.rules["11"] = group11

	return d.execute(0)
}

func (d *Day19) Handle(s string) ([]string, error) {
	d.init(s)
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
