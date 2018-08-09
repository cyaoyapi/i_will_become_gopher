package tempconv

func CToF(c Celsuis) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsuis {
	return Celsuis((f - 32) * 5 / 9)
}
