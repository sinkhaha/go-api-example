SHELL := /bin/bash
BASEDIR = $(shell pwd)

# 编译时传入版本信息参数 运行时带上-v即可查看
versionDir = "apiserver/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

# -w 为去掉调试信息（无法使用 gdb 调试），这样可以使编译后的二进制文件更小
ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"
all: gotool # swag init为生成swagger文档。go build生成go二进制文件，-ldflags传入版本等信息 ，makefile使用https://www.cnblogs.com/wang_yb/p/3990952.html  
	swag init
	@go build -v -ldflags ${ldflags} .
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
