PKGS ?= $(shell go list ./...)

run:
	go run main.go

generate-mock:
	cd services && mockery --name=IpasswordValidator --filename=validatePassword.go --outpkg=mock --output=../mock