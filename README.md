# cookbook

What is cookbook? an experiment, an excuse to build a simple REST api  using golang, to play with vanilla javascript, with sqlite (and then migrated to mysql), to test how it will run from a phone, a way to document recipes, all of them and maybe the most important was a way to have some fun and distract the mind amid so much uncertainty and at times fear during the moments of greatest restrictions during the pandemic

It is functional but still at a development level. With for example diferent hardcoded values. Originally used directly from my Android One phone with termux (2020) 

## mysql 

    docker pull mysql:latest
    docker run -p 3306:3306 --name cookbook -e MYSQL_ROOT_PASSWORD=root -d mysql:latest

    docker exec -it cookbook mysql -uroot -p
    create database cookbook;

## create schema

    mycli -h localhost -u root
    use cookbook
    \. schema.sql

## restore backup

    # restore a previously created backup mysqldump -u root --password=[password] [db-name] > backup.sql
    source backup.sql

## how to run

    go run main.go

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

[Install]
WantedBy=multi-user.target

then:

    sudo systemctl enable cookbook.service
    sudo systemctl start cookbook.service
    sudo systemctl status cookbook.service

view logs:

    journalctl -u cookbook -b