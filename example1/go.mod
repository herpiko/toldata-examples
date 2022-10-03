module example1

go 1.18

require (
	example2 v0.0.0-00010101000000-000000000000
	github.com/citradigital/toldata v0.1.5
	github.com/gogo/protobuf v1.2.1
	github.com/nats-io/go-nats v1.7.2
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092
	google.golang.org/grpc v1.23.0
)

require (
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/nats-io/nkeys v0.1.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/sys v0.0.0-20190412213103-97732733099d // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 // indirect
	google.golang.org/protobuf v1.23.0 // indirect
)

replace example2 => ./external/example2
