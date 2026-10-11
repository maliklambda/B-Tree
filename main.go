package main 

import(
    "fmt"
    "time"
    "github.com/MalikL2005/btree_SeliaDB/btree"
)



func main(){
    fmt.Println("Hello there")
    root := &btree.Node_t{}
    i := 1
    start := time.Now()
    var err error
    for ; i<=21; i++ {
        err = btree.Insert(&root, root, btree.Entry_t{Key: uint32(i), Value: uint32(i)})
        if err != nil {
            panic(1)
        }
    }
    fmt.Println("Traversing")
    btree.Traverse(root, root)
    btree.Insert(&root, root, btree.Entry_t{Key: uint32(i), Value: uint32(i)})
    fmt.Println("Traversing")
    btree.Traverse(root, root)
    end := time.Since(start)
    fmt.Println("Process to insert", i, "elements took", end)

		btree.WriteBtree(&root, root)

    // Deletion
    // if err := btree.Delete(&root, root, uint32(18)); err != nil {
    //     fmt.Println(err)
    // }
    // if err := btree.Delete(&root, root, uint32(19)); err != nil {
    //     fmt.Println(err)
    // }
    // if err := btree.Delete(&root, root, uint32(20)); err != nil {
    //     fmt.Println(err)
    // }
    // if err := btree.Delete(&root, root, uint32(21)); err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Println("Traversing")
    // btree.Traverse(root, root)

    
    fmt.Println()
    btree.FilterForGreaterKey (&root, root, 12)

}
