# DiscordRoleGranter
## 概要
Nginxなり、Apacheなりの認証システムをくぐれる人だけロールを付与できるようにしたかったため作成したCGI。汎用性は不明。

## 目次
<!-- TOC -->

- [DiscordRoleGranter](#discordrolegranter)
    - [概要](#概要)
    - [目次](#目次)
    - [使い方](#使い方)
    - [設定](#設定)

<!-- /TOC -->

## 使い方

```
$ go build -o index.up main/*
```

## 設定
index.upと同階層に次の`settings.json`を用意。

```json
{
    "Discord": {
        "Token": "DiscordBotToken",
        "GuildID": "******************",
        "RoleID": "******************"
    }
}
```



