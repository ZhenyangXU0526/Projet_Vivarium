package main

//go:generate file2byteslice -package=images -input=./images/spider.png -output=./images/spider.go -var=Spider_png
//go:generate file2byteslice -package=images -input=./images/runner.png -output=./images/runner.go -var=Runner_png
//go:generate file2byteslice -package=images -input=./images/tiles.png -output=./images/tiles.go -var=Tiles_png
//go:generate file2byteslice -package=images -input=./images/snail.png -output=./images/snail.go -var=Snail_png