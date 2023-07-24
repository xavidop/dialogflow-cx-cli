#!/bin/sh
set -e

SED="sed"
if which gsed >/dev/null 2>&1; then
	SED="gsed"
fi

cp -rf CONTRIBUTING.md docs/docs/community/contributing.md
cp -rf USERS.md docs/docs/community/users.md
cp -rf SECURITY.md docs/docs/security.md

rm -rf docs/docs/cmd/*.md
go run . docs
rm -rf docs/docs/static/*.json
go run . schema -f ./docs/docs/static/
"$SED" \
	-i'' \
	-e 's/SEE ALSO/See also/g' \
	-e 's/^## /# /g' \
	-e 's/^### /## /g' \
	-e 's/^#### /### /g' \
	-e 's/^##### /#### /g' \
	./docs/docs/cmd/*.md