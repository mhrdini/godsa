# Hash Maps and Sets

## Idea

Use a hash function to:

- map unique keys to values (in the case of hash map), or
- store unique keys (in the case of hash set)

## Motivations

- Immediate lookup

## Time Complexity

| Operation | Average | Worst |
| --------- | ------- | ----- |
| Insert    | O(1)    | O(n)  |
| Access    | O(1)    | O(n)  |
| Delete    | O(1)    | O(n)  |

- Can be O(n) when the hash function produces many collisions such that the
  operation must go through all entries

## When to use

- Storing key-value pairs
- Keeping track of state, aggregates
- Quick lookups
- Problems with keywords like:
  - frequency
  - unique
  - map
  - dictionary
  - fast lookup
