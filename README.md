# go-healthcheck-homework

Just a homework project. 
If you call GET /health it will return `{"status": "ok"}` json

Dockerfile checked with hadolint 😎😎

## Installation and Usage

`docker pull russbalabanov/docker-homework:latest`
`docker run -p 8000:8000 russbalabanov/docker-homework:latest`
`curl -h http://127.0.0.1:8000/health -v`
