build: clean
	cd cmd && go build -o main

clean:
	cd cmd && rm -rf main
