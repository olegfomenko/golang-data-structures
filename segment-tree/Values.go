package segment_tree

type Value interface {
	// Get current value
	Get() Value

	// Returns sum of values
	// Attention: current value would not be changed
	Add(value Value) Value
}

type IntValue int
type FloatValue float64
type StringValue string

func (cur IntValue) Get() Value    { return cur }
func (cur FloatValue) Get() Value  { return cur }
func (cur StringValue) Get() Value { return cur }

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

func GetInt(val int) Value {
	return IntValue(val)
}

func GetFloat(val float64) Value {
	return FloatValue(val)
}

func GetString(val string) Value {
	return StringValue(val)
}
