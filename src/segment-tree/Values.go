package segment_tree

type value interface {
	// Returns sum of values
	// Attention: current value would not be changed
	Add(value value) value
}

type IntValue int
type FloatValue float64
type StringValue string

func (cur IntValue) Add(value value) value {
	if value == nil {
		return cur
	} else {
		return cur + value.(IntValue)
	}
}

func (cur FloatValue) Add(value value) value {
	if value == nil {
		return cur
	} else {
		return cur + value.(FloatValue)
	}
}

func (cur StringValue) Add(value value) value {
	if value == nil {
		return cur
	} else {
		return cur + value.(StringValue)
	}
}
