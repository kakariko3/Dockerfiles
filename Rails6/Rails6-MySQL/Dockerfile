FROM ruby:2.7.1

# 必要なパッケージのインストール
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
  echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
  apt-get update -qq && \
  apt-get install -y nodejs yarn

WORKDIR /myapp
COPY Gemfile /myapp
COPY Gemfile.lock /myapp
RUN bundle install
COPY . /myapp

# コンテナが起動する度に実行されるスクリプトを追加
COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
EXPOSE 3000

# 起動時に実行するコマンド
CMD ["rails", "server", "-b", "0.0.0.0"]
