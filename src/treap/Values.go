package treap

type Value interface {
	// Is current values less then argument value
	Less(value Value) bool

	// Is current value equals to argument value
	Equals(value Value) bool
}

type IntValue int
type FloatValue float64
type StringValue string

func (cur IntValue) Less(value Value) bool    { return cur < value.(IntValue) }
func (cur FloatValue) Less(value Value) bool  { return cur < value.(FloatValue) }
func (cur StringValue) Less(value Value) bool { return cur < value.(StringValue) }

func (cur IntValue) Equals(value Value) bool    { return cur == value.(IntValue) }
func (cur FloatValue) Equals(value Value) bool  { return cur == value.(FloatValue) }
func (cur StringValue) Equals(value Value) bool { return cur == value.(StringValue) }
