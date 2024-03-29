all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./portainerImporter.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.22.1 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/portainerImporter.exe $(DESTDIR)$(PREFIX)/bin/portainerImporter
	strip $(DESTDIR)$(PREFIX)/bin/portainerImporter
