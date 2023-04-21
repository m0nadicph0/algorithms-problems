
# Segment Trees

Segment trees are a data structure used for efficiently performing range queries on an array or a list. They are particularly useful when the array is subject to updates, and the range queries need to be performed frequently.

The idea behind a segment tree is to divide the array into smaller segments and represent each segment using a node in a binary tree. The root of the tree represents the entire array, and each level of the tree represents a smaller segment of the array. Each node in the tree represents a segment, and it stores information about that segment that can be used to answer range queries.

For example, consider the following array of integers:

[1, 3, 5, 7, 9, 11, 13, 15]

A segment tree for this array would look like this:

```
              [1-8]
             /     \
          [1-4]    [5-8]
         /    \    /   \
      [1-2] [3-4] [5-6] [7-8]
     /    \   /   \   /    \
    [1]  [2] [3] [4] [5]  [6-8]
                       /    \
                      [7]  [8]
```

In this tree, each node represents a segment of the original array. For example, the node labeled [3-4] represents the segment [5, 7], and the node labeled [5-8] represents the segment [9, 11, 13, 15]. Each node stores some information that can be used to answer range queries over its corresponding segment.

For example, let's say we want to find the sum of the values in the segment [2, 6]. We can start at the root node, which represents the entire array, and recursively traverse down the tree, visiting only the nodes that overlap with our query range. At each node, we use the stored information to compute the answer to our query.

In this case, we would visit the nodes labeled [1-4], [3-4], and [5-6]. We can compute the sum of the segment represented by each of these nodes as follows:

```
[1-4] = 1 + 3 + 5 + 7 = 16
[3-4] = 5 + 7 = 12
[5-6] = 11 + 13 = 24
```

Then, we can add up these values to get the sum of the segment [2, 6]:

```
sum([2, 6]) = 16 + 12 + 24 = 52
```

This is just one example of how segment trees can be used to perform range queries efficiently. Other common range queries include finding the minimum or maximum value in a range, or counting the number of elements in a range that satisfy some condition.

## Range Queries



Range queries are operations that involve querying a subset or a range of elements in a data structure. These queries are common in a variety of applications, such as data analysis, machine learning, and database management. Here are five examples of range queries:

1. **Sum of elements in a range**: Given an array of integers and a range [L, R], find the sum of all integers in the range [L, R]. This type of query is common in many applications, such as computing the total revenue of a company over a specific period of time.

2. **Minimum/maximum element in a range**: Given an array of integers and a range [L, R], find the minimum or maximum integer in the range [L, R]. This type of query is useful in many applications, such as finding the highest or lowest temperature over a period of time.

3. **Count of elements in a range that satisfy a condition**: Given an array of integers and a range [L, R], count the number of integers in the range that satisfy a specific condition. For example, you may want to count the number of elements in a range that are greater than a certain threshold.

4. **Substring search**: Given a string and a range [L, R], find all occurrences of a substring within the range [L, R]. This type of query is common in text processing and data mining applications.

5. **Range update**: Given an array of integers and a range [L, R], update all elements in the range [L, R] by a specific amount. This type of query is useful in many applications, such as adjusting the prices of products in a specific category by a fixed amount.

## Data Structure
A segment tree is a binary tree where each leaf node represents an element of an input array, and each internal node represents a segment of the array. The root of the tree represents the entire array, and each internal node is recursively partitioned into two segments until the leaf nodes are reached. The segment represented by an internal node is typically the union of the segments represented by its two child nodes.

Each internal node of the segment tree stores some information computed from the information stored in its child nodes. The information stored in an internal node can be used to answer queries on the segment of the array represented by that node. Common types of information stored in each node include the minimum or maximum value in the segment, the sum of the elements in the segment, or the number of elements in the segment that satisfy a certain condition.

The segment tree can be constructed recursively in a top-down manner, by starting with the root node and dividing the input array into two halves, and recursively constructing the left and right child nodes. The base case of the recursion is when the segment represented by a node contains only one element, which is stored in the corresponding leaf node.

Once the segment tree is constructed, it can be used to answer range queries efficiently. For example, to compute the sum of the elements in a range [L, R], we start at the root of the tree and recursively descend the tree, querying each node to determine if its segment overlaps the query range [L, R]. If it does, we add the information stored in that node to the result and continue recursively descending to its child nodes until we reach the leaf nodes.

To update an element in the input array, we start at the root of the tree and recursively descend the tree to find the leaf node corresponding to the element to be updated. We update the value in that leaf node and then propagate the update information up the tree to update the values in the parent nodes, until we reach the root of the tree.

## Basic Operations

The basic operations on a segment tree are:

1. **Construction**: The segment tree is constructed from the input array in a top-down manner using recursion. The construction process involves dividing the input array into smaller segments, constructing the segment tree for each segment, and computing the information for each internal node based on the information stored in its child nodes.

2. **Query**: Given a query range [L, R], the segment tree can be used to compute some information about the elements in that range. The query operation starts at the root of the tree and recursively descends the tree, querying each node to determine if its segment overlaps the query range [L, R]. If it does, the operation adds the information stored in that node to the result and continues recursively descending to its child nodes until it reaches the leaf nodes.

3. **Update**: Given an index i and a new value v, the segment tree can be updated to reflect the change in the input array. The update operation starts at the root of the tree and recursively descends the tree to find the leaf node corresponding to the element to be updated. The value in that leaf node is updated to v, and then the operation propagates the update information up the tree to update the values in the parent nodes until it reaches the root of the tree.

4. **Lazy Propagation**: In some cases, updating a single element in the input array may require updating a large number of nodes in the segment tree. Lazy propagation is a technique used to delay the updates to the parent nodes until they are actually needed. When an update operation is performed on a node, instead of updating its parent node immediately, the operation stores the update information in the node and sets a flag to indicate that the node has pending updates. When a query operation encounters a node with pending updates, it applies the updates and propagates them to the child nodes before continuing with the query.

These basic operations form the building blocks for many advanced algorithms and data structures that use segment trees, such as range queries and updates, interval trees, and Fenwick trees.

Here's an example implementation of a segment tree in Go:

```go
package main

import "fmt"

type SegmentTree struct {
    tree   []int
    n      int
    update []int
}

func NewSegmentTree(arr []int) *SegmentTree {
    n := len(arr)
    tree := make([]int, 4*n)
    update := make([]int, 4*n)
    return &SegmentTree{tree, n, update}
}

func (st *SegmentTree) build(node, start, end int, arr []int) {
    if start == end {
        st.tree[node] = arr[start]
    } else {
        mid := (start + end) / 2
        st.build(2*node, start, mid, arr)
        st.build(2*node+1, mid+1, end, arr)
        st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
    }
}

func (st *SegmentTree) query(node, start, end, l, r int) int {
    if r < start || end < l {
        return 0
    }
    if l <= start && end <= r {
        return st.tree[node]
    }
    st.push(node, start, end)
    mid := (start + end) / 2
    left := st.query(2*node, start, mid, l, r)
    right := st.query(2*node+1, mid+1, end, l, r)
    return left + right
}

func (st *SegmentTree) update(node, start, end, idx, val int) {
    if start == end {
        st.tree[node] += val
    } else {
        mid := (start + end) / 2
        if start <= idx && idx <= mid {
            st.update(2*node, start, mid, idx, val)
        } else {
            st.update(2*node+1, mid+1, end, idx, val)
        }
        st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
    }
}

func (st *SegmentTree) updateRange(node, start, end, l, r, val int) {
    if r < start || end < l {
        return
    }
    if l <= start && end <= r {
        st.tree[node] += (end - start + 1) * val
        st.update[node] += val
    } else {
        st.push(node, start, end)
        mid := (start + end) / 2
        st.updateRange(2*node, start, mid, l, r, val)
        st.updateRange(2*node+1, mid+1, end, l, r, val)
        st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
    }
}

func (st *SegmentTree) push(node, start, end int) {
    if st.update[node] != 0 {
        mid := (start + end) / 2
        st.update[2*node] += st.update[node]
        st.update[2*node+1] += st.update[node]
        st.tree[2*node] += (mid - start + 1) * st.update[node]
        st.tree[2*node+1] += (end - mid) * st.update[node]
        st.update[node] = 0
    }
}

func main() {
    arr := []int{1, 3, 5, 7,     9, 11}
    st := NewSegmentTree(arr)
    st.build(1, 0, st.n-1, arr)
    fmt.Println(st.query(1, 0, st.n-1, 1, 3))
    st.update(1, 0, st.n-1, 2, 2)
    fmt.Println(st.query(1, 0, st.n-1, 1, 3))
    st.updateRange(1, 0, st.n-1, 0, 5, 2)
    fmt.Println(st.query(1, 0, st.n-1, 1, 3))
}
```

API:
- `NewSegmentTree(arr []int) *SegmentTree`: creates a new instance of SegmentTree struct with the provided integer array `arr`.
- `build(node, start, end int, arr []int)`: builds the segment tree from the given array `arr`.
- `query(node, start, end, l, r int) int`: returns the sum of elements in the range `[l, r]`.
- `update(node, start, end, idx, val int)`: updates the value of the element at index `idx` to `val`.
- `updateRange(node, start, end, l, r, val int)`: adds `val` to all elements in the range `[l, r]`.
- `push(node, start, end int)`: applies the pending updates to the node and its children.

A segment tree is a binary tree where each node represents a range of indices in an array. The root of the tree represents the entire array, and each leaf node represents a single element in the array. The tree is constructed in a way that allows for efficient range queries and updates.

Here is an example of a segment tree for an array of size 8:

```
            [0, 7]
           /      \
        [0, 3]    [4, 7]
        /    \     /    \
     [0,1] [2,3] [4,5] [6,7]
     / \   / \   / \   /  \
   [0] [1][2][3][4][5][6] [7]
```

Each node of the segment tree represents a range of indices in the array. For example, the root node represents the range `[0, 7]`, the left child of the root node represents the range `[0, 3]`, and so on. The leaf nodes represent single elements of the array, and the values stored at each node represent the result of some operation (e.g., sum, minimum, maximum) over the range of indices represented by that node.

In general, the height of a segment tree is `log(n)`, where `n` is the size of the array. This means that the number of nodes in the tree is `O(n)` and the space complexity of a segment tree is also `O(n)`.



Here are 20 examples of problems that can be solved using segment trees:

1. Range sum query: Given an array of integers, and a range [L, R], find the sum of all the integers in the range [L, R].

2. Range minimum query: Given an array of integers, and a range [L, R], find the minimum value in the range [L, R].

3. Range maximum query: Given an array of integers, and a range [L, R], find the maximum value in the range [L, R].

4. Range product query: Given an array of integers, and a range [L, R], find the product of all the integers in the range [L, R].

5. Range gcd query: Given an array of integers, and a range [L, R], find the greatest common divisor (GCD) of all the integers in the range [L, R].

6. Range lcm query: Given an array of integers, and a range [L, R], find the least common multiple (LCM) of all the integers in the range [L, R].

7. Range update query: Given an array of integers, and a range [L, R], add a value to all the integers in the range [L, R].

8. Range set query: Given an array of integers, and a range [L, R], set all the integers in the range [L, R] to a given value.

9. Range count query: Given an array of integers, and a range [L, R], count the number of integers in the range [L, R] that are equal to a given value.

10. Range mode query: Given an array of integers, and a range [L, R], find the mode (most frequent element) in the range [L, R].

11. Range median query: Given an array of integers, and a range [L, R], find the median (middle element) in the range [L, R].

12. Range kth query: Given an array of integers, and a range [L, R], find the kth smallest/largest element in the range [L, R].

13. Range reverse query: Given an array of integers, and a range [L, R], reverse the order of all the integers in the range [L, R].

14. Range binary search query: Given an array of integers sorted in ascending order, and a range [L, R], find the index of the first/last element in the range [L, R] that is greater than or equal to a given value.

15. Range update with lazy propagation: Given an array of integers, and a range [L, R], add a value to all the integers in the range [L, R] efficiently using lazy propagation.

16. Range gcd with lazy propagation: Given an array of integers, and a range [L, R], find the greatest common divisor (GCD) of all the integers in the range [L, R] efficiently using lazy propagation.

17. Range minimum with lazy propagation: Given an array of integers, and a range [L, R], find the minimum value in the range [L, R] efficiently using lazy propagation.

18. Range maximum with lazy propagation: Given an array of integers, and a range [L, R], find the maximum value in the range [L, R] efficiently using lazy propagation.

19. Range sum with lazy propagation: Given an array of integers, and a range [L, R], find the sum of all the integers in the range [L, R] efficiently using lazy propagation.

20. Range set with lazy propagation: Given an array of integers, and a range

Here are 20 problems from LeetCode that involve segment trees:

1. Range Sum Query - Mutable: Implement a data structure that supports range sum queries and range updates efficiently.

2. Count of Smaller Numbers After Self: Given an integer array, for each element, count the number of elements to its right that are smaller than it.

3. Range Sum Query 2D - Mutable: Implement a data structure that supports 2D range sum queries and range updates efficiently.

4. Maximum Sum of Subarray Close to K: Given an integer array and a target value K, find the maximum sum of a subarray such that the sum is closest to K.

5. Number of Inversions in an Array: Given an integer array, count the number of inversions (i.e., pairs of elements that are out of order).

6. Longest Repeating Character Replacement: Given a string and an integer k, find the length of the longest substring that can be obtained by replacing at most k characters.

7. Minimum Number of Arrows to Burst Balloons: Given a list of balloons, find the minimum number of arrows needed to burst all the balloons.

8. Largest Rectangle in Histogram: Given an array representing the heights of a histogram, find the area of the largest rectangle that can be formed.

9. Kth Smallest Element in a Sorted Matrix: Given a matrix of integers that is sorted both row-wise and column-wise, find the kth smallest element in the matrix.

10. The Skyline Problem: Given a list of buildings, find the skyline of the city.

11. Count of Range Sum: Given an integer array, count the number of range sums that fall within a given range.

12. Sliding Window Median: Given an integer array and a window size k, find the median of each window of size k that slides from left to right.

13. Range Module: Implement a data structure that supports adding, removing, and querying ranges of integers.

14. Count of Smaller Numbers Before Self: Given an integer array, for each element, count the number of elements to its left that are smaller than it.

15. Rectangle Area II: Given a list of rectangles, find the total area covered by the rectangles.

16. Maximal Rectangle: Given a matrix of 0's and 1's, find the area of the largest rectangle that can be formed by 1's.

17. The Number of Weak Characters in the Game: Given a list of game characters, a character is considered weak if it has strictly smaller attack and defense than another character.

18. Smallest Common Region: Given a list of regions and a pair of regions, find the smallest region that contains both regions.

19. Shortest Distance to Target Color: Given an integer array and a list of color values, for each index in the array, find the shortest distance to the nearest index that has the target color.

20. The Skyline Problem II: Given a list of buildings, find the silhouette of the city.
