package story

type god struct {
	item_storage map[godKeyType]bool
}

func (g *god) register(w watcher) {
	g.item_storage = append(g.item_storage, w.getGodKey())
}

type godKeyType string

type godWatcher struct {
	godKey godKeyType
}

func (g godWatcher) getGodKey() godKeyType {
	return g.godKey
}

type watcher interface {
	getGodKey() godKeyType
}

func (g *god) GodWatch(obj watcher) bool {
	if obj.getGodKey() != 221133 {
		panic("Великий создатель покарал вас за читерство")
	}
	return true
}
