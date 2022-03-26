FROM nginx:1.20-alpine

# ユーザーの追加
RUN adduser -S www-data -u 82 \
  && echo "www-data ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
  && echo 'www-data:www-data' | chpasswd

# 設定ファイルをコンテナにコピー
COPY ./conf.d/default.conf /etc/nginx/conf.d/default.conf
COPY ./nginx.conf /etc/nginx/nginx.conf
