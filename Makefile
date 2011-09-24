all: publishing
	./publishing

publishing: publishing.6
	6l -o publishing publishing.6

publishing.6: publishing.go pdf.go
	6g publishing.go pdf.go
