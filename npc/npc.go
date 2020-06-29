package npc

import (
	"fmt"
)

//var cl *client.TRPClient

//StartClientByVerifyKey
func StartClientByVerifyKey(serverAddr, verifyKey, connType, proxyUrl string) {
// 	cl := client.NewRPClient(serverAddr, verifyKey, connType, proxyUrl, nil, 60)
// 	cl.Start()
}

//GetClientStatus
func GetClientStatus() int {
// 	return client.NowStatus
	return 0
}

//CloseClient
//func CloseClient() {
//	if cl != nil {
//		cl.Close()
//	}
//}

//Version
func Version() string {
// 	return version.VERSION
	return "test"
}

//Logs
//func Logs() string {
//	return common.GetLogMsg()
//}

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
