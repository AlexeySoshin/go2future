git clone https://go.googlesource.com/go
cd ./go/src
git checkout dev.go2go
./make.bash
cd ../..
export GOROOT=$(pwd)/go
go version
cd ./src
go tool go2go test