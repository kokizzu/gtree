# gentree

## description
- cliです。
- マークダウンファイルの入力で、treeコマンドで出力するような結果が出力されます。
- マークダウンファイルは、testdataディレクトリ内のファイルを参考に作成してください。
    - ディレクトリ階層はインデントで表します。
    - インデントは以下のいずれかで統一してください。
        - タブ
        - 半角スペース２つ（`-ts`フラグ必須）
        - 半角スペース４つ（`-fs`フラグ必須）

## installation
```sh
go get github.com/ddddddO/gentree
```

## how to use

```sh
19:17:07 > cat testdata/sample1.md | gentree
a
├── vvv
│   └── jjj
├── kggg
│   ├── kkdd
│   └── tggg
├── edddd
│   └── orrr
└── gggg
```

or `gentree -f testdata/sample1.md`<br>
or `cat testdata/sample1.md | gentree -f -`

```sh
16:31:42 > cat testdata/sample2.md | gentree
k8s_resources
├── (Tier3)
│   └── (Tier2)
│       └── (Tier1)
│           └── (Tier0)
├── Deployments
│   └── ReplicaSet
│       └── Pod
│           └── containers
├── CronJob
│   └── Job
│       └── Pod
│           └── containers
├── (empty1)
│   └── DaemonSet
│       └── Pod
│           └── containers
└── (empty2)
    └── StatefulSet
        └── Pod
            └── containers
```

---
- Two spaces indent

```sh
01:15:25 > cat testdata/sample4.md | gentree -ts
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
```

- Four spaces indent

```sh
01:16:46 > cat testdata/sample5.md | gentree -fs
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
```

- Fixed bug!(2021/05/29)


```sh
00:16:59 > cat testdata/sample3.md | gentree
root
├── dddd
│   └── kkkkkkk
│       └── lllll
│           ├── ffff
│   │   │   ├── LLL
│   │   │   │   └── WWWWW
│   │   │   │       └── ZZZZZ
│           └── ppppp
│               └── KKK
│                   └── 1111111
│                       └── AAAAAAA
└── eee
```
↓
```sh
22:21:29 > cat testdata/sample3.md | gentree
root
├── dddd
│   └── kkkkkkk
│       └── lllll
│           ├── ffff
│           ├── LLL
│           │   └── WWWWW
│           │       └── ZZZZZ
│           └── ppppp
│               └── KKK
│                   └── 1111111
│                       └── AAAAAAA
└── eee
```