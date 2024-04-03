.PHONY:
.SILENT:



run:
	go run .\cmd\app\main.go --cfg=.\configs\prod.yaml


migrate:
	go run .\cmd\migrate\main.go --cfg=./configs/prod.yaml --migrations=./migrations/