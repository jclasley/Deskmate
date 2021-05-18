package slack

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_match(t *testing.T) {
	// ! `HandleMentionEvent` is not currently unit-testable d/t integration with `api.#` and `script.#` methods
	// ? Build out with interfaces to allow for mocking?
	
	// TODO: Make test integrate with `HandleMentionEvent`
	
	type test struct {
		content string
	}
	matchers := []string{"help", "set", "unset"}
	tt := []test{
		{"@Deskmate help"},
		{"@deskmate help"},
		{"@<U123456> help"},
		{"@Deskmate set"},
		{"@deskmate set"},
		{"@<U123456> set"},
		{"@Deskmate unset"},
		{"@deskmate unset"},
		{"@<U123456> unset"},
	}

	// re, err := regexp.Compile(`^<{0,1}@.*>{0,1} *`)
	re, err := regexp.Compile(`^@\S* *`)
	if err != nil {
		t.Error(err.Error())
	}

	for i := 0; i < len(matchers); i++ {
		for j := 0; j < 3; j++ { // loop through various mentions, only 3 variations
			c := tt[j + i * 3].content
			fmt.Println(re.FindString(c))
			c = re.ReplaceAllString(c, "")
			if !match(matchers[i], c) {
				t.Errorf("%s did not match regex as expected", c)
			}
		}
	}

}