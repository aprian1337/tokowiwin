package formatter

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"tokowiwin/constants"
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
		return fmt.Sprintf("-Rp%s", s)
	}

	return fmt.Sprintf("Rp%s", s)
}

func ToTimezoneJakarta(t time.Time) time.Time {
	location, _ := time.LoadLocation(constants.TimezoneAsiaJakarta)
	return t.In(location)
}
