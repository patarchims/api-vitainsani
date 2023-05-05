package entity

import "vincentcoreapi/modules/telegram"

// UserUseCase
type TelegramUseCase interface {
	SendMessage(url string, message telegram.Message) error
}
