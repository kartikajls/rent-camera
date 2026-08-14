package dto

type SendWAResponse struct {
	MsgID  int64  `json:"msgId"`
	JID    string `json:"jid"`
	Status string `json:"status"`
}
