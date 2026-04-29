build-windows:
    # cross compilation from Linux -> Windows
    # see https://github.com/go-gl/gl/issues/155#issuecomment-2480073175
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build
