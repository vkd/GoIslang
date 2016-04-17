package main

import "story"

func main() {
	day := story.Morning()
	_ = day

	food := day.GoToForest()
	day.Eat(food)

}
