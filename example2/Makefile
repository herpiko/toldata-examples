gen:
	echo $(shell pwd)
	docker run --platform linux/amd64 -v $(shell pwd):/gen -v $(shell pwd)/api:/api \
		citradigital/toldata:v0.1.4 \
		-I /api /api/example2.proto \
		--toldata_out=grpc:/gen --gogofaster_out=plugins=grpc:/gen
