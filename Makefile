default: gen

gen: api-code-gen-go

api-code-gen-go: #generate/refresh go clent code for idl, do this after update the idl file
	rm -Rf ./goapi ; true
	openapi-generator generate -i api-schema/xdb.yaml -g go -o goapi/xdbapi/ -p packageName=xdbapi -p generateInterfaces=true -p isGoSubmodule=false --git-user-id xdblab --git-repo-id xdb-apis
	rm ./goapi/xdbapi/go.* ; rm -rf ./goapi/xdbapi/test; gofmt -s -w goapi; true
