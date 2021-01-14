PKGNAME    = go-base
DESTDIR   ?=
PREFIX    ?= /usr/local
BINDIR     = $(PREFIX)/bin
DATADIR    = $(PREFIX)/share/$(PKGNAME)
MANDIR     = $(PREFIX)/share/man/man1

GOPROJROOT  = $(GOSRC)/$(PROJREPO)

GOLDFLAGS   = -ldflags "-s -w"
GOCC        = go
GOFMT       = $(GOCC) fmt -x
GOGET       = $(GOCC) get $(GOLDFLAGS)
GOBUILD     = CGO_ENABLED=0 $(GOCC) build -v $(GOLDFLAGS)
GOTEST      = $(GOCC) test
GOVET       = $(GOCC) vet
GOINSTALL   = $(GOCC) install $(GOLDFLAGS)

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
	@$(GOFMT) ./...
	@$(call pass,FORMAT)
	@$(call stage,VET)
	@$(call task,Running 'go vet'...)
	@$(GOVET) ./...
	@$(call pass,VET)
	@$(call stage,LINT)
	@$(call task,Running 'golint'...)
	@$(GOLINT) ./...
	@$(call pass,LINT)

install:
	@$(call stage,INSTALL)
	@$(call task,Installing 'go-base'...)
	install -Dm 00755 $(PKGNAME) $(DESTDIR)$(BINDIR)/$(PKGNAME)
	@$(call task,Creating symlinks...)
	@ln -s go-base gen-single-links
	./gen-single-links $(DESTDIR)$(BINDIR)
	@rm gen-single-links
	@$(call task,Creating manpages...)
	@ln -s go-base gen-man-pages
	./gen-man-pages
	@rm gen-man-pages
	@$(call task,Installing manpages...)
	install -t $(DESTDIR)$(MANDIR) -Dm00644 *.1
	@rm *.1
	@$(call pass,INSTALL)

clean:
	@$(call stage,CLEAN)
	@$(call task,Removing man-pages...)
	@rm *.1 || true
	@$(call task,Removing executable...)
	@rm $(PKGNAME)
	@$(call pass,CLEAN)
