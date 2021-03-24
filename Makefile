build:
	CGO_ENABLED=0 go build -o dist/podctl . 
	chmod +x dist/podctl

