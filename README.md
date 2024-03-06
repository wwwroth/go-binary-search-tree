# Go Binary Search Tree
This is a project written in go that accepts a binary tree represented in JSON and outputs an ASCII diagram for it as well as if it's a valid BST or not. You can also generate valid BSTs with a defined height.

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
go run main.go -tree '{"value":10,"left":{"value":5,"left":{"value":3},"right":{"value":7}},"right":{"value":15,"right":{"value":1}}}'
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

#### Generate JSON for a valid BST of N height

```bash
go run main.go -generate 4
```
```json
{
  "value": 8,
  "left": {
    "value": 0,
    "left": {
      "value": -4,
      "left": {
        "value": -6
      },
      "right": {
        "value": -2
      } 
    },
    "right": {
      "value": 4,
      "left": {
        "value": 2
      },
      "right": {
        "value": 6
      }
    }
  },
  "right": {
    "value": 16,
    "left": {
      "value": 12,
      "left": {
        "value": 10
      },
      "right": {
        "value": 14
      }
    },
    "right": {
      "value": 20,
      "left": {
        "value": 18
      },
      "right": {
        "value": 22
      }
    }
  }
}
```