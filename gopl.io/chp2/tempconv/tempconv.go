// Package tempconv performs Celsuis and Fahrenheit conversions.

package tempconv

import "fmt"

type Celsuis float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsuis = -273.15
	FreezingC     Celsuis = 0
	BoilingC      Celsuis = 100
)

func (c Celsuis) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
