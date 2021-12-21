package sum

import "errors"

type Sum []int

var ErrNegativeNumber = errors.New("negative numbers are not allowed")

func (s Sum) Compute() (int, error) {
	r := 0
	for _, value := range s {

		r = r + value
		if value < 0 {
			return 0, ErrNegativeNumber
		}

	}
	return r, nil

}
