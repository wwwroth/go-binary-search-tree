# Go Binary Search Tree
This is a project written in go that accepts a binary tree represented in JSON and outputs an ASCII diagram for it as well as if it's a valid BST or not.

## Usage

#### Valid Trees

```bash
go run main.go -validate '{"value":10,"left":{"value":5,"left":{"value":3},"right":{"value":7}},"right":{"value":15,"right":{"value":18}}}'
```

```
        ┌── 18
    ┌── 15
┌── 10
        ┌── 7
    └── 5
        └── 3
2024/03/05 19:58:01 The binary tree is a valid BST
```

#### Invalid Tree

```shell
go run main.go -validate '{"value":10,"left":{"value":5,"left":{"value":3},"right":{"value":7}},"right":{"value":15,"right":{"value":1}}}'
```

```
        ┌── 1
    ┌── 15
┌── 10
        ┌── 7
    └── 5
        └── 3
2024/03/05 19:57:24 The binary tree is not a valid BST
```