package login

import (
	"fmt"

	service "./talkservice"

	con "./config"
	"github.com/apache/thrift/lib/go/thrift"
)

func createSession(authToken, service, appName string) *service.TalkServiceClient {
	if service == "talk" {
		trans, err = thrift.NewTHttpClient(con.LINE_HOST + con.TALK_PATH)
	} else if service == "poll" {
		trans, err = thrift.NewTHttpClient(con.LINE_HOST + con.POLL_PATH)
	}
	if err != nil {
		fmt.Println(err)
	}
	var connect *thrift.THttpClient
	connect = trans.(*thrift.THttpClient)
	connect.SetHeader("X-Line-Access", authToken)
	connect.SetHeader("User-Agent", con.GetUserAgen(appName))
	connect.SetHeader("X-Line-Application", con.GetLineApp(appName))
	connect.SetHeader("x-lal", "ja_jp")
	setProtocol := thrift.NewTCompactProtocolFactory()
	protocol := setProtocol.GetProtocol(connect)
	return service.NewTalkServiceClientProtocol(connect, protocol, protocol)
}

func Talk(token, appName string) *service.TalkServiceClient {
	return createSession(token, "talk", appName)
}

func Poll(token, appName string) *service.TalkServiceClient {
	return createSession(token, "poll", appName)
}
