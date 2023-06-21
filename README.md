# ANZ-Test
ANZ-Test is a golang webserver that dynamically returns last commit's tag as a version, last commit's SHA and description of the application.
The code in test folder builds and pushes image of the serverless app to the google cloud registry (GCR)
The last commit needs to be git tagged. E.G 'git tag 1.1'
Decription is hard-coded. Version and SHA are injected during the build time by using the argument '-ldflags'


## Requirements
To run this web service locally, you will need following:
- [`docker`]
- [`go@1.13+`]
- [`git`]

execute: 
git tag [version]
build.sh to build the docker container
Access the web service via 
http://127.0.0.1:8080/version
