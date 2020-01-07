.PHONY: proto

proto:
	prototool format -w -f ./proto
	prototool generate --debug ./proto
	protoc-go-inject-tag -XXX_skip=bson,csv -input=./pkg/entity/v1/user.pb.go
	protoc-go-inject-tag -XXX_skip=bson,csv -input=./pkg/entity/v1/session.pb.go
	prototool lint ./proto