package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notifaction-service/pkg/loadenv"
	"os"
)

func Send(phone string, otp string) error {
loadenv.Load()

	log.Println("[TELEGRAM MOCK SMS REPLACEMENT]")

	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	text := fmt.Sprintf("🔐 OTP CODE\nPhone: %s\nOTP: %s", phone, otp)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	payload := map[string]any{
		"chat_id": chatID,
		"text":    text,
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("OTP sent to Telegram group")
	return nil
}