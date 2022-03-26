FROM php:8.1-fpm-alpine

# Composerのインストール(マルチステージビルド)
COPY --from=composer:2.0 /usr/bin/composer /usr/bin/composer

# イメージのビルド時に使用する一時的な変数を設定
ARG RUNTIME_PACKAGES='nodejs tzdata postgresql-dev postgresql git'

# 環境変数を設定
ENV LANG=C.UTF-8 \
    TZ=Asia/Tokyo \
    COMPOSER_ALLOW_SUPERUSER=1

# パッケージのインストール
RUN apk update && \
    apk add --no-cache ${RUNTIME_PACKAGES} && \
    docker-php-ext-install pgsql pdo_pgsql

# 設定ファイルをコンテナにコピー
COPY ./php-fpm.d/zz-docker.conf /usr/local/etc/php-fpm.d/zz-docker.conf
COPY ./php.ini /usr/local/etc/php/php.ini

COPY ./src /var/www/html

WORKDIR /var/www/html
