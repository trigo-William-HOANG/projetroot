# projetroot

docker envoy : docker run -d --name envoy -p 8080:8080 -v $(pwd)/envoy.yaml:/etc/envoy/envoy.yaml envoyproxy/envoy:v1.24.0

docker pull envoyproxy/envoy:v1.24.0