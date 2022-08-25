serve:
	echo "Serving api"

build:
	echo "Build"

up:
	echo "Docker up"
	dock

test:
	go test -tags testing ./...

mock-local:
	mockgen -source=usecase/user/interface.go -destination=usecase/user/mock/user.go -package=mock