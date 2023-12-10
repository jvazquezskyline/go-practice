package utils



// Function is responsible for returning a given value if 
// the value passed is empty, nil, or zero (0)
func ReturnSafeValue(value ,safeValue interface{}) interface{} {
	if value == "" || value == 0 || value == nil {
		return safeValue
	}

	return value
}


