build:
	cd bin && go build -o lxr /home/jack/LXR/LXR-cli/cmd/cli
run:
	cd bin && ./lxr

cpbin:
	cp /home/jack/LXR/LXR-cli/bin/lxr /usr/bin