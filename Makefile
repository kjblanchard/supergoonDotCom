#
# Makefile is used to handle local development easily, without having to manage remembering multiple commands.
# Handles varios methods of building, and running with different profiles.
#
.PHONY: build publish stop run package frontend bindir api apackage abuild

all: run

bindir:
	@$(foreach PACKAGE,$(PACKAGES), \
		cd ./$(PACKAGE) && mkdir -p $(BINARY_FOLDER_NAME) && cd - ; \
	)
clean:
	@rm -rf `find . -type d -name bin`

run:
	@cd frontend/src && go run .
