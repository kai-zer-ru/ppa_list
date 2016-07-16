# Repository Manager ppa_list

This application is designed to store and add a list of repositories and applications that are needed in a single repository for later use in a new installation (reinstalling) Ubuntu.

## Assembly and installation

To collect this service and install it into the system, you need the following:

* GoLang 1.6+

Clone repository and collect:

```
git clone https://github.com/kaizer666/ppa_list.git
cd ppa_list
make
sudo make install
```

Then start the service:

```
sudo service ppa_list start
```

Go to the address [http://localhost:3333](http://localhost:3333) and rejoice.

The wrapper ## in the domain/subdomain through Nginx

In order to make the service run on domain/subdomain on the server Nginx, do the following:

```
cd /etc/nginx/conf.d
sudo nano ppa.your.domain.com.conf
```

And insert the following lines:

```
server {
    listen 80;
    server_name ppa.your.domain.com;

    location / {
        proxy_pass http: // localhost: 3333;
    }
}
```

## Setting up the service

All settings are stored in the file `/opt/ppalist/main.cfg`

## Example

Service work can be found here - [http://ppatest.kai-zer.ru](http://ppatest.kai-zer.ru)