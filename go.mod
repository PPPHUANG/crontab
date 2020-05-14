module crontab

go 1.14

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	go.mongodb.org/mongo-driver v1.3.2
	go.uber.org/zap v1.15.0 // indirect
	google.golang.org/genproto v0.0.0-20200428115010-c45acf45369a // indirect
	google.golang.org/grpc v1.29.1 // indirect
)

replace (
    github.com/coreos/go-systemd => /Users/huangpeng/go/src/github.com/go-systemd
    google.golang.org/grpc => google.golang.org/grpc v1.26.0
)