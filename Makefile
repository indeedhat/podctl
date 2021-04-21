build:
	CGO_ENABLED=0 go build -o podctl .
	chmod +x podctl

install:
	CGO_ENABLED=0 go install 

