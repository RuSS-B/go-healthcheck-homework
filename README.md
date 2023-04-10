# go-homework

Just a homework project. 
If you call GET /health it will return `{"status": "ok"}` json
You also can call /users endpoint to do CRUD operations with GET / POST / PUT / DELETE methods

Dockerfiles checked with hadolint 😎😎

## Installation

### Preparations
First the namespace: `kubectl create namespace hw`
Next, let's use this namespace as a default one `kubectl config set-context --current --namespace=hw`

Make sure you have your hosts contain `127.0.0.1 arch.homework`

### Nginx-Ingress via Helm
`helm repo add nginx-stable https://helm.nginx.com/stable`

`helm repo update`

`helm install nginx-ingress-controller nginx-stable/nginx-ingress --set controller.service.httpPort.port=2080 --set controller.enableSnippets=true`

### Installation of the main project via Helm
`helm install hw3 ./homework3`

### Dashboard included
For login, please use this url: https://127.0.0.1:8443/#/login

Command to generate a token: `kubectl -n kubernetes-dashboard create token admin-user`

### Manual testing
Normal usage: `curl arch.homework:2080/users/1`

Note: I've used 2080 port just to avoid any port 80 issues on localhost, feel free to configure your ingress with port 80 instrad 