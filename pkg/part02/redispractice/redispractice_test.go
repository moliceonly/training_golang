package redispractice

import (
	"os"
	"testing"
)

func Test_all_question_redispractice(t *testing.T) {
	if os.Getenv("TRAINING_SKIP_REDIS") == "1" {
		t.Skip("TRAINING_SKIP_REDIS=1")
	}
	// 未装 Redis / 连不上时，各 Question 内部应自行容错打印；此处仍调用以练题干。
	Question122()
	Question123()
	Question124()
	Question125()
	Question126()
	Question127()
	Question128()
}
