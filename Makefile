PKGNAME    = go-base
DESTDIR   ?=
PREFIX    ?= /usr/local
BINDIR     = $(PREFIX)/bin
SYSCONFDIR = /etc
SYSTEMD   ?= $(SYSCONFDIR)/systemd/system

GOPROJROOT  = $(GOSRC)/$(PROJREPO)

GOLDFLAGS   = -ldflags "-s -w"
GOCC        = CGO_ENABLED=0 go
GOFMT       = $(GOCC) fmt -x
GOGET       = $(GOCC) get $(GOLDFLAGS)
GOBUILD     = $(GOCC) build -v $(GOLDFLAGS)
GOTEST      = $(GOCC) test
GOVET       = $(GOCC) vet
GOINSTALL   = $(GOCC) install $(GOLDFLAGS)

EXES = base64 basename cat cksum echo false true tty whoami yes

include Makefile.waterlog

GOLINT = golint -set_exit_status

all: build

build:
	@$(call stage,BUILD)
	@$(GOBUILD)
	@$(call pass,BUILD)

test: build
	@$(call stage,TEST)
	@$(GOTEST) ./...
	@$(call pass,TEST)

validate:
	@$(call stage,FORMAT)
	@$(GOFMT) ./... @$(call pass,FORMAT)
	@$(call stage,VET)
	@$(call task,Running 'go vet'...)
	@$(GOVET) ./... @$(call pass,VET)
	@$(call stage,LINT)
	@$(call task,Running 'golint'...)
	@$(GOLINT) ./...
	@$(call pass,LINT)

install:
	@$(call stage,INSTALL)
	@$(call task,Installing 'go-base'...)
	install -Dm 00755 $(PKGNAME) $(DESTDIR)$(BINDIR)/$(PKGNAME)
	@$(call task,Creating symlinks...)
	@for exe in $(EXES); do \
	    ln -s $(PKGNAME) $(DESTDIR)$(BINDIR)/$$exe; \
	done
	@$(call pass,INSTALL)

uninstall:
	@$(call stage,UNINSTALL)
	@$(call task,Removing 'go-base'...)
	rm -f $(DESTDIR)$(BINDIR)/$(PKGNAME)
	@$(call task,Removing symlinks...)
	@for exe in $(EXES); do \
	    unlink $(DESTDIR)$(BINDIR)/$$exe; \
	done
	@$(call pass,UNINSTALL)

clean:
	@$(call stage,CLEAN)
	@$(call task,Removing man-pages...)
	@rm *.1
	@$(call task,Removing executable...)
	@rm $(PKGNAME)
	@$(call pass,CLEAN)
