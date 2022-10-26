docker build . -t pre-test:v1.0.0
docker run -p 8080:8080 pre-test:v1.0.0