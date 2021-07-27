# gtree

[![ci](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml) [![codecov](https://codecov.io/gh/ddddddO/gtree/branch/master/graph/badge.svg?token=JLGSLF33RH)](https://codecov.io/gh/ddddddO/gtree) [![GitHub release](https://img.shields.io/github/release/ddddddO/gtree.svg)](https://github.com/ddddddO/gtree/releases)

(markdown || program) to tree.

```
├── CLI and Package(1).
│   ├── Given a markdown file or format, the result of linux tree command is printed.
│   ├── `gtree` does not temporarily create directories or files.
│   └── Create markdown file by referring to the file in the `testdata/` directory.
│       ├── Hierarchy is represented by indentation.
│       └── Indentation should be unified by one of the following.
│           ├── Tab
│           ├── Two half-width spaces（required: `-ts`）
│           └── Four half-width spaces（required: `-fs`）
└── Package(2).
    ├── You can also generate a tree programmatically.
    └── Markdown is irrelevant.
```
(generated by `cat testdata/sample0.md | gtree -fs`)


---

## As CLI

### Installation
```sh
go get github.com/ddddddO/gtree/cmd/gtree
```

or, download from [here](https://github.com/ddddddO/gtree/releases).


### Usage

```sh
20:25:28 > gtree -ts << EOS
> - a
>   - vvv
>     - jjj
>   - kggg
>     - kkdd
>     - tggg
>   - edddd
>     - orrr
>   - gggg
> EOS
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

#### OR
```
├── gtree -f testdata/sample1.md
├── cat testdata/sample1.md | gtree -f -
└── cat testdata/sample1.md | gtree
```

---

- Usage other than representing a directory.

```sh
16:31:42 > cat testdata/sample2.md | gtree
k8s_resources
├── (Tier3)
│   └── (Tier2)
│       └── (Tier1)
│           └── (Tier0)
├── Deployment
│   └── ReplicaSet
│       └── Pod
│           └── container(s)
├── CronJob
│   └── Job
│       └── Pod
│           └── container(s)
├── (empty)
│   └── DaemonSet
│       └── Pod
│           └── container(s)
└── (empty)
    └── StatefulSet
        └── Pod
            └── container(s)
```

---
- Two spaces indent

```sh
01:15:25 > cat testdata/sample4.md | gtree -ts
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
01:16:46 > cat testdata/sample5.md | gtree -fs
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

- Multiple roots

```sh
13:06:26 > cat testdata/sample6.md | gtree
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
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
---

## As Package(1) / like CLI

### Installation
```sh
go get github.com/ddddddO/gtree
```

### Usage

```go
package main

import (
	"bytes"
	"strings"

	"github.com/ddddddO/gtree"
)

func main() {
	r := bytes.NewBufferString(strings.TrimSpace(`
- root
	- dddd
		- kkkkkkk
			- lllll
				- ffff
				- LLL
					- WWWWW
						- ZZZZZ
				- ppppp
					- KKK
						- 1111111
							- AAAAAAA
	- eee`))

	conf := gtree.Config{
		IsTwoSpaces:  false, // `true` when indentation is two half-width spaces
		IsFourSpaces: false, // `true` when indentation is four half-width spaces
	}

	if err := gtree.Execute(os.Stdout, r, conf); err != nil {
		panic(err)
	}
	// root
	// ├── dddd
	// │   └── kkkkkkk
	// │       └── lllll
	// │           ├── ffff
	// │           ├── LLL
	// │           │   └── WWWWW
	// │           │       └── ZZZZZ
	// │           └── ppppp
	// │               └── KKK
	// │                   └── 1111111
	// │                       └── AAAAAAA
	// └── eee
}

```

## As Package(2) / generate a tree programmatically

### Installation
```sh
go get github.com/ddddddO/gtree
```

### Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

func main() {
	root := gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	root.Add("child 5")
	root.Add("child 1").Add("child 2").Add("child 4")
	if err := gtree.ExecuteProgrammably(root, os.Stdout); err != nil {
		panic(err)
	}
	// root
	// ├── child 1
	// │   └── child 2
	// │       ├── child 3
	// │       └── child 4
	// └── child 5

	primate := preparePrimate()
	if err := gtree.ExecuteProgrammably(primate, os.Stdout); err != nil {
		panic(err)
	}
	// Primate
	// ├── Strepsirrhini
	// │   ├── Lemuriformes
	// │   │   ├── Lemuroidea
	// │   │   │   ├── Cheirogaleidae
	// │   │   │   ├── Indriidae
	// │   │   │   ├── Lemuridae
	// │   │   │   └── Lepilemuridae
	// │   │   └── Daubentonioidea
	// │   │       └── Daubentoniidae
	// │   └── Lorisiformes
	// │       ├── Galagidae
	// │       └── Lorisidae
	// └── Haplorrhini
	//     ├── Tarsiiformes
	//     │   └── Tarsiidae
	//     └── Simiiformes
	//         ├── Platyrrhini
	//         │   ├── Ceboidea
	//         │   │   ├── Atelidae
	//         │   │   └── Cebidae
	//         │   └── Pithecioidea
	//         │       └── Pitheciidae
	//         └── Catarrhini
	//             ├── Cercopithecoidea
	//             │   └── Cercopithecidae
	//             └── Hominoidea
	//                 ├── Hylobatidae
	//                 └── Hominidae
}

func preparePrimate() *gtree.Node {
	primate := gtree.NewRoot("Primate")
	strepsirrhini := primate.Add("Strepsirrhini")
	haplorrhini := primate.Add("Haplorrhini")
	lemuriformes := strepsirrhini.Add("Lemuriformes")
	lorisiformes := strepsirrhini.Add("Lorisiformes")

	lemuroidea := lemuriformes.Add("Lemuroidea")
	lemuroidea.Add("Cheirogaleidae")
	lemuroidea.Add("Indriidae")
	lemuroidea.Add("Lemuridae")
	lemuroidea.Add("Lepilemuridae")

	lemuriformes.Add("Daubentonioidea").Add("Daubentoniidae")

	lorisiformes.Add("Galagidae")
	lorisiformes.Add("Lorisidae")

	haplorrhini.Add("Tarsiiformes").Add("Tarsiidae")
	simiiformes := haplorrhini.Add("Simiiformes")

	platyrrhini := haplorrhini.Add("Platyrrhini")
	ceboidea := platyrrhini.Add("Ceboidea")
	ceboidea.Add("Atelidae")
	ceboidea.Add("Cebidae")
	platyrrhini.Add("Pithecioidea").Add("Pitheciidae")

	catarrhini := simiiformes.Add("Catarrhini")
	catarrhini.Add("Cercopithecoidea").Add("Cercopithecidae")
	hominoidea := catarrhini.Add("Hominoidea")
	hominoidea.Add("Hylobatidae")
	hominoidea.Add("Hominidae")

	return primate
}

```