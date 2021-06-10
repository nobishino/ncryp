# ncryp

Learn cryptography with weak key and algorithms.

暗号学習のための弱い暗号ライブラリです。
## Install

### homebrew

- `brew tap nobishino/homebrew-tap`
- `brew install ncryp`

バージョンアップするときは

`brew upgrade ncryp`

を実行してください。

### Go

TBW
### Download binary 

TBW

## シナリオ

AliceはBobに共通鍵暗号でバイト列`0314159265`を送りたいです。しかしそのための共通鍵もネットワーク越しに送らなければなりません。そこでAliceとBobは公開鍵暗号であるRSA暗号を用いてまず共通鍵を送り、その後でその共通鍵を使って通信することにしました。

### BobがRSA鍵ペアを生成する

BobがRSA鍵ペアを生成します。

```
ncryp -keygen
83500100000000005F6C000000000000 83500100000000005345000000000000
```
### Bobが公開鍵を公開する

Bobは生成した鍵ペアのうち一方を公開鍵、もう一方を秘密鍵として使うことに決めます。
ここでは右の`83500100000000005345000000000000`を公開し、もう一方の`83500100000000005F6C000000000000`は秘密鍵として秘密にします。

公開鍵はAlice以外の人に見られても構いません。

### Aliceが共通鍵を生成して公開鍵で暗号化する

Aliceは通信で使いたい共通鍵`E39A`をBobの公開鍵で暗号化します。(共通鍵は2byteであるものと約束します)

```
ncryp -payload E39A -key 83500100000000005345000000000000 -rsa
EE93000000000000
```

そしてこの結果`EE93000000000000`をBobに送ります。これは盗聴されても構いません。

### Bobが秘密鍵を用いて復号する

```
ncryp -payload EE93000000000000 -key 83500100000000005F6C000000000000 -rsa
E39A000000000000
```

共通鍵は2byteと約束したので、得られる共通鍵は`E39A`です。

### Aliceが共通鍵で通信文を暗号化する

```
ncryp -payload 0314159265 -key E39A -naive
E08EF60886
```

### Bobが共通鍵で通信文を復号化する

```
ncryp -payload E08EF60886 -key E39A -naive
0314159265
```