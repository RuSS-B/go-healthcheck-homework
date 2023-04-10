include .make.env

run:
	go run *.go

build.app:
	docker build -f Dockerfile -t ${APP_IMAGE} .

push.app:
	docker push ${APP_IMAGE}

build.migrations:
	docker build -f Migrations.Dockerfile -t ${MIGRATIONS_IMAGE} .

push.migrations:
	docker push ${MIGRATIONS_IMAGE}

helm.upgrade:
	helm upgrade homework3 ./homework3 --values homework3/values.yaml
