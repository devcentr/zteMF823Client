package zteMF823Client

import "testing"

func TestClient_convertUTF2Unicode(t *testing.T) {
	client := Client{}
	cases := map[string]struct {
		value    string
		expected string
	}{
		"small word": {"word", "0077006F00720064"},
		"uppercase":  {"UPPERCASE", "005500500050004500520043004100530045"},
		"phrase":     {"One phrase", "004F006E00650020007000680072006100730065"},
		"cyrillic":   {"Привет", "041F04400438043204350442"},
		"symbols":    {":-)", "003A002D0029"},
	}

	for i, testCase := range cases {
		result := client.convertUTF2Unicode(testCase.value)
		if testCase.expected != result {
			t.Errorf("unexpected in %s case: value - %s, expected - %s, result - ", i, testCase.value, result)
		}
	}
}
