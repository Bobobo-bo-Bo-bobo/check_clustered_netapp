GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

depend:
	#

build:
	env GOPATH=$(GOPATH) GOBIN=${GOPATH}/bin go install ./...

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/lib64/nagios/plugins

strip: build
	strip --strip-all $(BINDIR)/check_shelf_state
	strip --strip-all $(BINDIR)/check_volume_state
	strip --strip-all $(BINDIR)/check_volume_utilisation

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/check_volume_state $(DESTDIR)/usr/lib64/nagios/plugins
	install -m 0755 $(BINDIR)/check_volume_utilisation $(DESTDIR)/usr/lib64/nagios/plugins

clean:
	/bin/rm -f bin/check_shelf_state
	/bin/rm -f bin/check_volume_state
	/bin/rm -f bin/check_volume_utilisation

distclean: clean
	#

uninstall:
	/bin/rm -f $(DESTDIR)/usr/lib64/nagios/plugins

all: depend build strip install

