package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func (r ingredient) score() int {
	if r.capacity <= 0 || r.durability <= 0 || r.flavor <= 0 || r.texture <= 0 {
		return 0
	}
	return r.capacity * r.durability * r.flavor * r.texture
}

func (r ingredient) add(t ingredient) ingredient {
	return ingredient{
		r.name,
		r.capacity + t.capacity,
		r.durability + t.durability,
		r.flavor + t.flavor,
		r.texture + t.texture,
		r.calories + t.calories,
	}
}

func main() {
	lines := util.ReadFile("input.txt")
	ingredients := parse(lines)
	acc := ingredient{"", 0, 0, 0, 0, 0}
	var max *int = new(int)
	*max = 0
	var calorieMax *int = new(int)
	*calorieMax = 0
	dfs(0, 100, acc, ingredients, max, calorieMax)
	fmt.Println("max =", *max)
	fmt.Println("500 calorie max =", *calorieMax)
}

func parse(lines []string) []ingredient {
	ingredients := []ingredient{}
	for _, line := range lines {
		var name string
		var capacity int
		var durability int
		var flavor int
		var texture int
		var calories int
		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)
		ingredients = append(ingredients, ingredient{name[:len(name)-1], capacity, durability, flavor, texture, calories})
	}

	return ingredients
}

func dfs(start int, turn int, acc ingredient, ingredients []ingredient, max, calorieMax *int) {
	if turn == 0 {
		score := acc.score()
		if score > *max {
			*max = score
		}
		if score > *calorieMax && acc.calories == 500 {
			*calorieMax = score
		}
		return
	}

	for i := start; i < len(ingredients); i++ {
		r := ingredients[i]
		dfs(i, turn-1, acc.add(r), ingredients, max, calorieMax)
	}
}
