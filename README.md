# go-healthcheck-homework

Just a homework project. 
If you call GET /health it will return `{"status": "ok"}` json

Dockerfile checked with hadolint 😎😎

## Installation and Usage with Docker

`docker pull russbalabanov/docker-homework:latest`
`docker run -p 8000:8000 russbalabanov/docker-homework:latest`
`curl -h http://127.0.0.1:8000/health -v`

## Installation and Usage with K8S

First the namespace: `kubectl create namespace hw`
Next, let's use this namespace as a default one `kubectl config set-context --current --namespace=hw`

Make sure you have your hosts contain `127.0.0.1 arch.homework`

### Helm for Nginx-Ingress

`helm repo add nginx-stable https://helm.nginx.com/stable`
`helm repo update`
`helm install nginx-ingress-controller nginx-stable/nginx-ingress --set controller.service.httpPort.port=80 --set controller.enableSnippets=true`


### Testing

Loading all manifests with one command `kubectl apply -f .`

Normal usage: `curl arch.homework/health`
Using rewrite: `curl arch.homework/otusapp/ruslan_balabanov`

### Troubleshooting
Sometimes Ingress might work funky with 80th port, changing it to a different port does the job. For example like this `--set controller.service.httpPort.port=8080` 