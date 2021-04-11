# 🍩 Spinning Donut

Spinning donut using ASCII characteres 

## Motivation

I saw this [video](https://www.youtube.com/watch?v=DEqXNfs_HhY) from Lex Fridman and try replicate in Go.
The result is not like in the video but I think is pretty similar. 


## How to run 


Use the ```build ``` script to build and run the code.

If you dont have go in your machine, and have docker, use to build

```
docker run -v $(pwd):/donut -w /donut golang:1.15-alpine sh -c "go build ."
```

