package test

import (
	"training_golang/basepractice"
	"testing"
)

type functionmethodpracticeQuestionInput struct {
	Q31 [2]int
	Q32 [2]int
	Q33 [2]int
	Q36 int
	Q37 int
	Q39 [2]float64
	Q40 [3]float64
}

func Test_all_question_functionmethodpractice(t *testing.T) {

	Input := functionmethodpracticeQuestionInput {
		Q31: [2]int{2, 20},
		Q32: [2]int{122,23},
		Q33: [2]int{32,42},
		Q36: 5,
		Q37: 10,
		Q39: [2]float64{2.12, 4.23},
		Q40: [3]float64{1.32, 3.14, 8.7},
	}

	basepractice.Question31(Input.Q31[0], Input.Q31[1])
	basepractice.Question32(Input.Q32[0], Input.Q32[1])
	basepractice.Question33(Input.Q33[0], Input.Q33[1])
	basepractice.Question34(1, 4, 4231, 23, 234)
	basepractice.Question35()
	basepractice.Question36(Input.Q36)
	basepractice.Question37(Input.Q37)
	basepractice.Question38()
	basepractice.Question39(Input.Q39[0], Input.Q39[1])
	basepractice.Question40(Input.Q40[0], Input.Q40[1],Input.Q40[2])
}