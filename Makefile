OBJ_NAME = 
LDFLAGS = 
PKG_NAME =
ARCH =
OS = 
scanner:
	$(eval PKG_NAME+= "scanner")
	$(eval ARCH += "amd64")
	$(eval OS += "linux")
	$(eval LDFLAGS += "-w -s")
	$(eval OBJ_NAME += port-scanner)
	cd ./cmd/$(PKG_NAME); GOOS=$(OS) GOARCH=$(ARCH) go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin 
doc:
	cd ./cmd/; godoc -http=:6060
