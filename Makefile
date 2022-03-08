.PHONY: smoketest fmt-check

smoketest:
	tinygo build -o /tmp/app.hex -target wioterminal -size short ./examples/slideshow
	go build -o /tmp/slideshow ./examples/slideshow

fmt-check:
	@unformatted=$$(gofmt -l `find . -name "*.go"`); [ -z "$$unformatted" ] && exit 0; echo "Unformatted:"; for fn in $$unformatted; do echo "  $$fn"; done; exit 1
