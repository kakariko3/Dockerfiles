FROM php:7.4-apache

ENV COMPOSER_ALLOW_SUPERUSER=1

COPY --from=composer:2.0 /usr/bin/composer /usr/bin/composer

RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - && \
  apt-get update && \
  apt-get install -y libpq-dev zip curl gnupg nodejs && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* && \
  a2enmod rewrite && \
  docker-php-ext-install pgsql pdo_pgsql && \
  npm install -g npm@latest && \
  npm install -g yarn

COPY ./php.ini /usr/local/etc/php/
COPY ./httpd.conf /etc/apache2/sites-available/000-default.conf

WORKDIR /var/www/html
