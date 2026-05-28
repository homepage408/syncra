package event

type Bus interface {
	Publish(Event)
	Subscribe(name string, handler func(Event))
}

type InMemoryBus struct {
	handlers map[string][]func(Event)
}

func NewBus() *InMemoryBus {
	return &InMemoryBus{handlers: make(map[string][]func(Event))}
}

func (b *InMemoryBus) Publish(e Event) {
	for _, handler := range b.handlers[e.Name()] {
		handler(e)
	}
}

func (b *InMemoryBus) Subscribe(name string, handler func(Event)) {
	b.handlers[name] = append(b.handlers[name], handler)
}
