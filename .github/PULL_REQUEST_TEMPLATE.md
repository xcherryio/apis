## Why and Why

[Explain why you are making this pull request and what problem it solves.]

## Checklist before merge
[ ] Run the make command to update the generated code
[ ] Bump the version for Typescript and Python APIs, if needed to be used in Python/Typescript SDKs. Publish actions will fail if version is not updated.
  [ ] Typescript API library version is defined in the top of ts-api/package.json
  [ ] Python API library version is defined in Makefile `api-code-gen-py` target, parameter `packageVersion`
[ ] Make sure the openapi generator is not getting downgraded
[ ] If it is a breaking changes, it's recommended to merge in together with SDK+server. And set the title with "[Breaking changes]"