export K8S_DASHBOARD_NAMESPACE=kubernetes-dashboard

kubectl create namespace $K8S_DASHBOARD_NAMESPACE
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --namespace=$K8S_DASHBOARD_NAMESPACE

kubectl apply -f ./dashboard-adminuser.yaml --namespace=$K8S_DASHBOARD_NAMESPACE

export POD_NAME=$(kubectl get pods -n kubernetes-dashboard -l "app.kubernetes.io/name=kubernetes-dashboard,app.kubernetes.io/instance=kubernetes-dashboard" -o jsonpath="{.items[0].metadata.name}")
kubectl -n $K8S_DASHBOARD_NAMESPACE port-forward $POD_NAME 8443:8443