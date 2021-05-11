# rbtree
The Red-Black Tree implementation in Go, by dissecting the vein of its relationship with 2-3-4 tree. To get more information, check my blog: https://blog.peven.me/post/2021/05/08/the-red-black-tree/

## How to use
To get interactive with rbtree, you need implement the KeyType interface, which defines the behavior `CompareTo`.
```go
// KeyType is the type with CompareTo behavior.
	KeyType interface {
		CompareTo(c interface{}) int
	}
```
There is a `KeyTypeInt` out of box, which implement the `CompareTo` of `int`.

## Example
```go
package main

import (
        "fmt"

        "github.com/phoon/rbtree"
)

type FooBar struct {
        Foo string
        Bar string
}

func main() {
        t := rbtree.NewRBT()

        data := FooBar{Foo: "foo", Bar: "bar"}

        // Insert
        t.Insert(rbtree.KeyTypeInt(1), data)

        // Search
        res := t.Search(rbtree.KeyTypeInt(1))
        var res_d FooBar
        if res != nil {
                res_d = res.(FooBar)
        }
        fmt.Println(res_d)

        // Delete
        t.Remove(rbtree.KeyTypeInt(1))
}
```