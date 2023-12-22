protoc:
	cd ../ &&\
	protoc --proto_path=go-tiktok-new/idl --go_out=. go-tiktok-new/idl/*