package question_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuestion2(t *testing.T) {
	assert.Equal(t, "210122", AnswerQuestion2("LLRR="))
	assert.Equal(t, "000210", AnswerQuestion2("==RLL"))
	assert.Equal(t, "221012", AnswerQuestion2("=LLRR"))
	assert.Equal(t, "012001", AnswerQuestion2("RRL=R"))
}
