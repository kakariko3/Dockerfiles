FROM ruby:2.7.2

# 必要なパッケージのインストール
RUN apt update -qq && apt install -y \
  wget \
  curl \
  apt-transport-https

# yarnのインストール
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
  echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
  apt install -y yarn

# Node.jsのインストール
RUN curl -sL https://deb.nodesource.com/setup_13.x | bash - && \
  apt install -y nodejs

WORKDIR /myapp
COPY Gemfile /myapp/Gemfile
COPY Gemfile.lock /myapp/Gemfile.lock
RUN bundle install
COPY . /myapp

# コンテナが起動する度に実行されるスクリプトを追加
COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
EXPOSE 3000

# 実行したいコマンド
CMD ["rails", "server", "-b", "0.0.0.0"]
