generate_grpc_code_backend:
	# Ensure the `invoicer` directory exists for the Go output
	protoc -I=. --go_out=./Backend/invoicer --go_opt=paths=source_relative --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative invoicer.proto

generate_grpc_code_frontend:
	# Make sure the paths are correct for the frontend files
	protoc -I=. --js_out=import_style=commonjs:./frontend/src --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src invoicer.proto
