# board_repository

## どのようなアプリか
### 概要
簡潔に言うと、「ホワイトボード×マッチングアプリのデザイン」よって、リードジェネレーションができるアプリです。
### 背景
スタートアップ企業でのtoB営業インターンで感じた課題を解決するために作成しました。

その課題とは、スタートアップ企業の営業が活用できるサービスが、極めて限定的であることです。

詳述すると、

・予算の都合上、広告を打つことが難しい。

・メルアポ・テレアポでは、リードナーチャリングができない。

・SNSのDMによるセールスでは、リードの質を担保できない（不確実性が伴う）。

上記のような、営業を取り巻く環境を、解決し得るサービスがあったらいいなと思い、このアプリを作成しました。

## 動作確認

### ①【1人目（佐藤）のユーザー作成、ログイン、プロフィール変更】
![1番目_佐藤ログイン](https://user-images.githubusercontent.com/75367572/113115343-4ef5d880-9247-11eb-85af-84d55bbc7f07.GIF)

### ②【1人目（佐藤）のユーザーでコンテンツの作成、削除、編集】

この時、自身が作成したコンテンツは、左側の大画面に表示されないことを確認できます。
![2番目_佐藤board作成](https://user-images.githubusercontent.com/75367572/113115403-5cab5e00-9247-11eb-9abb-897210e1fa59.GIF)

### ③【2人目（高橋）のユーザーでコンテンツの作成】
![3番目_高橋board作成](https://user-images.githubusercontent.com/75367572/113115452-6a60e380-9247-11eb-98b2-282c70b6f88d.GIF)

### ④【3人目のユーザー（伊藤）でログイン、コンテンツの閲覧 & 1人目と2人目のリードの閲覧】

下記が確認できます。

・左側の大画面で、先ほど作成されたコンテンツが表示されます。

・左下の両脇のボタンを押すと、表示されるコンテンツが入れ替わります。

・sendボタンを押したコンテンツの作成者に対しては、伊藤がリードとして表示されます。

・ゴミ箱ボタンを押したコンテンツの作成者に対しては、伊藤はリードとして表示されません。
![4番目_伊藤動作確認](https://user-images.githubusercontent.com/75367572/113116160-24f0e600-9248-11eb-8bd9-33b40e12128e.GIF)


## こだわり

ホワイトボードのランダム表示にこだわりました。

ホワイトボードの閲覧者にとって、興味のないホワイトボードは、興味ないものとして今後表示させないこと。

ホワイトボードの閲覧者にとって、興味のあるホワイトボードでは、リードとして自分自身のプロフィールを、ホワイトボードの作成者側に送ること。
そして、その作成者のホワイトボードは今後表示しないこと。

上記のロジックを、バックエンド、フロントエンド両側から実装しました。

