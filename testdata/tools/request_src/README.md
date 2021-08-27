# request.sh ユニット・テスト

[request.sh](../request.sh) が使う関数を外部ファイル化した [functions.sh](./functions.sh) のファイルと、その関数の<ruby>ユニット・テスト<rt>単体テスト</rt></ruby>置き場です。

ユニット・テストのフレームワークには [`shellspec`](https://github.com/shellspec/shellspec) を使っています。

## 実行方法

リポジトリのルートディレクトリから実行してください。（ルートディレクトリに、テストの実行に必要な設定ファイル（`.shellspec`）が設置されています）

```shellsession
$ shellspec
...
```

## 参考文献

- 基本的な使い方本は、以下の作者の日本語記事をご覧ください。
    - [ShellSpec - シェルスクリプト用のフル機能のBDDユニットテストフレームワーク](https://qiita.com/ko1nksm/items/2f01ff4f50e957ebf1de) @ Qiita
- 早見表と使用例
    - [References](https://github.com/shellspec/shellspec/blob/master/docs/references.md) | 英語 @ GitHub
