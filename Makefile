run:
	gotest -v main_test.go

test:
	gotest -v day_$(day)_test.go


all:
	for i in {01..15}; do \
		make test day=$$(printf '%02d\n' $$i); \
	done
