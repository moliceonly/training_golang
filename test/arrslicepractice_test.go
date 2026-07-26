package test

import (
	"training_golang/basepractice"
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

	basepractice.Question21()
	basepractice.Question22()
	basepractice.Question23()
	basepractice.Question24()
	basepractice.Question25()
	basepractice.Question26(Input.Q26)
	basepractice.Question27()
	basepractice.Question28()
	basepractice.Question29(Input.Q29)
	basepractice.Question30(Input.Q30)
}