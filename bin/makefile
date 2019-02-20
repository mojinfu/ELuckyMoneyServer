RM= rm -rf

all: clean release

release: ELuckyMoneyServer

clean:
	$(RM) ../log/* 
	$(RM)  ELuckyMoneyServer

ELuckyMoneyServer:
	go build -ldflags "-s -w" -o $@ ../src/main/*.go
