package complex_sum

import (
	"test/sum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0Items(t *testing.T) {

	s := ComplexSumItems{}
	actual, err := s.Compute()
	expected := 0
	assert.Equal(t, expected, actual)
	assert.Nil(t, err)

}
func Test2Numbers(t *testing.T) {
	s := ComplexSumItems{
		{
			Number: Int(5),
		},
		{
			Number: Int(5),
		},
	}
	actual, err := s.Compute()
	expected := 10
	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}

func Test2NumbersOneNil(t *testing.T) {
	s := ComplexSumItems{
		{
			Number: Int(5),
		},
		{
			Number: nil,
		},
	}
	_, err := s.Compute()
	expected := ErrNilNumber

	assert.Equal(t, expected, err)
}

func Test1NumberOneSum(t *testing.T) {
	s := ComplexSumItems{
		{
			Number: Int(5),
		},
		{
			Sum: sum.Sum{4, 6},
		},
	}
	actual, err := s.Compute()
	expected := 15
	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}

func Test1NumberOneSumOneComplex(t *testing.T) {
	s := ComplexSumItems{
		{
			Number: Int(5),
		},
		{
			Sum: sum.Sum{4, 6},
		},
		{
			Complex: &ComplexSumItem{
				Sum: sum.Sum{4, 6},
			},
		},
	}
	actual, err := s.Compute()
	expected := 25
	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}
