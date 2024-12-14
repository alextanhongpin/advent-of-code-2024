run:
	gotest -v main_test.go

test:
	gotest -v day_$(day)_test.go


all:
	for i in {00..13}; do \
		make test day=$$(printf '%02d\n' $$i); \
	done
