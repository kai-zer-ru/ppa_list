# Менеджер репозиториев ppa_list

Данное приложение предназначено для хранения и добавления списка репозиториев и необходимых приложений в одно хранилище для дальнейшего использования при новой установке (переустановке) Ubuntu.

## Сборка и установка

Что бы собрать данный сервис и установить его в систему, необходимо следующее:

* GoLang 1.6+

Клонируем репозиторий и собираем:

```
git clone https://github.com/kaizer666/ppa_list.git
cd ppa_list
make
sudo make install
```

Затем запускаем сервис:

```
sudo service ppa_list start
```

Переходим по адресу [http://localhost:3333](http://localhost:3333) и радуемся.

## Обёртка в домен/поддомен через Nginx

Для того, что бы заставить сервис работать на домене/поддомене на сервере с Nginx, сделайте следующее:

```
cd /etc/nginx/conf.d
sudo nano ppa.your.domain.com.conf
```

И вставьте следующие строки:

```
server {
    listen 80;
    server_name ppa.your.domain.com;

    location / {
        proxy_pass http://localhost:3333;
    }
}
```

## Настройка сервиса

Все настройки хранятся в файле `/opt/ppalist/main.cfg`

## Пример

Работу сервиса можно посмотреть тут - [http://ppatest.kai-zer.ru](http://ppatest.kai-zer.ru)