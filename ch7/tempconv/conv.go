package tempconv

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func FtoK(f Fahrenheit) Kelvin { return CtoK(FtoC(f)) }
func KtoF(k Kelvin) Fahrenheit { return CtoF(KtoC(k)) }
