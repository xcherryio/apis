default: gen tests

gen: api-code-gen-go api-code-gen-ts api-code-gen-py

api-code-gen-go: #generate/refresh go clent code for idl, do this after update the idl file
	rm -Rf ./goapi ; true
	openapi-generator generate -i api-schema/xcherry.yaml -g go -o goapi/xcapi/ -p packageName=xcapi -p generateInterfaces=true -p isGoSubmodule=false --git-user-id xcherryio --git-repo-id apis
	rm ./goapi/xcapi/go.* ; rm -rf ./goapi/xcapi/test; gofmt -s -w goapi; true

api-code-gen-ts: #generate/refresh typescript apis
	rm -Rf ./ts-api/src/api-gen ; true
	openapi-generator generate -i ./api-schema/xcherry.yaml -g typescript-axios -o ./ts-api/src/api-gen --git-user-id xcherryio --git-repo-id apis

api-code-gen-py: #generate/refresh python apis
	rm -Rf ./pyapi/* ; true
	openapi-generator generate -i ./api-schema/xcherry.yaml -g python -o ./pyapi -p packageVersion=0.0.3 -p packageName=xcherryapi --git-user-id xcherryio --git-repo-id apis

tests:
	$Q go test -v ./gotests 
