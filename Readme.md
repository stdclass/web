web(1)
===

Create a vhost-configuration file under `/etc/apache2/sites-available`, activate the site with `a2enmod` and reload the apache server configuration with 1 command.

Usage
---

Simple:

    $ web domain.com /var/www/domain.com 

Domain with alias:

    $ web domain.com /var/www/domain.com www.domain.com


