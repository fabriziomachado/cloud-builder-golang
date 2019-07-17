# GoLang no GCP Cloud Build

#### 1) Você deverá criar uma simples aplicação que possua uma função chamada soma que receberá dois parâmetros e retornará a adição desses dois valores.

```golang
package main

import ("fmt")

func add(x, y int) int {
    return x + y
}

func main () {
    var x,y = 5, 5
    total := add(x, y)
    fmt.Printf("Total da Soma de %v + %v é %v", x, y, total);
}
```

#### 2) Crie um teste unitário para essa função.

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(5, 5)
    want := 10
    if got != want {
       t.Errorf("Sum was error got: %d want %d!", got, want)
    }
}
```

#### 3) Ative um processo de CI que execute os seguintes passos:
![CI](/ci-gcp.png)

## Imagem registrada no GPC

$ docker run gcr.io/codeeducation-245712/calculator

```
Total da Soma de 5 + 5 é 10
```
