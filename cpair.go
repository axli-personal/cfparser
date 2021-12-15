package cfparser

type CPair struct {
	left   *CPair
	right  *CPair
	key    string
	val    string
	signal chan *CPair
}

// Convert pair's value to string.
func (pair *CPair) String() string {
	return pair.val
}

// Convert pair's value to bool. Return true only when value is "true", "True" or "TRUE".
func (pair *CPair) Bool() bool {
	switch pair.String() {
	case "true":
		return true
	case "True":
		return true
	case "TRUE":
		return true
	}
	return false
}

// When the pair change in the future, you will receive the modified pair immediately.
// The channel will create when you call this, and it's buffered channel with size of 1.
func (pair *CPair) Watch() chan *CPair {
	if pair.signal == nil {
		pair.signal = make(chan *CPair, 1)
	}
	return pair.signal
}

// Constructs pair and store it.
func put(pair *CPair, key string, val string) *CPair {
	if pair == nil {
		return &CPair{key: key, val: val}
	}

	if key < pair.key {
		pair.left = put(pair.left, key, val)
	} else if key > pair.key {
		pair.right = put(pair.right, key, val)
	} else {
		pair.val = val
		if pair.signal != nil {
			pair.signal <- pair
		}
	}

	return pair
}

// Find pair and get it.
func get(pair *CPair, key string) *CPair {
	if pair == nil {
		return nil
	}

	if key < pair.key {
		return get(pair.left, key)
	} else if key > pair.key {
		return get(pair.right, key)
	} else {
		return pair
	}
}
