# vscode用GO Remote container init


## 環境構築

前提：Docker 、VSCodeインストール済みであること

- VSCodeにRemote Containerをインストールします。
    - 左のメニューから拡張機能のインストールを開き、「remote container」等で検索し、インストール。
- GitHubからプロジェクトをクローンします。
- クローンしたプロジェクトを、VSCodeで開きます。
- 左下のマークから、「Reopen in Container」を選択します。
    - すると自動的にdocker-compose.ymlを見に行きコンテナがビルドされ、立ち上がります。
- 下の方に「Dev Container:Go」と表示されていれば立ち上がっています。
- go.modで記載しているパケジーをインストールする

- これで環境構築は完了です。


