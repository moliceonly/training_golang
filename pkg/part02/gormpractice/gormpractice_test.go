package gormpractice

import (
	"os"
	"testing"
)

func Test_all_question_gormpractice(t *testing.T) {
	if os.Getenv("TRAINING_MYSQL_DSN") == "" && os.Getenv("TRAINING_SKIP_MYSQL") == "1" {
		t.Skip("TRAINING_SKIP_MYSQL=1")
	}
	// 未装 MySQL / 连不上时，各 Question 内部应自行容错打印；此处仍调用以练题干。
	Question114()
	Question115()
	Question116()
	Question117()
	Question118()
	Question119()
	Question120()
	Question121()
}
