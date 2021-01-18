BIN = golvgl
GOCMD = go
GOBUILD = ${GOCMD} build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"'
GOCLEAN = ${GOCMD} clean

build:
	${GOBUILD} -o ${BIN} -v *.go

clean:
	${GOCLEAN}
