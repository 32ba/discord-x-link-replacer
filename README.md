# discord-x-link-replacer

English version is below.

X(旧Twitter)のリンクを含んだメッセージに対して，Discord上で埋め込み表示可能なURLをリプライで返すBotです．

## 使い方

1. Discord Developer PortalからBotを作成し，トークンを取得してください．

1. [compose.sample.yml](compose.sample.yml)を参考に，`compose.yml`を作成してください．

   基本的には，`DISCORD_TOKEN`に取得したトークンを設定するだけです．

1. `docker compose up -d`を実行してください．

 　ログにエラーが出力されず，Botがオンラインになれば成功です．

## ライセンス

このプロジェクトは，MIT Licenseの下で公開されています．詳細は，[LICENSE](LICENSE)を参照してください．

---

This is a bot that replies to messages containing X (formally Twitter) links with a URL that can be embedded on Discord.

## Usage

1. Create a bot from the Discord Developer Portal and get the token.

1. Create `compose.yml` based on [compose.sample.yml](compose.sample.yml).

   Basically, just set the token you got to `DISCORD_TOKEN`.

1. Run `docker compose up -d`.

    If no errors are output to the log and the bot is online, it was successful.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.