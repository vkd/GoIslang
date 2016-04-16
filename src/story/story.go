package story

import "log"

func Morning() Dayer {
	log.Printf("Доброе утро. Вы проснулись на острове.")
	return &Day1{}
}

type Dayer interface {
	GoToForest() *Item
}

type Item struct {
	Name string
}

type Day1 struct {
}

func (d *Day1) GoToForest() *Item {
	log.Printf("Вы нашли еду")
	return &Item{"food"}
}
