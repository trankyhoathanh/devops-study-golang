docker build -t go-fasthttp-test .
docker run -p 8081:8080 --name test go-fasthttp-test