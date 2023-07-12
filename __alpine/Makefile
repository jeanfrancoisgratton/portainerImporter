all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./p2r.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.20.6 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/p2r.exe $(DESTDIR)$(PREFIX)/bin/p2r
	strip $(DESTDIR)$(PREFIX)/bin/p2r
