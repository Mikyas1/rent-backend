package conv

import "strconv"

func ToFloat32(receiver *float32, val interface{}) {
	if val == nil {
		return
	}
	f, err := strconv.ParseFloat(val.([]string)[0], 8)
	if err == nil {
		if receiver != nil {
			*receiver = float32(f)
		} else {
			l := float32(f)
			receiver = &l
		}
	}
	return
}

func ToFloat64(receiver *float64, val interface{}) {
	if val == nil {
		return
	}
	f, err := strconv.ParseFloat(val.([]string)[0], 8)
	if err == nil {
		if receiver != nil {
			*receiver = f
		} else {
			receiver = &f
		}
	}
	return
}

func ToInt(receiver *int, val interface{}) {
	if val == nil {
		return
	}
	intVar, err := strconv.Atoi(val.([]string)[0])
	if err == nil {
		if receiver != nil {
			*receiver = intVar
		} else {
			receiver = &intVar
		}
	}
	return
}

func ToBool(receiver *bool, val interface{}) {
	if val == nil {
		return
	}
	in := val.([]string)[0]
	if in == "false" {
		if receiver != nil {
			*receiver = false
		} else {
			l := false
			receiver = &l
		}
	} else if in == "true" {
		if receiver != nil {
			*receiver = true
		} else {
			l := true
			receiver = &l
		}
	}
	return
}

func GetFloat32(val interface{}) *float32 {
	if val == nil {
		return nil
	}
	f, err := strconv.ParseFloat(val.([]string)[0], 8)
	if err != nil {
		return nil
	}
	l := float32(f)
	return &l
}

func GetFloat64(val interface{}) *float64 {
	if val == nil {
		return nil
	}
	f, err := strconv.ParseFloat(val.([]string)[0], 8)
	if err != nil {
		return nil
	}
	return &f
}

func GetInt(val interface{}) *int {
	if val == nil {
		return nil
	}
	intVar, err := strconv.Atoi(val.([]string)[0])
	if err != nil {
		return nil
	}
	return &intVar
}

func GetBool(val interface{}) *bool {
	if val == nil {
		return nil
	}
	if val == "false" {
		l := false
		return &l
	}
	if val == "true" {
		l := true
		return &l
	}
	return nil
}