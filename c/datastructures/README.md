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
| 1 | **Stack (LIFO)** | Generics, push/pop/peek, error handling |
| 2 | **Queue (FIFO)** | Slices, circular buffer, enqueue/dequeue |
| 3 | **Linked List** | Structs, pointers, singly vs doubly linked lists |

---

## 🟨 Intermediate Assignments

| # | Assignment | Key Concepts |
|---|------------|--------------|
| 4 | **Hash Map (basic)** | Buckets, hashing, collision handling |
| 5 | **Binary Search Tree (BST)** | Recursion, tree traversal, delete operation |
| 6 | **Min/Max Heap** | Heapify, priority queues |
| 7 | **Trie (Prefix Tree)** | String processing, nested maps/structs |

---

## 🟧 Advanced Assignments

| # | Assignment | Key Concepts |
|---|------------|--------------|
| 8  | **Graph Implementation** | Adjacency list vs matrix, BFS/DFS |
| 9  | **LRU Cache** | Linked list + hash map, O(1) operations |
| 10 | **Disjoint Set / Union-Find** | Path compression, union by rank |
| 11 | **Bloom Filter** | Probabilistic DS, multiple hash functions |

---

## 🟪 Go-Idiomatic Enhancements

- **Testing:** Use `go test` to write unit tests for each structure.
- **Benchmarks:** Use `go test -bench .` to measure performance.
- **Interfaces:** Implement `fmt.Stringer` for pretty printing.
- **Generics:** Make data structures reusable with type parameters.
- **Examples:** Provide small `main.go` demos of each data structure.

---

## ✅ Acceptance Criteria for Each Assignment

Each data structure must:

1. **Be implemented as a reusable type** (struct + methods, using generics where appropriate).
2. **Include error handling** (e.g., popping from an empty stack returns an error).
3. **Include unit tests** using the `testing` package:
    - Happy paths.
    - Edge cases.
5. **Optional:** Implement `fmt.Stringer` for readable printing.
6. **Optional:** Add benchmarks with `go test -bench .`.

---

### 🟩 Beginner Assignments

#### 1. Stack (LIFO)
- **Required Methods:** `Push`, `Pop`, `Peek`, `Size`.
- **Behavior:** Last-in-first-out. Popping or peeking from an empty stack returns an error.
- **Tests:** Push multiple items, pop all, check size.

#### 2. Queue (FIFO)
- **Required Methods:** `Enqueue`, `Dequeue`, `Peek`, `Size`.
- **Behavior:** First-in-first-out. Dequeue/peek on empty queue returns an error.
- **Bonus:** Implement circular buffer behavior.

#### 3. Linked List
- **Required Methods:** `Append`, `Prepend`, `InsertAfter(node, value)`, `Delete(value)`, `Find(value)`.
- **Behavior:** Maintain head (and tail for doubly linked).
- **Tests:** Insert and remove at various positions.

---

### 🟨 Intermediate Assignments

#### 4. Hash Map (basic)
- **Required Methods:** `Put(key, value)`, `Get(key)`, `Delete(key)`, `Size`.
- **Behavior:** Handle collisions (separate chaining or open addressing).
- **Tests:** Insert, retrieve, overwrite, delete, collision cases.

#### 5. Binary Search Tree
- **Required Methods:** `Insert`, `Find`, `Delete`, Traversals (`InOrder`, `PreOrder`, `PostOrder`).
- **Behavior:** Maintain BST property. Delete must handle nodes with 0, 1, or 2 children.
- **Tests:** Insert, delete, traverse.

#### 6. Min/Max Heap
- **Required Methods:** `Insert`, `ExtractMin` (or `ExtractMax`), `Peek`.
- **Behavior:** Maintain heap property after each operation.
- **Tests:** Insert random numbers, extract in order.

#### 7. Trie (Prefix Tree)
- **Required Methods:** `Insert(word)`, `Search(word)`, `StartsWith(prefix)`.
- **Behavior:** Support efficient prefix search.
- **Tests:** Insert multiple words, search prefixes.

---

### 🟧 Advanced Assignments

#### 8. Graph Implementation
- **Required Methods:** `AddVertex`, `AddEdge`, `RemoveVertex`, `RemoveEdge`, `BFS`, `DFS`.
- **Behavior:** Support directed or undirected graphs.
- **Tests:** Build a small graph, traverse.

#### 9. LRU Cache
- **Required Methods:** `Get(key)`, `Put(key, value)`.
- **Behavior:** O(1) average time for both operations using a doubly linked list + hash map.
- **Tests:** Insert over capacity, verify eviction.

#### 10. Disjoint Set / Union-Find
- **Required Methods:** `Find(x)`, `Union(x, y)`.
- **Behavior:** Path compression & union by rank.
- **Tests:** Multiple unions and finds.

#### 11. Bloom Filter
- **Required Methods:** `Add(item)`, `MightContain(item)`.
- **Behavior:** Multiple hash functions, low memory, false positives allowed.
- **Tests:** Known elements return true; non-elements return mostly false.

---

## 📝 Example First Assignment (Stack)

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
