package chrotable

import "github.com/elliotchance/orderedmap/v3"

type Chrotable[T any] struct {
	columns   []Column
	constants *orderedmap.OrderedMap[string, any]
	variables *orderedmap.OrderedMap[string, any]
	rows      [][]Cell
	convertor func(item T) []Cell
}

func (c *Chrotable[T]) init() {
	if c.constants == nil {
		c.constants = orderedmap.NewOrderedMap[string, any]()
	}
	if c.variables == nil {
		c.variables = orderedmap.NewOrderedMap[string, any]()
	}
}

func (c *Chrotable[T]) SetConvertor(convertor func(item T) []Cell) {
	c.convertor = convertor
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
