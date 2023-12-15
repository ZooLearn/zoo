.PHONY: gen-ent
gen-ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./ent/template/*.tmpl" ./ent/schema --feature sql/execquery,intercept

.PHONY: migrate-up
migrate-up:

.PHONY: migrate-down
migrate-down:

.PHONY: test
test:
	go test ./...

.PHONY: gen-domain-test
gen-domain-test:
	mockery --dir=domain --output=domain/mocks --outpkg=mocks --all