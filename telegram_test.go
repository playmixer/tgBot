package tgbot

import (
	"fmt"
	"testing"
)

const (
	TOKEN = "5002576807:AAFyj3XEBRISbse7Pp4HUwpnuV5me6ayRjQ"
)

func TestGetMe(t *testing.T) {
	tg := New(TOKEN)
	res, err := tg.GetMe()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
