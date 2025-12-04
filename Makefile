YEAR?=2025
DAY?=0
padded_day=$(shell printf '%02d' $(DAY))
GO_VERSION?=1.25.5

new_day:
	mkdir -p $(YEAR)
	mkdir -p $(YEAR)/$(padded_day)
	cp -R template/* $(YEAR)/$(padded_day)/
	sed -i 's/template/$(YEAR)\/$(padded_day)/g' $(YEAR)/$(padded_day)/cmd/part_1/main.go
	sed -i 's/template/$(YEAR)\/$(padded_day)/g' $(YEAR)/$(padded_day)/cmd/part_2/main.go
	git config --get remote.origin.url | sed 's/^.*\(github.*\).git/module \1\/$(YEAR)\/$(padded_day)\n\ngo $(GO_VERSION)\n/' > $(YEAR)/$(padded_day)/go.mod
	curl https://adventofcode.com/$(YEAR)/day/$(DAY)/input -H "Cookie: $(shell cat cookie)" > $(YEAR)/$(padded_day)/input.txt
