port=81
echo "http://localhost:$port"
docker run -p$port:8080 -P cloudtest
