package fileiopractice

import (
	"testing"
	"strings"
)

func Test_all_question_fileiopractice(t *testing.T) {
	tmp := t.TempDir()
	Question94(tmp)
	Question95(tmp)
	t.Run("Question95", func(t *testing.T) {
		TestSafeJoin(t)
	})
	Question96(tmp)
	Question97(tmp)
	Question98(tmp)
	Question99(tmp)
	Question100(tmp)
	Question101()
}

// 可选：单独测 SafeJoin，实现后取消注释补全断言
func TestSafeJoin(t *testing.T) {
	root := t.TempDir()
	got, err := SafeJoin(root, "a/b.txt")
	if err != nil {
		t.Logf("合法路径不应报错: %v", err)
	}
	if !strings.HasPrefix(got, root) {
		t.Logf("got=%q 应在 root=%q 下", got, root)
	}
	_, err = SafeJoin(root, "../../etc/passwd")
	if err == nil {
		t.Logf("穿越路径应返回 error")
	}
}
