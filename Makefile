GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

PROGRAMS = check_volume_utilisation

depend:
	#

build:
	env GOPATH=$(GOPATH) GOBIN=${GOPATH}/bin go install ./...

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/lib64/nagios/plugins

strip: build
	strip --strip-all $(BINDIR)/$(PROGRAMS)

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/$(PROGRAMS) $(DESTDIR)/usr/lib64/nagios/plugins

clean:
	/bin/rm -f bin/$(PROGRAMS)

distclean: clean
	/bin/rm -rf src/github.com/

uninstall:
	/bin/rm -f $(DESTDIR)/usr/lib64/nagios/plugins

all: depend build strip install

