.PHONY: back-build back-run back-test back-fmt back-lint

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