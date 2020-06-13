package login

import (
	"log"
	talk "../talkservice"
	con "../config"
	"github.com/apache/thrift/lib/go/thrift"
)

func createSession(authToken, service, appName string) *talk.TalkServiceClient {
	var trans thrift.TTransport
	var err error
	if service == "talk" {
		trans, err = thrift.NewTHttpClient(con.LINE_HOST + con.TALK_PATH)
	} else if service == "poll" {
		trans, err = thrift.NewTHttpClient(con.LINE_HOST + con.POLL_PATH)
	}
	if err != nil {
		log.Fatal(err)
	}
	var connect *thrift.THttpClient
	connect = trans.(*thrift.THttpClient)
	connect.SetHeader("X-Line-Access", authToken)
	connect.SetHeader("User-Agent", con.GetUserAgent(appName))
	connect.SetHeader("X-Line-Application", con.GetLineApp(appName))
	connect.SetHeader("x-lal", "ja_jp")
	setProtocol := thrift.NewTCompactProtocolFactory()
	protocol := setProtocol.GetProtocol(connect)
	return talk.NewTalkServiceClientProtocol(connect, protocol, protocol)
}

func Talk(token, appName string) *talk.TalkServiceClient {
	return createSession(token, "talk", appName)
}

func Poll(token, appName string) *talk.TalkServiceClient {
	return createSession(token, "poll", appName)
}
