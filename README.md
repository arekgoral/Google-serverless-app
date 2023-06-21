# ANZ-Test
An example  webserver written in golang that dynamically returns last commit's tag as a version number, last commit's SHA and description of the application.
The code in test2 folder builds and pushes image of the serverless app to the google Cloud Registry (GCR).
The image can be then used to deploy a serverless app using google Cloud Run service.



## Local Testing Requirements
The last commit needs to be git tagged. E.G 'git tag 1.1'
The decription returned along with the version is hard-coded. Version and SHA are injected during the build time by using the argument '-ldflags'
To run this web service locally, you will need following:
- [`docker`]
- [`go@1.13+`]
- [`git`]

execute: 
git tag [version]
build.sh to build the docker container
Access the web service via 
http://127.0.0.1:8080/version
