install_cross:
	go install github.com/fyne-io/fyne-cross@latest

build:
	go build -tags wayland -o controllercontrol

build_darwin:
	GOOS=darwin GOARCH=amd64 go build -o controllercontrol-darwin
