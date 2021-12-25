package healthsug

import "testing"

func TestHealthSug(t *testing.T) {
	sug := HeslthSug("男", 30, 0.147)
	if sug != "标准，继续保持！" {
		t.Fatalf("预期是：标准，继续保持！ 但测试为：%s。", sug)
	}

}
