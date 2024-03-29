# 20K1105-Exercise3-LineBot
リクエスト集中講義「プログラミング言語Go入門」の演習課題3 です。

# 実行方法
![lineQR](https://user-images.githubusercontent.com/78438738/109186576-69660d80-77d4-11eb-95d9-63bb1e0dc96b.png)
1. Line にて、上記の QR コードより「学習塾用通知Bot」を友達追加します。
2. 「鈴宮 花子」と入力します。(名字と名前の間は、半角スペース 1 つ空ける)
3. 鈴宮 花子の学習時間の制限である 1 分経過後、通知が Line にて届きます。
4. 他の生徒で試したい場合は、以下の名前を入力します。
  * 鈴宮 太郎 --> 2 分
  * 鈴宮 次郎 --> 5 分

# 参考文献等
* golangとHerokuとLinebotでオウム返しするbotを作成する。  
 ( heroku との連携や LINE API の使用法の参考にさせていただきました。) (https://satoki252595.com/2020/08/10/linebot_01/#golang)
* echo_bot  
(golang の LineBot 専用のライブラリ`line-bot-sdk-go`を用いた LineBot の一例。上記のサイトでも参考にしている LineBot です。)  
(https://github.com/line/line-bot-sdk-go/blob/master/examples/echo_bot/server.go)
* 【Golang】Goで決まった時間に論語を配信(Push)するLINE Botを作った  
( Push を学ぶ際に参考にさせていただきました。) (https://tanaken.me/posts/190222/)
