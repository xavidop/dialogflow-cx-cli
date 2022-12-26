#!/bin/sh
set -e

SED="sed"
if which gsed >/dev/null 2>&1; then
	SED="gsed"
fi

cp -rf CONTRIBUTING.md docs/docs/contributing.md
cp -rf USERS.md docs/docs/users.md
cp -rf SECURITY.md docs/docs/security.md

rm -rf docs/docs/cmd/*.md
go run . docs
go run . schema -f ./docs/docs/static/schema.json
"$SED" \
	-i'' \
	-e 's/SEE ALSO/See also/g' \
	-e 's/^## /# /g' \
	-e 's/^### /## /g' \
	-e 's/^#### /### /g' \
	-e 's/^##### /#### /g' \
	./docs/docs/cmd/*.md