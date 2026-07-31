package shortlinkpractice

import (
	"os"
	"testing"
)

func Test_all_question_shortlinkpractice(t *testing.T) {
	if os.Getenv("TRAINING_SKIP_MYSQL") == "1" || os.Getenv("TRAINING_SKIP_REDIS") == "1" {
		t.Skip("TRAINING_SKIP_MYSQL=1 or TRAINING_SKIP_REDIS=1")
	}
	// 未装 MySQL/Redis / 连不上时，各 Question 内部应自行容错打印；此处仍调用以练题干。
	Question129()
	Question130()
	Question131()
	Question132()
	Question133()
	Question134()
}
