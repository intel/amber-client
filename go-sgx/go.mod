module github.com/intel/amber/v1/client/sgx

go 1.17

require github.com/intel/amber/v1/client v0.0.0

require (
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)

replace github.com/intel/amber/v1/client => ../go-client