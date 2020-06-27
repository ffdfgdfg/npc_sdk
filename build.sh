cd /go
go get golang.org/x/mobile/cmd/gomobile
gomobile init
echo "gomobile install finish"
go get -d -v ehang.io/nps
cd /go/src/ehang.io/nps
go mod vendor
cd vendor
cp -R * /go/src
cd ..
rm -rf vendor
echo "dependences install finish"
mkdir -p /go/src/github.com/ffdfgdfg/npc_sdk
cp -R /app/* /go/src/github.com/ffdfgdfg/npc_sdk
cd /go/src/github.com/ffdfgdfg/npc_sdk
gomobile bind github.com/ffdfgdfg/npc_sdk/npc
echo "bind finish"
cp npc.aar /app
cp npc-sources.jar /app
