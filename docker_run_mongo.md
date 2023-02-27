1. 下载mongo 的 image 镜像文件
docker pull mongo
2. 测试镜像是否可用
docker run  \ --name mongo \ -p 27017:27017  \ -d mongo
3. 选择本地目录挂载mongo数据（永久保存在本地，否则镜像关闭数据就会丢失）
/Users/jxc_year/data/docker/mongodb
  a. 目录下细分conf和data/db两个文件夹
4. 之后的管理就可以直接使用docker的桌面版管理即可