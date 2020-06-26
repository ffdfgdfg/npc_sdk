package npc

import (
	"ehang.io/nps/client"
	"ehang.io/nps/lib/version"
	"fmt"
)

//var cl *client.TRPClient

//StartClientByVerifyKey
func StartClientByVerifyKey(serverAddr, verifyKey, connType, proxyUrl string) {
	cl := client.NewRPClient(serverAddr, verifyKey, connType, proxyUrl, nil, 60)
	cl.Start()
}

//GetClientStatus
func GetClientStatus() int {
	return client.NowStatus
}

//CloseClient
//func CloseClient() {
//	if cl != nil {
//		cl.Close()
//	}
//}

//Version
func Version() string {
	return version.VERSION
}

//Logs
//func Logs() string {
//	return common.GetLogMsg()
//}

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
