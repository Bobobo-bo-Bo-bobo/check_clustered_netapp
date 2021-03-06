GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

depend:
	#

build:
	env GOPATH=$(GOPATH) GOBIN=${GOPATH}/bin go install ./...

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/lib64/nagios/plugins

strip: build
	strip --strip-all $(BINDIR)/check_aggregate_state
	strip --strip-all $(BINDIR)/check_aggregate_utilisation
	strip --strip-all $(BINDIR)/check_disk_state
	strip --strip-all $(BINDIR)/check_shelf_state
	strip --strip-all $(BINDIR)/check_snapmirror_relationship_state
	strip --strip-all $(BINDIR)/check_volume_state
	strip --strip-all $(BINDIR)/check_volume_utilisation

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/check_aggregate_state $(DESTDIR)/usr/lib64/nagios/plugins
	install -m 0755 $(BINDIR)/check_aggregate_utilisation $(DESTDIR)/usr/lib64/nagios/plugins
	install -m 0755 $(BINDIR)/check_disk_state $(DESTDIR)/usr/lib64/nagios/plugins
	install -m 0755 $(BINDIR)/check_shelf_state $(DESTDIR)/usr/lib64/nagios/plugins
	install -m 0755 $(BINDIR)/check_snapmirror_relationship_state $(DESTDIR)/usr/lib64/nagios/plugins
	install -m 0755 $(BINDIR)/check_volume_state $(DESTDIR)/usr/lib64/nagios/plugins
	install -m 0755 $(BINDIR)/check_volume_utilisation $(DESTDIR)/usr/lib64/nagios/plugins

clean:
	/bin/rm -f bin/check_aggregate_state
	/bin/rm -f bin/check_aggregate_utilisation
	/bin/rm -f bin/check_disk_state
	/bin/rm -f bin/check_shelf_state
	/bin/rm -f bin/check_snapmirror_relationship_state
	/bin/rm -f bin/check_volume_state
	/bin/rm -f bin/check_volume_utilisation

distclean: clean
	#

uninstall:
	/bin/rm -f $(DESTDIR)/usr/lib64/nagios/plugins

all: depend build strip install

