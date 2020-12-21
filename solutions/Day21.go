package solutions

import (
	"fmt"
	"strings"
)

type Day21 struct {
	foods       []food
	ingredients map[string]struct{}
	allergens   map[string]struct{}
	ident       map[string]string
}

type food struct {
	ingredients map[string]struct{}
	allergens   map[string]struct{}
}

func (d *Day21) init(s string) {
	d.ingredients = make(map[string]struct{})
	d.allergens = make(map[string]struct{})
	d.ident = make(map[string]string)
	d.foods = make([]food, 0)
	for _, foodString := range strings.Split(s, "\n") {
		g := strings.Split(foodString, "(contains")
		foodInstance := food{
			ingredients: make(map[string]struct{}),
			allergens:   make(map[string]struct{}),
		}

		for _, ingredient := range strings.Split(strings.TrimSpace(g[0]), " ") {
			value := strings.TrimSpace(ingredient)
			foodInstance.ingredients[value] = struct{}{}
			d.ingredients[value] = struct{}{}
		}

		for _, allergen := range strings.Split(strings.Trim(strings.TrimSpace(g[1]), ")"), ",") {
			value := strings.TrimSpace(allergen)
			foodInstance.allergens[value] = struct{}{}
			d.allergens[value] = struct{}{}
		}
		d.foods = append(d.foods, foodInstance)
	}
}

func (d *Day21) executeA() int {
	changed := true
	for changed {
		changed = false
		for allergen, _ := range d.allergens {
			found := make(map[string]int)
			count := 0
			for _, food := range d.foods {
				if _, ok := food.allergens[allergen]; !ok {
					continue
				}
				count++
				for ingredient, _ := range food.ingredients {
					if _, ok := d.ident[ingredient]; ok {
						continue
					}
					found[ingredient]++
				}
			}
			countFound := 0
			var ing string
			for ig, igCount := range found {
				if igCount == count {
					ing = ig
					countFound++
				}
			}
			if countFound == 1 {
				d.ident[ing] = allergen
				changed = true
			}
		}
	}

	count := 0
	for _, food := range d.foods {
		for ig, _ := range food.ingredients {
			if _, ok := d.ident[ig]; ok {
				continue
			}
			count++
		}
	}

	return count
}

func (d *Day21) executeB() string {

	refMap := make(map[string]string)
	list := make([]string, 0)
	for id, al := range d.ident {
		refMap[al] = id
		list = append(list, al)
	}
	resList := make([]string, len(list))
	for x, al := range list {
		resList[x] = refMap[al]
	}
	return strings.Join(resList, ",")
}

func (d *Day21) Handle(s string) ([]string, error) {
	d.init(s)
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%s", d.executeB()))
	return results, nil
}
