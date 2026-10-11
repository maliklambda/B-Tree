package btree

import (
	"errors"
	"fmt"
	"slices"
)



func Delete (root **Node_t, current *Node_t, key uint32) error {
    if current == nil {
        return errors.New("Current may not be nil")
    }

    // current is leaf node and current has more than MIN_KEYS+1
    if (current.Children == nil || len(*current.Children) == 0) && len(*current.Entries) >= MIN_KEYS+1 {
        fmt.Println(current.Children)
        fmt.Println(len(*current.Entries))
        fmt.Println("Found right node")
        fmt.Println("Deleting", key, "from", *current.Entries)
        // entry can simply be removed
        deleteIndex := findIndex(current.Entries, key)
        if deleteIndex < 0 {
            return errors.New("Entry not found in current.Entries")
        }
        *current.Entries = slices.Delete(*current.Entries, deleteIndex, deleteIndex+1)
        current.NumOfEntries--
        return nil
    }

    // current is leaf node and current current has less than MIN_KEYS+1
    if (current.Children == nil || len(*current.Children) == 0) && len(*current.Entries) < MIN_KEYS+1 {
        parent, childIndex := findParent(*root, current)
        if childIndex < 0 {
            return errors.New("No parent found")
        }
        // can borrow from right sibling
        if len(*parent.Children)>childIndex+1 && checkRightSibling(parent, childIndex){
            fmt.Println("Switching right sibling:", (*parent.Children)[childIndex].Entries, "and", (*parent.Children)[childIndex+1].Entries)
            err := insertEntry(current, (*parent.Entries)[childIndex])
            if err != nil {
                return err
            }
            deleteIndex := findIndex(current.Entries, key)
            *current.Entries = slices.Delete(*current.Entries, deleteIndex, deleteIndex+1)
            (*parent.Entries)[childIndex] = ((*(*parent.Children)[childIndex+1].Entries)[0])
            return nil
        }
        // can borrow from left sibling
        if checkLeftSibling(parent, childIndex){
            fmt.Println("Switching left sibling:", (*parent.Children)[childIndex].Entries, "and", (*parent.Children)[childIndex-1].Entries)
            err := insertEntry(current, (*parent.Entries)[childIndex-1])
            if err != nil {
                return err
            }
            indexNewParent := len(*(*parent.Children)[childIndex-1].Entries) -1
            (*parent.Entries)[childIndex-1] = (*(*parent.Children)[childIndex-1].Entries)[indexNewParent-1]
            fmt.Println(*(*parent.Children)[childIndex-1].Entries)
            fmt.Println(childIndex)
            *(*parent.Children)[childIndex-1].Entries = slices.Delete(*(*parent.Children)[childIndex-1].Entries, indexNewParent, indexNewParent+1)
            deleteIndex := findIndex(current.Entries, key)
            if deleteIndex < 0 {
                return errors.New("index not found")
            }
            *current.Entries = slices.Delete(*current.Entries, deleteIndex, deleteIndex+1)
            return nil
        }
        // merge parent and child (pull parent down)
        if len(*parent.Entries) >= MIN_KEYS+1 {
            fmt.Println("Pulling parent down")
        }


        return errors.New("Can't delete this")
    }

    if len(*current.Children) > 0 {
        childIndex := findChildIndex(*current, key)
        return Delete(root, &(*current.Children)[childIndex], key)
    }

    return nil
}



func findIndex (entries *[]Entry_t, key uint32) int {
    for i, entry := range *entries {
        if entry.Key == key {
            return i
        }
    }
    return -1
}
