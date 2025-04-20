package krlistsgo

import "fmt"

var (
	ListLastPageErr       = fmt.Errorf("이 페이지가 마지막입니다.")
	ListFirstPageErr      = fmt.Errorf("이 페이지가 가장 처음입니다.")
	ListPositiveNumberErr = fmt.Errorf("page 인수의 값은 0보다 커야합니다.")
	ServerIdentifyIsNil   = fmt.Errorf("해당 기능을 사용할려면 KrLists.SetServerIdentify로 자격 증명을 추가해야합니다.")
	BotIdentifyIsNil      = fmt.Errorf("해당 기능을 사용할려면 KrLists.SetBotIdentify로 자격 증명을 추가해야합니다.")
)
