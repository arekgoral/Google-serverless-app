# A serverless app using Google Cloud Run service dynamically showing its own version number
An example  webserver written in golang that dynamically returns last commit's tag as a version number, last commit's SHA and description of the application.
The code in test2 folder builds and pushes image of the serverless app to the google Cloud Registry (GCR).
The image can be then used to deploy a serverless app using google Cloud Run service.
