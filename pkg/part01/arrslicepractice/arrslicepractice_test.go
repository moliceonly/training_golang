package arrslicepractice

import (
	"testing"
)

type arrslicepracticeQuestionInput struct {
	Q26 []int
	Q29 []int
	Q30 []int
}

func Test_all_question_arrslicepractice(t *testing.T) {

	Input := arrslicepracticeQuestionInput{
		Q26: []int{23, 82, 12, 42, 213, 2, 92, 72},
		Q29: []int{23, 82, 12, 42, 213, 2, 92, 72},
		Q30: []int{23, 82, 12, 42, 213, 2, 92, 72},
	}

	Question21()
	Question22()
	Question23()
	Question24()
	Question25()
	Question26(Input.Q26)
	Question27()
	Question28()
	Question29(Input.Q29)
	Question30(Input.Q30)
}
