build:
	CGO_ENABLED=0 go build -o podctl .
	chmod +x podctl

