# Two Pointers

## Motivations

- To compare values
- To keep track of values
- As an intermediary step of another process, e.g. search

## When to use

- On a linear data structure, such as array or linked list
  - Brute force: O(n^2)
  - If data structure has predictable dynamics, it can be cheaper
- On structures with predictable dynamics
  - Easy to predict whether the next value that a pointer is being moved to is greater or smaller
  - So we can use a logical approach instead of brute force to move pointers
  - Example: Sorted arrays
- On problems that ask for a pair of values or a result that can be generated
  from two values

## Strategies: How to use

Two-pointer strategies take O(n) time

1. **Inward traversal**
   - Pointers start at opposite ends and move/converge toward the center
   - Usually pointers adjust their positions based on comparisons, until a
     certain condition is met, or until they meet/cross each other
2. **Unidirectional traversal**
   - Pointers start on same side and move in same direction
   - Usually right pointer is used to find information and the left pointer to
     keep track of information
3. **Staged traversal**
   - Traverse with first pointer then derive start position of second pointer
     from first pointer
   - Usually first pointer is used to search for something, then once found,
     second pointer finds additional information pertaining to the value at the
     first pointer
