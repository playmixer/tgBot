package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TGetMe struct {
	Ok     bool `json:"ok"`
	Result struct {
		Id                      int64  `json:"id"`
		IsBot                   bool   `json:"is_bot"`
		FirstName               string `json:"first_name"`
		Username                string `json:"username"`
		CanJoinGroups           bool   `json:"can_join_groups"`
		CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
		SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	} `json:"result"`
}

type TGetUpdate struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateId int64 `json:"update_id"`
		Message  struct {
			MessageId int `json:"message_id"`
			From      struct {
				Id           int    `json:"id"`
				IsBot        bool   `json:"is_bot"`
				FirstName    string `json:"first_name"`
				LastName     string `json:"last_name"`
				Username     string `json:"username"`
				LanguageCode string `json:"language_code"`
			} `json:"from"`
			Chat struct {
				Id        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date  int    `json:"date"`
			Text  string `json:"text"`
			Voice struct {
				Duration     int    `json:"duration"`
				MimeType     string `json:"mime_type"`
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				FileSize     int    `json:"file_size"`
			} `json:"voice"`
		} `json:"message"`
	} `json:"result"`
}

type TelegramBot struct {
	Token   string
	api     string
	apiFile string
}

func New(token string) TelegramBot {
	tg := TelegramBot{
		api:     fmt.Sprintf("https://api.telegram.org/bot%s/", token),
		apiFile: fmt.Sprintf("https://api.telegram.org/file/bot%s/", token),
	}

	return tg
}

func post(url string, payload io.Reader) ([]byte, error) {
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body, nil
}

func (b *TelegramBot) methodURL(name string) string {
	return fmt.Sprintf("%s%s", b.api, name)
}

func (b *TelegramBot) GetMe() (TGetMe, error) {
	url := b.methodURL("getMe")

	body, _ := post(url, nil)

	var v TGetMe
	json.Unmarshal(body, &v)

	return v, nil
}

func (b *TelegramBot) GetUpdate(offset int64) (TGetUpdate, error) {
	url := b.methodURL("getMe")
	payload := strings.NewReader(fmt.Sprintf("{\"offset\":%v,\"limit\":null,\"timeout\":null}", offset))
	body, _ := post(url, payload)
	var v TGetUpdate
	json.Unmarshal(body, &v)

	return v, nil
}
