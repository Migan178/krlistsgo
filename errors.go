package krlistsgo

import "fmt"

var (
	ListLastPageErr       = fmt.Errorf("이 페이지가 마지막입니다.")
	ListFirstPageErr      = fmt.Errorf("이 페이지가 가장 처음입니다.")
	ListPositiveNumberErr = fmt.Errorf("page 인수의 값은 0보다 커야합니다.")
)
