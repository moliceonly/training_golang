package controlpractice

import (
	"testing"
)

type controlpracticeQuestionInput struct {
	Q11 int
	Q12 int
	Q13 int
	Q14_1 int
	Q14_2 int
	Q14_3 int
	Q17 int
	Q20 string
}

func Test_all_question_controlpratice(t *testing.T) {

	Input := controlpracticeQuestionInput{
		Q11: 10,
		Q12: 12,
		Q13: 2023,
		Q14_1: 12,
		Q14_2: 24,
		Q14_3: 13,
		Q17: 13,
		Q20: "哇，云朵，哒哒哒哒哒",
	}

	Question11(Input.Q11)
	Question12(Input.Q12)
	Question13(Input.Q13)
	Question14(Input.Q14_1, Input.Q14_2, Input.Q14_3)
	Question15()
	Question16()
	Question17(Input.Q17)
	Question18()
	Question19()
	Question20(Input.Q20)
}