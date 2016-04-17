package story

import "log"

var (
	g = &god{}
)

func Morning() *Day1 {
	log.Printf("Доброе утро. Вы проснулись на острове.")
	log.Printf("Ваша задача: съесть банан")
	return &Day1{}
}

type Food interface {
	watcher
	GetName() string
}

type Banana struct {
	godWatcher
}

func (b Banana) GetName() string {
	return "banana"
}

type Day1 struct {
	godWatcher
}

func (d *Day1) GoToForest() Food {
	log.Printf("Вы нашли еду")
	f := &Banana{}
	f.godKey = 221133
	g.register(f)
	return f
}

func (d *Day1) Eat(eat Food) {
	if !g.GodWatch(eat) {
		return
	}
	if _, ok := eat.(Banana); !ok {
		log.Printf("Вы скушали странный фрукт")
		return
	}
	log.Printf("Победа. Вы выжили на острове.")
}
