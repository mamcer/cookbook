# cookbook

What is cookbook? an experiment, an excuse to build a simple REST api  using golang, to play with vanilla javascript, with sqlite (and then migrated to mysql), to test how it can be hosted on a phone (using termux), a way to document recipes, all of them and maybe the most important was a way to have some fun and distract the mind amid so much uncertainty and at times fear during the moments of greatest restrictions during the pandemic

It is functional but still at a development level. With for example diferent hardcoded values.

## development db 

```bash
docker pull mysql:latest
docker run -p 3366:3306 --name cookbook -e MYSQL_ROOT_PASSWORD=root -d mysql:latest

docker exec -it cookbook mysql -uroot -p
create database cookbook;
```

create schema

```bash
mycli -h localhost -P 3366 -u root
use cookbook
\. schema.sql
```

(or) restore backup

```bash
# restore a previously created backup 
mysqldump -u root --password=[password] [db-name] > backup.sql
source backup.sql
```
## how to run

on repository root directory

    make build
    ./bin/main

> api is on port 5001 'curl http://localhost:5001/ping', web app on port 5000 'http://localhost:5000'

## configure as a service

    sudo vim /etc/systemd/system/cookbook.service

add the following content:

[Unit]
Description=Cookbook

[Service]
WorkingDirectory=/home/mario/bin/cookbook
ExecStart=/home/mario/bin/cookbook/main
Restart=always
RestartSec=10
User=mario
Environment=DISPLAY=:0
Environment="GIN_MODE=release"

[Install]
WantedBy=multi-user.target

then:

    sudo systemctl enable cookbook.service
    sudo systemctl start cookbook.service
    sudo systemctl status cookbook.service

view logs:

    journalctl -u cookbook -b
