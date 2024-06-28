package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/LombardiDaniel/ip-checker/controllers"
	"github.com/LombardiDaniel/ip-checker/services"
	"github.com/LombardiDaniel/ip-checker/utils"
)

var (
	ipCheckerService			services.IpCheckerService
	notifierService				services.NotifierService
)

func init() {
	// senderEmail := os.Getenv("EMAIL_HOST_USER")
    // senderEmailPass := os.Getenv("EMAIL_HOST_PASSWORD")
    telegramBotToken := os.Getenv("TELEGRAM_BOT_KEY")
    // mailList := strings.Split(os.Getenv("MAIL_LIST"), ",")
    telegramChatList := strings.Split(os.Getenv("TELEGRAM_CHAT_LIST"), ",")
	deviceName := utils.GetEnvVarDefault("DEVICE_NAME", "raspberry")

	ipCheckerService = services.NewIpCheckerImpl()
	notifierService = services.NewNotifierServiceTelegramImpl(
		telegramBotToken,
		telegramChatList,
		deviceName,
	)
}

func main() {
	controller := controllers.NewController(ipCheckerService, notifierService)

	err := controller.Loop()
	if err != nil {
		slog.Error(fmt.Sprintf("error in loop, exiting: %s", err.Error()))
	}
}