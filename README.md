## メモ
- バックエンド: https://ongakconnbe-qy2hw5cluq-an.a.run.app/dev/tracktag
- フロントエンド: https://ongakuconnection.com/


## cloud run に デプロイしたいとき 
- docker build -t ongk .
- (デバッグ docker run -p 8000:8080 ongk  (docker os 上の 8080 ポートを ローカルマシンの 8000 ポートへマウント))
- docker tag ongk gcr.io/ongakuconnection/ongkconn:latest
- docker push gcr.io/ongakuconnection/ongkconn:latest

## cloud sql に対して cloud run から接続したいとき
### DSN をこれに設定: 
  - DSN := "root:@unix(/cloudsql/ongakuconnection:asia-northeast1:ongkdb)/tomozoudb"
  - DSN := "root:@unix(/cloudsql/ongakuconnection:asia-northeast1:ongkdb)/tomozoudb?charset=utf8&parseTime=True"
  - parseTime を 忘れるとうまく動かない(DB へ保存できてなかったりする)
  - cloud run のデプロイ時の設定で、環境変数を持たせることもできる
  - やり直したいとき (drop database tomozoudb) 

## デプロイするとき .git が邪魔な場合
  - rm -rf .git/


## 変更したい時
- バックエンド: setting/const.go と main.go のDSN をローカル用に変更
- フロント: axios　の basURL を変更
- 作業環境: master ブランチにpush or pull request で デプロイされちゃうからそれ以外で
