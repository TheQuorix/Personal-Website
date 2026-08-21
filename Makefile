.PHONY: back-build back-run back-test back-fmt back-lint front-build front-run front-dev front-prepare

back-build:
	$(MAKE) -C backend build

back-run:
	$(MAKE) -C backend run

back-test:
	$(MAKE) -C backend test

back-fmt:
	$(MAKE) -C backend test

back-lint:
	$(MAKE) -C backend lint


front-build:
	$(MAKE) -C frontend build

front-run:
	$(MAKE) -C frontend run

front-dev:
	$(MAKE) -C frontend dev

front-prepare:
	$(MAKE) -C frontend prepare