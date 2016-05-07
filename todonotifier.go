// TodoNotifier
package main

//SOAP API リクエストのテンプレート
const SOAPTemplate {
	<?xml version="1.0" encoding="UTF-8"?>
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope">
	<soap:Header>
		<Action>{{.Action}}</Action>
    	<Security></Security>
    	<Timestamp>
			<Created>2010-08-12T14:45:00Z</Created>
			<Expires>2037-08-12T14:45:00Z</Expires>
		</Timestamp>
		<Locale>jp</Locale>
	</soap:Header>
	<soap:Body>
    	<{{.Action}}>
	    	{{.Parameters}}
    	</{{.Action}}>
	</soap:Body>
	</soap:Envelope>
}

//POST処理先
const LoginAPI = "https://madtak.cybozu.com/g/util_api/util/api.csp?"
const ScheduleAPI = "https://madtak.cybozu.com/g/cbpapi/schedule/api.csp?"


import (
	"fmt"
	"net/http"
)
//XMLをパースするための構造体
type XMLexclusive struct {
	XMLName xml.Name `xml:"exclusive_datetime"`
	Start   string   `xml:"start,attr"`
	End     string   `xml:"end,attr"`
}
type XMLrepeat_condition struct {
	XMLName   xml.Name `xml:"condition"`
	Type      string   `xml:"type,attr"`
	Day       string   `xml:"day,attr"`
	Week      string   `xml:"week,attr"`
	StartDate string   `xml:"start_date,attr"`
	EndDate   string   `xml:"end_date,attr"`
	StartTime string   `xml:"start_time,attr"`
	EndTime   string   `xml:"end_time,attr"`
}
type XMLrepeat struct {
	XMLName   xml.Name             `xml:"repeat_info"`
	Condition *XMLrepeat_condition `xml:"condition"`
	Exclusive []*XMLexclusive      `xml:"exclusive_datetimes>exclusive_datetime"`
}

type XMLdatetime struct {
	XMLName xml.Name `xml:"datetime"`
	Start   string   `xml:"start,attr"`
	End     string   `xml:"end,attr"`
}
type XMLschedule_event struct {
	XMLName    xml.Name     `xml:"schedule_event"`
	Id         int          `xml:"id,attr"`
	EventType  string       `xml:"event_type,attr"`
	PublicType string       `xml:"public_type,attr"`
	Plan       string       `xml:"plan,attr"`
	Detail     string       `xml:"detail,attr"`
	TimeZone   string       `xml:"timezone,attr"`
	Version    int          `xml:"version,attr"`
	Allday     bool         `xml:"allday,attr"`
	StartOnly  bool         `xml:"start_only,attr"`
	Datetime   *XMLdatetime `xml:"when>datetime"`
	Repeat     *XMLrepeat   `xml:"repeat_info"`
}
type XMLSoap struct {
	XMLName        xml.Name             `xml:"Envelope"`
	Schedule_event []*XMLschedule_event `xml:"Body>ScheduleGetEventsResponse>returns>schedule_event"`  // 「>」で接続することでelementツリーをその順にたどってくれるようです
}
type Request struct {
	requestType http.Response
	url string
}

//TODO : mainをUIスレッドに変更
//認証実行、スケジュールの取得処理を行う
func main() {
	
}

//SOAP API 実行で受け取ったcookieを返す
func execLogin()string{
	
}
//スケジュールを取得してXMLデコードした構造体のポインタ配列を返す
func getSchedule () []*XMLschedule_event {
	//Templateを利用してパラメータセット
	
	//httpRequest呼び出し
	
	//decodeXML呼び出し
}
//XMLデコード処理
func decodeXML (res string) []*XMLschedule_event {
	
}
//HTTPリクエスト処理
func httpRequest(req Request){
	
}

//例外処理
func catchError(err error){
	
}
