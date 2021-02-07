# stacked_books(サービス名考え中)

## サービス概要

読書の積み上げによって達成感を得たい人に<br>
**読書量を「本の厚さ(cm, m)」で可視化**することによって達成感を与える<br>
読書量記録サービス

## マーケット
- 読書量を記録することで、達成感を得たい人
  - 紙だけでなく電子書籍で読書をする人も想定

## 登場人物
- エンドユーザー
  - 読書量を記録して達成感を得たい人
- 管理者
  - エンドユーザーの情報を管理する人

## ユーザーの課題
日々読書や勉強に励んでいるが、
電子書籍を使用したときに、どれくらいボリュームのある本を読んだのか目に見えないため、あまり達成感を得られない。<br>
また、読了した「本の冊数」を記録しても、本によってページ数は異なるため、どれくらいの読書量を積み重ねてきたのかピンと来ない。

## 解決方法
読書量を「本の厚さ(cm, m)」という単位で可視化することで、達成感を与える。

## 未来


## なぜ作るのか

### その人の本当の読書量を可視化したいという想いから生まれた

机の上に積み重ねられた本。ギッシリと本が詰まった本棚。それを見たときにこのサービスのアイデアが生まれました。<br>

私は読書や書籍を使った勉強が好きで、ふと机の上に積み重ねられた本や本棚を見た時に「今までこれだけの量の本を読んできたのか」と達成感を得ることがよくあります。<br>
一方で、電子書籍を使うと「実際に本を手に取ることができず、どれくらいボリュームのある一冊を読み終えたのか実感しづらい」「積み重ねられた本を見た時の達成感を得られない」ということに気付きました。
また、既存の読書管理サービスは読了数を記録できるものの、100ページの本も300ページの本も「同じ1冊」として括られてしまうため、読書量を記録するという観点で、少しばかり違和感を覚えていました。<br>

そこで、読書量を本の厚さ(cm, m)で記録すれば、**「紙の本も電子書籍も関係なく、その人の本当の読書量を可視化できるのではないか」**と思い、このサービスを作りました。

## プロダクト
あなたが読んだ本を積み重ねると何cm、何mになる？<br>
読書量を本の厚さという単位で記録するサービス

## サービスページでは、どのように読書量が記録されていくのか

### 積み上げられた本がスカイツリーの高さを超える！？
サービスページでは、アナタが積み上げた本が、奈良の大仏、マンション、スカイツリーの高さを超え、いずれは宇宙にまで到達するような世界観を提供したいと思っています。<br>
一方、そのような世界観だけでは、どれくらいの読書量を積み上げてきたのかイメージが付かないため、実際の高さの記録はもちろん、一定の高さまで本を積み上げると「自動販売機(183cm)...バスケットゴール(3m)...信号機(5m)の高さを超えました！」といった表示もしてくれる仕様になっています。

### 他のユーザーと比較できる
積み上げた本の高さを他のユーザーとランキング形式で比較できるようにもなっています。<br>
また、本のカテゴリごとに積み上げ量を表示することも可能です。<br>

### 積み上げ記録はどうやって算出されるの？
例えば300ページの書籍を読了した場合、その本をアプリ上に登録することで、<br>
300ページ x 0.1mm/1ページ(※参考値) = 3cm<br>
という計算から、読書量が3cm加算されるような設定となっています。

## 画面遷移図
https://www.figma.com/file/OTN0BtC3RrgPhMgW9wILm2/Stacked_Books?node-id=0%3A1

## ER図
