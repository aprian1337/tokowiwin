package formatter

import (
	"fmt"
	"regexp"
	"strings"
)

func Int64ToRupiah(amount int64) string {
	s := fmt.Sprintf("%d", amount)
	r := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != s; {
		n = s
		s = r.ReplaceAllString(s, "$1.$2")
	}
	if amount < 0 {
		s = strings.Replace(s, "-", "", -1)
		return fmt.Sprintf("-Rp%s,00", s)
	}

	return fmt.Sprintf("Rp%s,00", s)
}
