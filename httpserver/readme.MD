1、容器运行
docker pull firemancoder/httpserver:xwjv1.0

docker run -itd -p 8181:80 firemancoder/httpserver:xwjv1.0

curl http://192.168.80.50:8181/header


2、通过nsenter查看ip
获取进程ip：
docker inspect -f {{.State.Pid}} d5cbdb70bb54
进入容器进程的network空间并查看
nsenter -n -t pid ip addr