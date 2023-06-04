package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elizabeth-dev/plush_bot/internal/types"
	"net/http"
)

type TelegramAdapter struct {
	c *http.Client
}

func NewTelegramAdapter(c *http.Client) *TelegramAdapter {
	return &TelegramAdapter{c: c}
}

func (a *TelegramAdapter) SendMessage(token string, req types.TelegramSendMessage) (*int, error) {
	payloadBuf := new(bytes.Buffer)
	if err := json.NewEncoder(payloadBuf).Encode(req); err != nil {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE ENCODING PAYLOAD: %v\n", err)
		fmt.Printf("REQUEST: %v\n", req)

		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, "https://api.telegram.org/bot"+token+"/sendMessage", payloadBuf)
	if err != nil {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE CREATING REQUEST: %v\n", err)
		fmt.Printf("TOKEN: %s\n", token)
		fmt.Printf("REQUEST: %v\n", req)

		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	res, err := a.c.Do(httpReq)

	if err != nil {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE SENDING REQUEST: %v\n", err)
		fmt.Printf("TOKEN: %s\n", token)
		fmt.Printf("REQUEST: %v\n", req)

		return nil, err
	}

	type response struct {
		Ok          bool                  `json:"ok"`
		Description string                `json:"description"`
		Result      types.TelegramMessage `json:"result"`
	}

	var resp response
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE DECODING RESPONSE: %v\n", err)
		fmt.Printf("BODY: %v\n", res.Body)

		return nil, nil
	}

	if (res.StatusCode != http.StatusOK) || !resp.Ok {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE SENDING MESSAGE: %s\n", resp.Description)
		fmt.Printf("TOKEN: %s\n", token)
		fmt.Printf("REQUEST: %v\n", req)

		return nil, fmt.Errorf("error while sending message: %s", resp.Description)
	}

	fmt.Printf("RESPONSE: %v\n", resp)

	return &resp.Result.Id, nil
}

func (a *TelegramAdapter) SendVoice(token string, req types.TelegramSendVoice) error {
	payloadBuf := new(bytes.Buffer)
	if err := json.NewEncoder(payloadBuf).Encode(req); err != nil {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE ENCODING PAYLOAD: %v\n", err)
		fmt.Printf("REQUEST: %v\n", req)

		return err
	}

	httpReq, err := http.NewRequest(http.MethodPost, "https://api.telegram.org/bot"+token+"/sendVoice", payloadBuf)
	if err != nil {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE CREATING REQUEST: %v\n", err)
		fmt.Printf("TOKEN: %s\n", token)
		fmt.Printf("REQUEST: %v\n", req)

		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	res, err := a.c.Do(httpReq)

	if err != nil {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE SENDING REQUEST: %v\n", err)
		fmt.Printf("TOKEN: %s\n", token)
		fmt.Printf("REQUEST: %v\n", req)

		return err
	}

	type response struct {
		Ok          bool   `json:"ok"`
		Description string `json:"description"`
	}

	var resp response
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE DECODING RESPONSE: %v\n", err)
		fmt.Printf("BODY: %v\n", res.Body)

		return nil
	}

	if (res.StatusCode != http.StatusOK) || !resp.Ok {
		fmt.Printf("ERROR IN SENDMESSAGE WHILE SENDING MESSAGE: %s\n", resp.Description)
		fmt.Printf("TOKEN: %s\n", token)
		fmt.Printf("REQUEST: %v\n", req)

		return fmt.Errorf("error while sending message: %s", resp.Description)
	}

	return nil
}
