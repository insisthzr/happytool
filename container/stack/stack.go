package stack

type Stack struct {
	stack []interface{}
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) Push(v interface{}) {
	s.stack = append(s.stack, v)
}

func (s *Stack) Pop() (interface{}, bool) {
	len := s.Len()
	if len == 0 {
		return nil, false
	}
	v := s.stack[len-1]
	s.stack = s.stack[:len-1]
	return v, true
}

func (s *Stack) For(fn func(value interface{}) bool) {
	for i := 0; i < s.Len(); i++ {
		conti := fn(s.stack[i])
		if !conti {
			break
		}
	}
}

type Config struct {
	Capacity int
}

var (
	DefaultConfig = &Config{
		Capacity: 16,
	}
)

func NewStack() *Stack {
	return NewStackWithConfig(DefaultConfig)
}

func NewStackWithConfig(config *Config) *Stack {
	if config == nil {
		return newStackWithConfig(DefaultConfig)
	}
	if config.Capacity == 0 {
		config.Capacity = DefaultConfig.Capacity
	}
	return newStackWithConfig(config)
}

func newStackWithConfig(config *Config) *Stack {
	return &Stack{
		stack: make([]interface{}, 0, config.Capacity),
	}
}

type StringStack struct {
	*Stack
}

func (s *StringStack) Len() int {
	return s.Stack.Len()
}

func (s *StringStack) Push(value string) {
	s.Stack.Push(value)
}

func (s *StringStack) Pop() (string, bool) {
	value, ok := s.Stack.Pop()
	return value.(string), ok
}

func (s *StringStack) For(fn func(value string) bool) {
	for i := 0; i < s.Len(); i++ {
		conti := fn(s.stack[i].(string))
		if !conti {
			break
		}
	}
}

func NewStringStack() *StringStack {
	return &StringStack{Stack: NewStack()}
}

func NewStringStackWithConfig(config *Config) *StringStack {
	return &StringStack{Stack: NewStackWithConfig(config)}
}
