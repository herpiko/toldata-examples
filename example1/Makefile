gen:
	echo $(shell pwd)
	docker run --platform linux/amd64 -v $(shell pwd)/external/example2:/api citradigital/toldata:v0.1.4 \
        -I /api/ \
        /api/example2.proto \
        --toldata_out=grpc:/api --gogofaster_out=plugins=grpc:/api
	docker run --platform linux/amd64 -v $(shell pwd):/gen -v $(shell pwd)/api:/api \
		citradigital/toldata:v0.1.4 \
		-I /api /api/example1.proto \
		--toldata_out=grpc:/gen --gogofaster_out=plugins=grpc:/gen
