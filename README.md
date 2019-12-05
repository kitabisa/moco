# moco
moco/mo·co/ - Membaca
 
Library to parse Bank mutation CSV and extract information from it.

## Installation

```bash
$ go get github.com/kitabisa/moco
```
## How to use

```go
//Open file
f, err := os.Open("/dir/to/filename.csv")
if err != nil {
	panic(err)
}

r := moco.NewReader(f, moco.BankBCA)
err = r.ReadMutation()
if err != nil {
	panic(err)
}

success := r.GetSuccess() //Array of MutationBank
failure := r.GetFail() //Array of FailRecord
```

