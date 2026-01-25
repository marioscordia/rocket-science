module github.com/marioscordia/rocket-science/payment

go 1.25.1

replace github.com/marioscordia/rocket-science/shared => ../shared

require (
	github.com/google/uuid v1.6.0
	github.com/marioscordia/rocket-science/shared v0.0.0-20260123092452-744905b44df8
)

require (
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/grpc v1.78.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
