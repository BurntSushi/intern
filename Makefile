tag:
	gotags *.go > TAGS

fmt:
	gofmt -w *.go

push:
	git push origin master
	git push github master

