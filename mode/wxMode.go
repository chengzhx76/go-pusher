package mode

import "encoding/xml"

type CommonError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

/*
<?xml version="1.0" encoding="utf-8"?>

<xml>
  <ToUserName><![CDATA[gh_5bdec3e424d4]]></ToUserName>
  <FromUserName><![CDATA[o1Zfm1ODIStd268G04DKuNCVNLWM]]></FromUserName>
  <CreateTime>1554210101</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>

  <Event><![CDATA[SCAN]]></Event>
  <EventKey><![CDATA[cheng]]></EventKey>
  <Ticket><![CDATA[gQFi8DwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAyc25BX0lKbG5mQ2sxRVRtSU5zY00AAgS3W6NcAwSAOgkA]]></Ticket>
</xml>

<?xml version="1.0" encoding="utf-8"?>

<xml>
  <ToUserName><![CDATA[gh_5bdec3e424d4]]></ToUserName>
  <FromUserName><![CDATA[o1Zfm1ODIStd268G04DKuNCVNLWM]]></FromUserName>
  <CreateTime>1554212481</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>

  <Content><![CDATA[好好]]></Content>
  <MsgId>22250973400809876</MsgId>
</xml>


*/
type CommonWxMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}
type WxScanEvent struct {
	CommonWxMsg
	Event    string `xml:"Event"`
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}
type WxTextEvent struct {
	CommonWxMsg
	Content string `xml:"Content"`
	MsgId   string `xml:"MsgId"`
}

type PushMessage struct {
	touser      string
	template_id string
	url         string
	data        PushMessageData
}

type PushMessageData struct {
	title PushMessageDataValue
	text  PushMessageDataValue
	desc  PushMessageDataValue
}

type PushMessageDataValue struct {
	value string
}
