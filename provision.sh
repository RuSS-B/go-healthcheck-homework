kubectl create namespace hw
kubectl config set-context --current --namespace=hw

helm install -f homework3/values.yaml homework3 ./homework3/

