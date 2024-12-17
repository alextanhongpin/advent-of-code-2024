main:
	gotest -v day_16_unsolved_p2_test.go
	#gotest -v main_test.go

run:
	gotest -v day_$(day)_test.go


all:
	for i in {01..15}; do \
		make run day=$$(printf '%02d\n' $$i); \
	done
