
.PHONY: all
all: castai/generated.go

castai/generated.go: openapi.json
	@echo Generating castai package...
	@oapi-codegen -package castai -generate types,client $< >$@

openapi.json: FORCE
	@echo -n Checking for newer $@ file...
	@wget -q -O $@.new https://api.cast.ai/v1/spec/openapi.json
	@if ! cmp -s $@ $@.new; then \
		echo " found!";\
		mv $@.new $@ ;\
	else \
		echo " up-to-date" ;\
	fi
	@$(RM) $@.new

FORCE: ;
