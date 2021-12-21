package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_case0elements(t *testing.T) {
	s := Sum{}
	actual,err:= s.Compute()
	expected := 0
	assert.Nil(t,err)
	assert.Equal(t, expected, actual)

}
func Test_case2elements(t *testing.T) {
	s := Sum{3, 4}
	actual,err := s.Compute()
	expected := 7
	assert.Nil(t,err)
	assert.Equal(t, expected, actual)

}

func Test_caseNotNegative(t *testing.T) {
	s := Sum{3, -4}
	_, err := s.Compute()
	expected := ErrNegativeNumber

	assert.Equal(t, expected, err)
}
