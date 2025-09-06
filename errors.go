package krlistsgo

import "fmt"

var (
	ErrListLastPage            = fmt.Errorf("this page is the last")
	ErrListFirstPage           = fmt.Errorf("this page is the first")
	ErrListIsNotPositiveNumber = fmt.Errorf("a page argument value should bigger than 0")
	ErrServerIdentifyIsNil     = fmt.Errorf("you should set your server's KrLists identify to use this endpoint")
	ErrBotIdentifyIsNil        = fmt.Errorf("you should set your bot's KrLists identify to use this endpoint")
	ErrAnyTokenIsNil           = fmt.Errorf("you should set any KrLists identify to use this endpoint")
)
