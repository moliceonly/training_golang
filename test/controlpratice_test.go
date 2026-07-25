package test

import (
	"training_golang/basepractice"
	"testing"
)

type QuestionInput struct {
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

	Input := QuestionInput{
		Q11: 10,
		Q12: 12,
		Q13: 2023,
		Q14_1: 12,
		Q14_2: 24,
		Q14_3: 13,
		Q17: 13,
		Q20: "哇，云朵，哒哒哒哒哒",
	}

	basepractice.Question11(Input.Q11)
	basepractice.Question12(Input.Q12)
	basepractice.Question13(Input.Q13)
	basepractice.Question14(Input.Q14_1, Input.Q14_2, Input.Q14_3)
	basepractice.Question15()
	basepractice.Question16()
	basepractice.Question17(Input.Q17)
	basepractice.Question18()
	basepractice.Question19()
	basepractice.Question20(Input.Q20)
}