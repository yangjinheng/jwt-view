rsrc -manifest main.manifest -ico favicon.ico -o main.syso
go build -ldflags -H=windowsgui
