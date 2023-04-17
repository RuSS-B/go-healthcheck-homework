include .make.env

run:
	go run *.go

namespace:
	kubectl create namespace hw

build.app:
	docker build -f Dockerfile -t ${APP_IMAGE} .

push.app:
	docker push ${APP_IMAGE}

build.migrations:
	docker build -f Migrations.Dockerfile -t ${MIGRATIONS_IMAGE} .

push.migrations:
	docker push ${MIGRATIONS_IMAGE}

helm.install:
	helm install homework3 ./homework3 --values homework3/values.yaml

helm.delete:
	helm delete homework3

helm.upgrade:
	helm upgrade homework3 ./homework3 --values homework3/values.yaml

helm.deps.upd:
	helm dependency update ./homework3

helm.deps:
	helm dependency build ./homework3

dashboard.token:
	kubectl -n kubernetes-dashboard create token admin-user