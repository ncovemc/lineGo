package login

import (
	"log"

	con "../config"
	talk "../talkservice"
	"github.com/apache/thrift/lib/go/thrift"
)

func createSession(authToken, appName string) *talk.TalkServiceClient {
	var trans thrift.TTransport
	var err error
	trans, err = thrift.NewTHttpClient(con.LINE_HOST + con.TALK_PATH)
	if err != nil {
		log.Fatal(err)
	}
	var connect *thrift.THttpClient
	connect = trans.(*thrift.THttpClient)
	connect.SetHeader("X-Line-Access", authToken)
	connect.SetHeader("User-Agent", con.GetUserAgent(appName))
	connect.SetHeader("X-Line-Application", con.GetLineApp(appName))
	connect.SetHeader("x-lal", "ja_jp")
	protocol := thrift.NewTCompactProtocolFactory().GetProtocol(connect)
	return talk.NewTalkServiceClientProtocol(connect, protocol, protocol)
}

func createSession_poll(authToken, appName string) *talk.TalkServiceClient {
	var trans thrift.TTransport
	var err error
	trans, err = thrift.NewTHttpClient(con.LINE_HOST + con.POLL_PATH)
	if err != nil {
		log.Fatal(err)
	}
	var connect *thrift.THttpClient
	connect = trans.(*thrift.THttpClient)
	connect.SetHeader("X-Line-Access", authToken)
	connect.SetHeader("User-Agent", con.GetUserAgent(appName))
	connect.SetHeader("X-Line-Application", con.GetLineApp(appName))
	connect.SetHeader("x-lal", "ja_jp")
	protocol, err := thrift.NewTTransportFactory().GetTransport(connect)
	return talk.NewTalkServiceClientFactory(protocol, thrift.NewTCompactProtocolFactory())
}

func Talk(token, appName string) *talk.TalkServiceClient {
	return createSession(token, appName)
}

func Poll(token, appName string) *talk.TalkServiceClient {
	return createSession_poll(token, appName)
}
