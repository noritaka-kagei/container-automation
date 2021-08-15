# keronori/container-automation
This project has the following behavior:
1. Compile the Java source code (Maven project).
2. It makes the compiled code into a distribution (.jar and/or .war).
3. Generates an executable container image of the distribution.
4. Push the generated container image to an external repository(ex. Docker Hub).

### Check your environment for execute this project
$ go run init.go

### (Optional) Copy your Maven project (source code) to an appropriate directory
* In this repository, a sample project [sample-app/] is provided.

### Execute for creating container image and pushing external repository.
$ go run execute.go