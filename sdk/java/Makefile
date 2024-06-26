GOROOT=$(shell go env GOROOT)

all: package

ceph: libjfs-ceph
	mvn package -B -Dmaven.test.skip=true

libjfs-ceph: ../../pkg/*/*.go libjfs/*.go
	make -C libjfs ceph

libjfs/libjfs: ../../pkg/*/*.go libjfs/*.go
	make -C libjfs

compile:
	mvn compile -B --quiet
test: libjfs
	mvn test -B --quiet
package: libjfs/libjfs
	mvn package -B -Dmaven.test.skip=true

win: win-package package

win-package: ../../pkg/*/*.go libjfs/*.go
	make -C libjfs win

package-all: libjfs-all
	mvn clean package -B -Dmaven.test.skip=true

libjfs-all: libjfs.so
	docker run --rm \
		-v ~/go/pkg/mod:/go/pkg/mod \
		-v ~/work/juicefs/juicefs:/go/src/github.com/leonatone/juicefs \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-w /go/src/github.com/leonatone/juicefs/sdk/java/libjfs \
		--entrypoint=/bin/bash \
		leonatone/golang-cross:latest \
		-c 'make mac win linux-arm64 mac-arm64'

libjfs.so:
	docker run --rm \
		-v ~/go/pkg/mod:/go/pkg/mod \
		-v $(GOROOT):/go \
        -v ~/work/juicefs/juicefs:/go/src/github.com/leonatone/juicefs \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -w /go/src/github.com/leonatone/juicefs/sdk/java/libjfs \
        leonatone/sdk-builder \
        /bin/bash -c 'make'
