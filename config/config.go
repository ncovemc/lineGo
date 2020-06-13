package config

var LINE_HOST string = "https://ga2.line.naver.jp"
var TALK_PATH string = "/S4"
var POLL_PATH string = "/P4"
var SYSTEM_NAME string = "SakuraGo"

var SystemVer map[string]string = map[string]string{
	"IOS": "13.4.1",
	"MAC": "10.15.1",
}

var AppVer map[string]string = map[string]string{
	"IOS": "10.6.5",
	"MAC": "5.24.1",
}

func GetLineApp(appName string) string {
	switch appName {
	case "IOS":
		return "IOS\t" + AppVer["IOS"] + "\tiOS\t" + SystemVer["IOS"]
	case "MAC":
		return "DESKTOPMAC\t" + AppVer["MAC"] + "\tOS X\t" + SystemVer["MAC"]
	default:
		return "DESKTOPMAC"
	}
}

func GetUserAgent(appName string) string {
	switch appName {
	case "IOS":
		return "Line/" + AppVer["IOS"] + " iPhone8,1 " + SystemVer["IOS"]
	case "MAC":
		return "Line/" + AppVer["MAC"]
	default:
		return "Line/" + AppVer["MAC"]
	}
}
