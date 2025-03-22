package chrotable

import "github.com/elliotchance/orderedmap/v3"

type Option[T any] func(*config[T])

type config[T any] struct {
	convertor func(item T) []Cell
}

func WithConvertor[T any](convertor func(item T) []Cell) Option[T] {
	return func(c *config[T]) {
		c.convertor = convertor
	}
}

type Chrotable[T any] struct {
	config    *config[T]
	columns   []Column
	constants *orderedmap.OrderedMap[string, any]
	variables *orderedmap.OrderedMap[string, any]
	rows      [][]Cell
}

func NewChrotable[T any](options ...Option[T]) *Chrotable[T] {
	c := &Chrotable[T]{
		config: &config[T]{},
	}
	for _, option := range options {
		option(c.config)
	}
	return c
}

func (c *Chrotable[T]) init() {
	if c.config == nil {
		c.config = &config[T]{}
	}
	if c.constants == nil {
		c.constants = orderedmap.NewOrderedMap[string, any]()
	}
	if c.variables == nil {
		c.variables = orderedmap.NewOrderedMap[string, any]()
	}
}

func (c *Chrotable[T]) SetColumns(columns []Column) {
	c.columns = columns
}

func (c *Chrotable[T]) SetConstant(name string, value any) {
	c.init()
	c.constants.Set(name, value)
}

func (c *Chrotable[T]) GetConstant(name string) any {
	if c.constants == nil {
		return nil
	}
	v, _ := c.constants.Get(name)
	return v
}

func (c *Chrotable[T]) SetVariable(name string, value any) {
	c.init()
	c.variables.Set(name, value)
}

func (c *Chrotable[T]) GetVariable(name string) any {
	if c.variables == nil {
		return nil
	}
	v, _ := c.variables.Get(name)
	return v
}
