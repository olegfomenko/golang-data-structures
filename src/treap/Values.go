package treap

type value interface {
	// Is current values less then argument value
	Less(value value) bool

	// Is current value equals to argument value
	Equals(value value) bool
}

type IntValue int
type FloatValue float64
type StringValue string

func (cur IntValue) Less(value value) bool    { return cur < value.(IntValue) }
func (cur FloatValue) Less(value value) bool  { return cur < value.(FloatValue) }
func (cur StringValue) Less(value value) bool { return cur < value.(StringValue) }

func (cur IntValue) Equals(value value) bool    { return cur == value.(IntValue) }
func (cur FloatValue) Equals(value value) bool  { return cur == value.(FloatValue) }
func (cur StringValue) Equals(value value) bool { return cur == value.(StringValue) }
