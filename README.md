# Golang for Beginner

## Running the GO inside the container <br/>

```
docker build -t dev-go --target dev .
```

## Run container in interactive mode <br/>
```
docker run -it -v $(pwd):/work dev-go bash
```