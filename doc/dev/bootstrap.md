# 项目开发环境搭建

详细的步骤请参考 [项目开发设定](./setup.md), 这里只给出基本的配置流程步骤。

## setup machine

### Mac 环境

- [Install Homebrew](https://brew.sh/)

### Linux 环境

都用 linux 了，你应该会的了

### Windows 环境

Windows 是啥？

## -2. 安装 Go

- [Downloads Go](https://golang.org/dl/)

一般情况下使用最新和次新的两个版本均可使用。

## -1. setup docker

### -1.1 Mac 环境

- [Install Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
- [macOS 安装 docker](https://yeasy.gitbooks.io/docker_practice/install/mac.html)

### -1.2 Linux 环境

- [Get Docker CE for Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
- [Docker - ArchWiki](https://wiki.archlinux.org/index.php/Docker)

安装好 docker 后，安装 `docker-compose`:

- [Install Docker Compose](https://docs.docker.com/compose/install/)

## 0. clone

```
git clone ${GIT_REPO_PATH} /path/to/workspace
```

## 1. 编译命令行工具

aegis-tool 提供了一些辅助工具

```
$ make binary-tools
$ ./bin/aegis-tools -h
```

## 2. 启动 dev 环境依赖
      
目前开发环境依赖使用 docker / docker-compose 管理，可以通过以下方式初始化：

```
$ cd dev
$ docker-compose up -d
# 等若干分钟...
$ docker-compose ps
         Name                       Command             State                             Ports                           
--------------------------------------------------------------------------------------------------------------------------
dev_clickhouse-server_1   /entrypoint.sh                Up      0.0.0.0:18123->8123/tcp, 0.0.0.0:19000->9000/tcp, 9009/tcp
dev_mysql_1               docker-entrypoint.sh mysqld   Up      0.0.0.0:13306->3306/tcp, 33060/tcp    

这个时候，可以在外部使用 13306 访问 mysql 服务，使用 18123/19000 访问 clickhouse 服务。
```

## 3. 初始化配置

```
$ cp config/aegis.yml.local config/aegis.yml
```

如果你本地配置和默认配置有不同，请根据需要进行编辑，本地开发配置请不要加入到版本控制。

## 4. 初始化表结构

```
#初始化表结构
$ cd /path/to/workspace/aegis
$ ./bin/aegis-tools migration:metadata up
$ ./bin/aegis-tools migration:warehouse up
```

## 5. 启动本地开发服务

```
$ go run cmd/engine/*.go

{"event":"modelManager.init.done","hostname":"helios.local","mode":"production","server":"engine","time":"2019-11-07T11:31:51.793331Z"}
{"event":"propertyManager.init.done","hostname":"helios.local","mode":"production","server":"engine","time":"2019-11-07T11:31:51.793383Z"}
{"event":"abstractionManager.init.done","hostname":"helios.local","mode":"production","server":"engine","time":"2019-11-07T11:31:51.793394Z"}
{"event":"activationManager.init.done","hostname":"helios.local","mode":"production","server":"engine","time":"2019-11-07T11:31:51.793401Z"}
```

## 6. 初始化本地测试配置


本地测试的时候会自动使用 `aegis_test.yml`:

```
$ cp config/aegis.yml.example config/aegis_test.yml
```

如果想区分两个环境的数据库使用，修改配置文件中相关配置即可。

## 7. 开发

```
$ $EDITOR .
```

## 8. 测试

```
$ make test
```
