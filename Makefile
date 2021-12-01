all: gotool #生成go二进制文件   makefile使用https://www.cnblogs.com/wang_yb/p/3990952.html
	@go build -v .
clean:  # 清理工作：删除二进制文件、删除 vim swp 文件
	rm -f go-api-example
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool: #格式化代码和源码静态检查
	gofmt -w .
	go vet . | grep -v vendor;true
help: # 打印 help 信息
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"

.PHONY: clean gotool ca help
