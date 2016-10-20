
VAR1 := "var 1"
VAR2 ?= "var 2"

check:
	echo $(VAR1)
	echo $(VAR2)

test: 
	echo $(MAKE)
	echo $$(which make)
	echo $$USER
	@if [ $$USER != vagrant ]; then \
		echo $$USER ; \
	fi
