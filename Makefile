serve:
	echo "Serving api"

build:
	echo "Build"

up:
	echo "Docker up"
	dock

test:
	go test -tags testing ./...