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

type IDingDingAPI interface {
    SendDingDingMessage(string, bool) bool
}

type dingDingClient struct {

}

var (
    DingDingClient IDingDingAPI = &dingDingClient{}
)

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

func (d *DingDingMessage) marshal() []byte {
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
		logger.Log.Errorf(context.Background(), "json.Marshal fail when DingDingMessage.marshal, msg: %s, err: %v", d.Message, err)
	}

	return b
}

func (d *dingDingClient)SendDingDingMessage(msg string, isAtAll bool) bool {
	client := &http.Client{}

	dm := DingDingMessage{
        Message: msg,
		Msgtype: "text",
        IsAtAll: isAtAll,
	}
	body := dm.marshal()

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/robot/send", DingDingAPIUrl), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	query := req.URL.Query()
	query.Add("access_token", config.Config.DingDing.AccessToken)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Errorf(context.Background(), "DingDingClient POST /robot/send err: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorf(context.Background(), "read DingDingClient body err: %v", err)
	}
	dingResp := DingDingResponse{}
	err = json.Unmarshal(respBody, &dingResp)
	if err != nil {
		logger.Log.Errorf(context.Background(), "json.Unmarshal DingDingClient body err: %v", err)
	}
	if dingResp.Errcode != 0 {
		logger.Log.Errorf(context.Background(), "SendDingDingMessage fail. msg: %v, dingResp: %v", msg, dingResp)
        return false
	}

    return true
}
