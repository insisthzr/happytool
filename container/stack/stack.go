package stack

type Stack struct {
	data []interface{}
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Top() (interface{}, bool) {
	if s.Len() == 0 {
		return nil, false
	}
	return s.data[s.Len()-1], true
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (interface{}, bool) {
	len := s.Len()
	if len == 0 {
		return nil, false
	}
	v := s.data[len-1]
	s.data = s.data[:len-1]
	return v, true
}

func (s *Stack) For(fn func(value interface{}) bool) {
	for i := 0; i < s.Len(); i++ {
		conti := fn(s.data[i])
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
		data: make([]interface{}, 0, config.Capacity),
	}
}

type StringStack struct {
	stack *Stack
}

func (s *StringStack) Len() int {
	return s.stack.Len()
}

func (s *StringStack) Top() (string, bool) {
	if s.Len() == 0 {
		return "", false
	}
	return s.stack.data[s.Len()-1].(string), true
}

func (s *StringStack) Push(value string) {
	s.stack.Push(value)
}

func (s *StringStack) Pop() (string, bool) {
	value, ok := s.stack.Pop()
	return value.(string), ok
}

func (s *StringStack) For(fn func(value string) bool) {
	for i := 0; i < s.Len(); i++ {
		conti := fn(s.stack.data[i].(string))
		if !conti {
			break
		}
	}
}

func NewStringStack() *StringStack {
	return &StringStack{stack: NewStack()}
}

func NewStringStackWithConfig(config *Config) *StringStack {
	return &StringStack{stack: NewStackWithConfig(config)}
}
