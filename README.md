## General Info
This project is web application to translate from Bahasa Indonesia (ID) to English (EN). 

![image](https://user-images.githubusercontent.com/38047246/123521140-00899600-d6df-11eb-84e8-6d788750f175.png)

## Technologies
Project is created with:
* Go version: go1.16.3
* [Gtranslate](https://github.com/bregydoc/gtranslate)
* [Testify](https://github.com/stretchr/testify)
## Setup in local
To run this project, install it locally:

``` bash
# clone the repo
$ git clone https://github.com/anonymousliem/transhift 

# go into app's directory
$ cd transhift

# ignore it if you use linux
# for windows machine ->  in main.go change
from Addr: ":8080" to Addr: "localhost:8080"

# to run application
go run main.go

# open browser and paste this example
http://localhost:8080/transhift?q=CHANGE_IT_WITH_YOUR_WORD
```

## What's included

```
├──  main.go         #main code           
|__  main_test.go    #test function 	
|	
|__  Dockerfile      #build to docker
```

## Run Application With Docker
```
docker run -p 8080:8080 anonymousliem/transhift:latest
```