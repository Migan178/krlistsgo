package krlistsgo

import "fmt"

// WidgetStyle은 한디리 위젯 스타일의 타입입니다.
type WidgetStyle string

// BotWidgetType은 봇의 위젯이 담을 정보의 타입입니다.
type BotWidgetType string

// ServerWidgetType은 서버의 위젯이 담을 정보의 타입입니다.
type ServerWidgetType string

// 위젯의 스타일입니다.
const (
	WidgetClassic WidgetStyle = "classic"
	WidgetFlat    WidgetStyle = "flat"
)

// 봇 위젯의 타입입니다.
const (
	WidgetBotVotes   BotWidgetType = "votes"
	WidgetBotServers BotWidgetType = "server"
	WidgetBotStatus  BotWidgetType = "status"
)

// 서버 위젯의 타입입니다.
const (
	WidgetServerVotes     ServerWidgetType = "votes"
	WidgetServerMembers   ServerWidgetType = "members"
	WidgetServerBoostTier ServerWidgetType = "boost"
)

func getWidget(isBot bool, id, widgetType, style string, scale float32, icon bool) (string, error) {
	botOrServer := "bots"
	if scale < 0.5 || scale > 3.0 {
		return "", fmt.Errorf("scale의 값은 0.5이상, 3.0이하입니다.")
	}

	if !isBot {
		botOrServer = "servers"
	}
	return fmt.Sprintf(
		"%s/widget/%s/%s/%s.svg?style=%s&scale=%f&icon=%t",
		API_URL,
		botOrServer,
		widgetType,
		id,
		style,
		scale,
		icon,
	), nil
}

// BotWidget은 봇의 위젯 URL을 반환합니다.
func (k *KrLists) BotWidget(id string, widgetType BotWidgetType, style WidgetStyle, scale float32, icon bool) (string, error) {
	return getWidget(true, id, string(widgetType), string(style), scale, icon)
}

// ServerWidget은 서버의 위젯 URL을 반환합니다.
func (k *KrLists) ServerWidget(id string, widgetType ServerState, style WidgetStyle, scale float32, icon bool) (string, error) {
	return getWidget(false, id, string(widgetType), string(style), scale, icon)
}

// Widget은 해당 봇의 위젯의 URL을 반환합니다.
func (b *Bot[T]) Widget(widgetType BotWidgetType, style WidgetStyle, scale float32, icon bool) (string, error) {
	return getWidget(true, b.ID, string(widgetType), string(style), scale, icon)
}

// Widget은 해당 서버의 위젯의 URL을 반환합니다.
func (s *Server[T]) Widget(widgetType ServerState, style WidgetStyle, scale float32, icon bool) (string, error) {
	return getWidget(false, s.ID, string(widgetType), string(style), scale, icon)
}
