
# Segment Trees

Range queries are a type of query that involves finding an aggregate value 
(such as minimum, maximum, sum, or average) over a contiguous subsequence 
or range of elements in an array or data structure.

For example, consider an array of integers:

```
 0  1  2  3  4  5  6  7
[4, 6, 1, 8, 3, 9, 7, 5]
``` 
Here are some examples of range queries that can be performed on this array:

1. **Finding the minimum value between index 2 and index 5**: The subsequence [1, 8, 3, 9] lies between index 2 and index 5. The minimum value in this subsequence is 1.
2. **Finding the sum of all values between index 1 and index 6**: The subsequence [6, 1, 8, 3, 9, 7] lies between index 1 and index 6. The sum of all values in this subsequence is 34.
3. **Finding the maximum value in the entire array**: The maximum value in the array is 9.

Range queries are commonly used in a variety of applications, such as 
* data analysis
* image processing
* computational geometry. 

* By using data structures such as `segment trees` or _binary indexed trees_, range queries can be performed 
* efficiently in logarithmic time, allowing for the processing of large data sets in real-time.

## Segment Trees

Segment Trees, also known as interval trees, are a type of data structure used for 
efficient querying and updating of intervals or segments in an array. 
They are particularly useful for solving problems that involve range queries, such as finding the minimum, maximum, or sum of all elements in a given range.

A segment tree works by recursively partitioning the array into smaller sub-arrays and storing the values that correspond to each subarray in a tree data structure. 
Each node of the tree represents a segment or interval of the original array, and the values stored at each node are determined by aggregating the values of its children.

For example, in a segment tree for an array of integers
* each leaf node would represent a single element of the array, 
* while the internal nodes would represent segments of the array. 
* The value stored at each internal node would be the aggregate of its children's values, such as the sum or minimum of their values.


## Structure
A segment tree is a binary tree data structure that is used to represent intervals or segments of an array. The structure of a segment tree is as follows:

1. Each node in the tree represents a segment or interval of the input array. The root node represents the entire array, while the leaf nodes represent individual elements of the array.
2. Each internal node in the tree has two children, which correspond to the left and right halves of its interval. For example, if a node represents the interval [i, j], its left child represents the interval [i, (i+j)/2] and its right child represents the interval [(i+j)/2+1, j].
3. The values stored at each node of the tree are determined by aggregating the values of its children. This aggregation operation can be any associative function, such as sum, maximum, minimum, or bitwise AND.
4. The tree is constructed recursively, by dividing the input array into smaller sub-arrays and building the corresponding segments of the tree.
5. The size of the segment tree is typically 4n, where n is the size of the input array, since each node in the tree corresponds to a segment of at most length n.

The structure of a segment tree allows for efficient querying and updating of intervals in the input array, 
by traversing the tree and aggregating the values of its nodes. The logarithmic depth of the tree ensures 
that the time complexity of these operations is O(log n), where n is the size of the input array.

## Basic Operations

A segment tree supports several basic operations that can be performed efficiently in logarithmic time. Here are the most commonly used operations:

1. `Build`: This operation constructs the segment tree for a given input array. It is performed recursively by dividing the input array into smaller subarrays and building the corresponding segments of the tree.
2. `Query`: This operation computes an aggregate value (such as minimum, maximum, sum, or average) over a given interval or range of the input array. It is performed by traversing the tree and aggregating the values of its nodes that overlap with the given range.
3. `Update`: This operation modifies the value of a single element in the input array and updates the corresponding node in the segment tree. It is performed by traversing the tree from the root to the leaf node that represents the target element, and then updating the node and its ancestors as necessary to maintain the correct aggregation values.
4. `Lazy propagation`: This is an optimization technique that allows for efficient updates of a large range of elements in the input array. Instead of updating all the affected nodes immediately, lazy propagation defers the updates until they are actually needed. This can greatly reduce the number of updates performed and improve the performance of the segment tree.
5. `Range modification`: This operation modifies the values of all elements in a given interval or range of the input array, and updates the corresponding nodes in the segment tree. It is similar to the update operation, but it affects multiple nodes in the tree instead of just one.
6. `Range queries with condition`: This operation computes an aggregate value over a given interval or range of the input array that satisfies a certain condition. For example, it can compute the minimum value over a range that is greater than a certain threshold. This operation can be implemented using a modified version of the query operation that checks the condition at each node of the tree before aggregating its values.

### Examples


#### Example 1: 

Input array:

```
 0  1  2  3
[1, 3, 2, 4]
```


```
graph TD
  A["[0,3] 10"]
  B["[0,1] 4"]
  C["[2,3] 6"]
  D["[0,0] 1"]
  E["[1,1] 3"]
  F["[2,2] 2"]
  G["[3,3] 4"]
  A-->B
  A-->C
  B-->D
  B-->E
  C-->F
  C-->G
```

Explanation: 
* This segment tree represents the input array [1, 3, 2, 4] with 0-based indexing. 
* The root node (A) represents the interval [0,3], which is the entire array. 
* Its left child (B) represents the interval [0,1], and its right child (C) represents the interval [2,3]. 
* The leaf nodes contain the array elements themselves. 
* The value next to each node represents the aggregate value of the interval that it represents. E.g: the root node A represents the sum of all elements in the array, which is 10.

#### Example 2: 
Input array:

```
 0  1  2  3  4   5  6
[7, 5, 3, 9, 11, 6, 2]
```


```
graph TD
  A["[0,6] 43"]
  B["[0,3] 24"]
  C["[4,6] 19"]
  D["[0,1] 12"]
  E["[2,3] 12"]
  F["[4,5] 17"]
  G["[6,6] 2"]
  H["[0,0] 7"]
  I["[1,1] 5"]
  J["[2,2] 3"]
  K["[3,3] 9"]
  L["[4,4] 11"]
  M["[5,5] 6"]
  N["[6,6] 2"]
  A-->B
  A-->C
  B-->D
  B-->E
  C-->F
  C-->G
  D-->H
  D-->I
  E-->J
  E-->K
  F-->L
  F-->M
  G-->N
```

Explanation: 
* This segment tree represents the input array [7, 5, 3, 9, 11, 6, 2] with 0-based indexing. 
* The root node (A) represents the interval [0,6], which is the entire array. 
* Its left child (B) represents the interval [0,3], 
* and its right child (C) represents the interval [4,6]. 
* The leaf nodes contain the array elements themselves. For example, the node H represents the element at index 0, which is 7. 
* The value next to each node represents the aggregate value of the interval that it represents. For example, the root node A represents the sum of all elements in the array, which is 43.


## Examples:

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
