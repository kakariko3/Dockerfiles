FROM php:7.4-apache

RUN apt-get update && \
  # PDO PostgreSQL 拡張
  apt-get install -y libpq-dev && \
  docker-php-ext-install pgsql pdo_pgsql

# COPY php.ini /usr/local/etc/php/
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer
