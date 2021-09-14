FROM php:7.4-apache

# COPY ./php.ini /usr/local/etc/php/

RUN apt-get update && \
  # PDO PostgreSQL 拡張
  apt-get install -y libpq-dev && \
  docker-php-ext-install pgsql pdo_pgsql
