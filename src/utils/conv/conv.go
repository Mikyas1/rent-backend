package conv

import "strconv"

func ToFloat32(receiver *float32, val interface{}) {
	if val == nil {
		return
	}
	f, err := strconv.ParseFloat(val.([]string)[0], 8)
	if err == nil {
		*receiver = float32(f)
	}
	return
}

func ToFloat64(receiver *float64, val interface{}) {
	if val == nil {
		return
	}
	f, err := strconv.ParseFloat(val.([]string)[0], 8)
	if err == nil {
		*receiver = f
	}
	return
}

func ToInt(receiver *int, val interface{}) {
	if val == nil {
		return
	}
	intVar, err := strconv.Atoi(val.([]string)[0])
	if err == nil {
		*receiver = intVar
	}
	return
}

func ToBool(receiver *bool, val interface{}) {
	if val == nil {
		return
	}
	in := val.([]string)[0]
	if in == "false" {
		*receiver = false
	} else if in == "true" {
		*receiver = true
	}
	return
}
