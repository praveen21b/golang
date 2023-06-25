# Golang for Beginner

* Running the GO inside the container <br/>

    ```
    docker build -t dev-go --target dev .
    ```

* Run container in interactive mode <br/>
    ```
    docker run -it -v $(pwd):/work dev-go bash
    ```

# Code Structure <br/>

* Module path 
    - Check [here](https://go.dev/doc/tutorial/create-module), a guide to create the `go module/s`
    - lets our module path for this project is `github.com/praveen21b/golang/introduction/app`
        Now go inside the container and create a folder named `app` and change working directory to app, then the below commad to create a module.
        ```
        go mod init github.com/praveen21b/golang/introduction/app
        ```
* Create a file named `app.go` under `app` and this will be out main entrypoint to the project.
* Copy hello-world program [here](https://gobyexample.com/hello-world) and paste it inside the `app.go'
* To run code 
    ```
    go run app.go
    ```

### Building a Binary
```
go build
```
A binary file called `app` gets created. To run the binary run the below command.
```
./app
```

### Running code from everywhere like CLI/kubectl
This helps if we want to go to commandline interface
```
go install github.com/praveen21b/golang/introduction/app

# To run the code
app
```