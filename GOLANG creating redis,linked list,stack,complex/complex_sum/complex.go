package complex_sum

import (
	"errors"
	"test/sum"
)

func Int(i int) *int {
	return &i
}

var ErrNilNumber = errors.New("null numbers are not allowed")

type ComplexSumItem struct {
	Number  *int
	Sum     sum.Sum
	Complex *ComplexSumItem
}

type ComplexSumItems []ComplexSumItem

func (c ComplexSumItems) Compute() (int, error) {
	r := 0

	for _, value := range c {
		if value.Number == nil && value.Sum == nil &&value.Complex == nil {

			return 0, ErrNilNumber
		}
		if value.Number != nil {

			r = r + *value.Number
		}

		if value.Sum != nil {
			sum, err := value.Sum.Compute()
			if err != nil {
				return 0, err
			}

			r = r + sum
		}

		if value.Complex!=nil{
			com,err:=value.Complex.Sum.Compute()
			if err!=nil{
				return 0,err
			}
			r= r + com

		}
	}
	return r, nil
}
