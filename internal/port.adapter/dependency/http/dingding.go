package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

    "quants/config"
    "quants/util/logger"
)

/**
 * @author Rancho
 * @date 2021/12/3
 */

type DingDingMessage struct {
	Msgtype string `json:"msgtype"`
	At      struct {
		AtMobiles []string    `json:"atMobiles"`
		IsAtAll   interface{} `json:"isAtAll"`
	} `json:"at"`
	Text struct {
		Content interface{} `json:"content"`
	} `json:"text"`
    IsAtAll bool
    Message string
}

type DingDingResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (d DingDingResponse) String() string {
	return fmt.Sprintf("{errcode: %d, errmsg: %s}", d.Errcode, d.Errmsg)
}

func (d *DingDingMessage) Marshal() []byte {
	d.Msgtype = "text"
	d.At = struct {
		AtMobiles []string    `json:"atMobiles"`
		IsAtAll   interface{} `json:"isAtAll"`
	}{
		AtMobiles: []string{"1111"},
		IsAtAll:   d.IsAtAll,
	}
	d.Text = struct {
		Content interface{} `json:"content"`
	}{
		Content: "Notice.\n" + d.Message,
	}

	b, err := json.Marshal(d)
	if err != nil {
		logger.Log.Errorf(context.Background(), "json.Marshal fail when DingDingMessage.Marshal, msg: %s, err: %v", d.Message, err)
	}

	return b
}

func SendDingDingMessage(msg string, isAtAll bool) bool {
	client := &http.Client{}

	d := DingDingMessage{
		Msgtype: "text",
        IsAtAll: isAtAll,
	}
	body := d.Marshal()

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/robot/send", DingDingAPI), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	query := req.URL.Query()
	query.Add("access_token", config.Config.DingDing.AccessToken)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Errorf(context.Background(), "DingDingAPI POST /robot/send err: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorf(context.Background(), "read DingDingAPI body err: %v", err)
	}
	dingResp := DingDingResponse{}
	err = json.Unmarshal(respBody, &dingResp)
	if err != nil {
		logger.Log.Errorf(context.Background(), "json.Unmarshal DingDingAPI body err: %v", err)
	}
	if dingResp.Errcode != 0 {
		logger.Log.Errorf(context.Background(), "SendDingDingMessage fail. msg: %v, dingResp: %v", msg, dingResp)
        return false
	}

    return true
}
