## 介绍
该项目使用Golang进行构建，具体参见：https://mlog.club

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=server *.go


scp -r server root@194.163.147.93:/opt/bbs-go/server/.

[//]: # (pm2 start npm --name "bbs-go-site" -- run start)

pm2 start server --name "bbs-go-server"


# 启动进程/应用
pm2 start bin/www

# 重命名进程/应用
pm2 start app.js --name wb123、

# 添加进程/应用
pm2 start bin/www

# 结束进程/应用
pm2 stop www

# 结束所有进程/应用
pm2 stop all

# 删除进程/应用 pm2
pm2 delete www

# 删除所有进程/应用
pm2 delete all

# 列出所有进程/应用
pm2 list

# 查看某个进程/应用具体情况
pm2 describe www

# 查看进程/应用的资源消耗情况
pm2 monit

# 查看pm2的日志
pm2 logs 序号/名称

# 若要查看某个进程/应用的日志,使用
pm2 logs www

# 重新启动进程/应用
pm2 restart www

# 重新启动所有进程/应用
pm2 restart all

