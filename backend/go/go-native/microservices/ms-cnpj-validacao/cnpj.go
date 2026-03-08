package main

func normalizeCNPJ(input *string) *string {
	if input == nil {
		return nil
	}

	digits := make([]byte, 0, len(*input))
	for i := 0; i < len(*input); i++ {
		c := (*input)[i]
		if c >= '0' && c <= '9' {
			digits = append(digits, c)
		}
	}

	if len(digits) == 0 {
		return nil
	}

	n := string(digits)
	return &n
}

func isValidCNPJ(input *string) bool {
	normalized := normalizeCNPJ(input)
	if normalized == nil {
		return false
	}

	digits := *normalized
	if len(digits) != 14 {
		return false
	}

	if allDigitsSame(digits) {
		return false
	}

	d1 := calculateCheckDigit(digits, 12)
	d2 := calculateCheckDigit(digits, 13)

	return digits[12] == byte('0'+d1) && digits[13] == byte('0'+d2)
}

func allDigitsSame(digits string) bool {
	first := digits[0]
	for i := 1; i < len(digits); i++ {
		if digits[i] != first {
			return false
		}
	}
	return true
}

func calculateCheckDigit(digits string, length int) byte {
	var weights []int
	if length == 12 {
		weights = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	} else {
		weights = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	}

	sum := 0
	for i := 0; i < length; i++ {
		d := int(digits[i] - '0')
		sum += d * weights[i]
	}

	mod := sum % 11
	if mod < 2 {
		return 0
	}
	return byte(11 - mod)
}
