package tempconv

import "fmt"
import "flag"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// *celsiusFlagはflag.Valueインタフェースを満足します。
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) //エラー検証は必要ない
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FtoC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KtoC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("Invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
