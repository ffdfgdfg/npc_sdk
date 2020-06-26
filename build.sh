cd /go
go get golang.org/x/mobile/cmd/gomobile
gomobile init
go get -d ehang.io/nps
mkdir -p /go/src/github.com/ffdfgdfg/npc_sdk/npc
cp -R /app/* /go/src/github.com/ffdfgdfg/npc_sdk/npc
cd /go/src/github.com/ffdfgdfg/npc_sdk/npc
gomobile bind github.com/ffdfgdfg/npc_sdk/npc
cp npc.aar /app
cp npc-sources.jar /app