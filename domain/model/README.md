# ドメイン層について

## Golangでのドメインモデルの実装方法
### 実装方針
- ドメインルールを守ったオブジェクトであることを保証したい。
- Goの言語使用上外部からの変更（同一パッケージ内からの変更）を完全に防ぐことは難しいが、それでもなるべくドメインオブジェクトの不変性を保ちたい。

を実現する必要があるため、以下実装方法を守ること。

### 実装方法
インターフェイスを用いてエンティティや値オブジェクトを生成すればコンストラクタの使用を強制することができるため、<br>
ドメインルールを守ったオブジェクトであることを保証できる。

#### 1. エンティティの実装

```go
// インターフェイス
type UserBookInterface interface {
    // ゲッター
    UserId() UserIdInterface
    BookId() BookIdInterface
    Status() StatusInterface
    Memo() MemoInterface

    // セッター
    ChangeMemo(value *string) error
}

// 実装側の構造体
// => 先頭を小文字にすることで初期化できないようにする
type userBook struct {
    userId    UserIdInterface
    bookId    BookIdInterface
    status    StatusInterface
    memo      MemoInterface
}

// コンストラクタで初期化する
func NewUserBook(
    userId int,
    bookId int,
    status int,
    memo *string, // nilを許可したい場合ポインタ型にする
) (UserBookInterface, error) {
    s, err := NewStatus(status)
    if err != nil {
        return &userBook{}, err
    }

    m, err := NewMemo(memo)

    if err != nil {
        return &userBook{}, err
    }

    return &userBook{
        userId: NewUserId(userId),
        bookId: NewBookId(bookId),
        status: s,
        memo:   m,
    }, nil
}

// インターフェイスを満たすためのメソッド
// => Goではインターフェースの中にある同じ名前のメソッドを
//    全て定義すれば、自動的にインターフェイスを実装したことになる。
func (ub *userBook) UserId() UserIdInterface {
    return ub.userId
}

func (ub *userBook) ChangeMemo(value *string) error {
    return ub.memo.changeMemo(value)
}
```

#### 2. 値オブジェクトの実装
エンティティの実装と基本的に同じ。

- nil 許可していない場合
```go
package userbook

type UserIdInterface interface {
    Value() int
}

type userId struct {
    value int
}

func NewUserId(value int) UserIdInterface {
    return &userId{value}
}

func (ui *userId) Value() int {
    return ui.value
}
```

- nil 許可している場合
```go
package userbook

import (
    "fmt"
    "unicode/utf8"
)

// メモの最大文字数
const maxCount = 255

type MemoInterface interface {
    Value() *string
    changeMemo(value *string) error
}

// memo : 本のメモ
type memo struct {
    value *string
}

func NewMemo(value *string) (MemoInterface, error) {
    if err := validate(value); err != nil {
        return nil, err
    }

    // &でポインタ型を生成
    return &memo{value}, nil
}

func (m *memo) Value() *string {
    return m.value
}

func (m *memo) changeMemo(value *string) error {
    if err := validate(value); err != nil {
        return err
    }

    m.value = value
    return nil
}

// バリデーション
func validate(value *string) error {
    if value != nil {
        memoCount := utf8.RuneCountInString(*value)

        if memoCount > maxCount {
            return fmt.Errorf("メモは255文字以下で入力ください。")
        }
    }

    return nil
}
```

## 関連URL
- [【Go】Goでドメインオブジェクトをどのようにして生成すべきか](https://ryota21silva.hatenablog.com/entry/2023/03/21/173934)
- [インターフェイスを用いることで、ドメインオブジェクトを生成する際にコンストラクタの使用を強制する #52](https://github.com/ryota1116/stacked_books/pull/52)
