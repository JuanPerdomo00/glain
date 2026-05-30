PREFIX ?= /usr/local
BINARY = glain

build:
	go build -o $(BINARY) .

install: build
	install -Dm755 $(BINARY) $(DESTDIR)$(PREFIX)/bin/$(BINARY)
	install -Dm644 .gif-index $(DESTDIR)$(HOME)/.config/glain/.gif-index
	install -Dm644 margin.txt $(DESTDIR)$(HOME)/.config/glain/margin.txt

uninstall:
	rm -f $(PREFIX)/bin/$(BINARY)

clean:
	rm -f $(BINARY)

.PHONY: build install uninstall clean
