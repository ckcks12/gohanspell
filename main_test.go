package gohanspell

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPostPusanUniv(t *testing.T) {
	str := `
나는 그렇지 않다고 생각합니다.
왜 않되?

앙기모띠

라이센스
라이선스
`
	txt, err := PostPusanUniv(str)
	if err != nil {
		t.Error(err)
	}

	licenseAlt := strings.Split("사용권|라이선스|면허증|면허|허가장|면허장", "|")
	answer1 := `나는 그렇지 않다고 생각합니다.
왜 안 돼?

기분 좋아

`
	answer2 := `
라이선스`

	txt = strings.ReplaceAll(txt, answer1, "")
	txt = strings.ReplaceAll(txt, answer2, "")
	assert.Contains(t, licenseAlt, txt)
}
