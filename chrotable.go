package chrotable

import (
	"fmt"
	"github.com/elliotchance/orderedmap/v3"
)

type Option[T any] func(*config[T])

type config[T any] struct {
	convertor  func(item T) []Cell
	calculator func(item T) (T, bool)
}

func WithConvertor[T any](convertor func(item T) []Cell) Option[T] {
	return func(c *config[T]) {
		c.convertor = convertor
	}
}

func WithConsumer[T any](calculator func(item T) (T, bool)) Option[T] {
	return func(c *config[T]) {
		c.calculator = calculator
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

func (c *Chrotable[T]) LoadConstants(constants map[string]any) {
	c.init()
	for k, v := range constants {
		c.SetConstant(k, v)
	}
}

func (c *Chrotable[T]) LoadVariables(variables map[string]any) {
	c.init()
	for k, v := range variables {
		c.SetVariable(k, v)
	}
}

func (c *Chrotable[T]) SetColumns(columns []Column) {
	c.columns = columns
}

func (c *Chrotable[T]) SetConstant(name string, value any) {
	c.init()
	if !c.constants.Has(name) {
		c.constants.Set(name, value)
	} else {
		fmt.Printf("[chrotable] warning: constant %s set duplicatedly\n", name)
	}
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

func (c *Chrotable[T]) Push(state T) (err error) {

	if c.config.convertor == nil {
		err = fmt.Errorf("convertor is not initialized")
		return
	}

	row := c.config.convertor(state)
	fmt.Println(row)

	return
}

func (c *Chrotable[T]) Calc(state T) (err error) {
	if c.config.calculator == nil {
		err = fmt.Errorf("calculator is not initialized")
		return
	}

	n, ok := c.config.calculator(state)
	if !ok {
		return
	}

	err = c.Push(n)
	if err != nil {
		return
	}

	return
}
