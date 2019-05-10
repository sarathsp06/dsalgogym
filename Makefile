test:
	go	test	-race	-coverprofile=coverage.txt	-covermode=atomic	./...	\
	
	@NODE_ENV=test	./node_modules/.bin/mocha	\

.PHONY:	test