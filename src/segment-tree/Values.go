package segment_tree

type Value interface {
	// Returns sum of values
	// Attention: current value would not be changed
	Add(value Value) Value
}

type IntValue int
type FloatValue float64
type StringValue string

func (cur IntValue) Add(value Value) Value {
	if value == nil {
		return cur
	} else {
		return cur + value.(IntValue)
	}
}

func (cur FloatValue) Add(value Value) Value {
	if value == nil {
		return cur
	} else {
		return cur + value.(FloatValue)
	}
}

func (cur StringValue) Add(value Value) Value {
	if value == nil {
		return cur
	} else {
		return cur + value.(StringValue)
	}
}
