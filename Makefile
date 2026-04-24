build:
	cd client && go build -o lxr-cli . && ./lxr-cli
run:
	cd client && ./lxr-cli