package entity

import "vincentcoreapi/modules/telegram"

type TelegramUseCase interface {
	SendMessage(url string, message telegram.Message) error
}
