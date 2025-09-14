# Go Data Structures Practice

This repository is a self-directed set of “homework” problems to practice implementing classic data structures in **Go**.  
The goal: learn Go syntax, generics, and idioms while refreshing core CS concepts.

---

## 📚 Overview

Each assignment focuses on a classic data structure. You’ll:

- Implement it in Go using idiomatic patterns.
- Write unit tests with the `testing` package.
- Benchmark performance against Go’s built-in equivalents where applicable.
- Use **generics** to make reusable components (Go 1.18+).

---

## 🟩 Beginner Assignments

| # | Assignment | Key Concepts |
|---|------------|--------------|
| 1 | **Dynamic Array (like ArrayList)** | Slices, resizing, append/insert/remove |
| 2 | **Stack (LIFO)** | Generics, push/pop/peek, error handling |
| 3 | **Queue (FIFO)** | Slices, circular buffer, enqueue/dequeue |
| 4 | **Linked List** | Structs, pointers, singly vs doubly linked lists |

---

## 🟨 Intermediate Assignments

| # | Assignment | Key Concepts |
|---|------------|--------------|
| 5 | **Hash Map (basic)** | Buckets, hashing, collision handling |
| 6 | **Binary Search Tree (BST)** | Recursion, tree traversal, delete operation |
| 7 | **Min/Max Heap** | Heapify, priority queues |
| 8 | **Trie (Prefix Tree)** | String processing, nested maps/structs |

---

## 🟧 Advanced Assignments

| # | Assignment | Key Concepts |
|---|------------|--------------|
| 9  | **Graph Implementation** | Adjacency list vs matrix, BFS/DFS |
| 10 | **LRU Cache** | Linked list + hash map, O(1) operations |
| 11 | **Disjoint Set / Union-Find** | Path compression, union by rank |
| 12 | **Bloom Filter** | Probabilistic DS, multiple hash functions |

---

## 🟪 Go-Idiomatic Enhancements

- **Testing:** Use `go test` to write unit tests for each structure.
- **Benchmarks:** Use `go test -bench .` to measure performance.
- **Interfaces:** Implement `fmt.Stringer` for pretty printing.
- **Generics:** Make data structures reusable with type parameters.
- **Examples:** Provide small `main.go` demos of each data structure.

---

## 📝 Example First Assignment

**Implement a Stack in Go using Generics**

- Create a file `stack.go` with:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) { /* ... */ }
func (s *Stack[T]) Pop() (T, error) { /* ... */ }
func (s *Stack[T]) Peek() (T, error) { /* ... */ }
func (s *Stack[T]) Size() int { /* ... */ }
func (s *Stack[T]) String() string { /* ... */ }
