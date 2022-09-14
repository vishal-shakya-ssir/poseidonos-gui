module dagent

go 1.14

replace kouros => ./../kouros

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1
	kouros v0.0.0-00010101000000-000000000000
)
