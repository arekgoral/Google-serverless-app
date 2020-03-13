# ANZ-Test - Task 1
This Dockerfile is refactored version of original Dockerfile from Anz-test 1. It builds a GO container and has fixed issue that was causing compilation error.
It then builds the same web service as in Task 2.
To test locally run: sudo docker build -t testgocontainer .