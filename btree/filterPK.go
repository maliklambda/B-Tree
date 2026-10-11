package btree

import(
    "fmt"
)


func FilterForGreaterKey (root **Node_t, current *Node_t, comparedKey uint32) {
    if current == nil {
        return
    }

    fmt.Println("New Node", current.NumOfEntries)
    for i, entry := range *current.Entries {
        if entry.Key < comparedKey {
            continue
        }
        fmt.Println(i, ":", entry)
    }

    if current.Children == nil {
        return
    }
    for i, child := range *current.Children {
        if i < len(*current.Entries) && (*current.Entries)[i].Key < comparedKey {
            continue
        }
        FilterForGreaterKey (root, &child, comparedKey)
    }
}



