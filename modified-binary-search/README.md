# Modified binary search

* variation of the traditional binary search algorithm that is used to search for an element in a sorted array. 
* used when the array is not sorted in a traditional ascending or descending order, but rather has a specific property that can be exploited to efficiently search for an element.

In modified binary search, the search space is divided into two parts, just like in traditional binary search. 
However, the condition for moving left or right is modified based on the specific property of the array. 
For example, in a rotated sorted array, the mid-element is compared with the first and last elements of the 
array to determine whether to move left or right.

## Uses
* finding the peak element in an array 
* searching for an element in a matrix. 

## Problems
1. `Search in Rotated Sorted Array`: Given a rotated sorted array, where the rotation point is unknown, find a specific element in the array.
2. `Find Minimum in Rotated Sorted Array`: Given a rotated sorted array, where the rotation point is unknown, find the minimum element in the array.
3. `Peak Element`: Given an array of integers, find the peak element in the array. A peak element is an element that is greater than or equal to its neighbors.
4. `Find Kth Smallest Element in a Sorted Matrix`: Given a matrix where each row and column is sorted in non-decreasing order, find the kth smallest element in the matrix.
5. `Find Median from Data Stream`: Given an array of integers, find the median element in the array as the elements are being inserted one by one.
6. `Search in a Bitonic Array`: Given a bitonic array, where the elements first increase and then decrease, find a specific element in the array.
7. `Find First and Last Position of Element in Sorted Array`: Given a sorted array, find the first and last position of a specific element in the array.
8. `Count of Smaller Numbers After Self`: Given an array of integers, for each element in the array, count the number of elements to the right of it that are smaller than it.
9. `Capacity To Ship Packages Within D Days`: Given an array of weights and a capacity limit, find the minimum capacity required to ship all the packages within D days.
10. `Split Array Largest Sum`: Given an array of integers and a number k, split the array into k non-empty sub-arrays, such that the maximum sum of any subarray is minimized.

## Leetcode
11. Search in Rotated Sorted Array: <ins>https://leetcode.com/problems/search-in-rotated-sorted-array/</ins>
12. Find Minimum in Rotated Sorted Array: <ins>https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/</ins>
13. Peak Index in a Mountain Array: <ins>https://leetcode.com/problems/peak-index-in-a-mountain-array/</ins>
14. Find Kth Smallest Element in a Sorted Matrix: <ins>https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/</ins>
15. Median of Two Sorted Arrays: <ins>https://leetcode.com/problems/median-of-two-sorted-arrays/</ins>
16. Capacity To Ship Packages Within D Days: <ins>https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/</ins>
17. Split Array Largest Sum: <ins>https://leetcode.com/problems/split-array-largest-sum/</ins>
18. Koko Eating Bananas: <ins>https://leetcode.com/problems/koko-eating-bananas/</ins>
19. Minimum Limit of Balls in a Bag: <ins>https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/</ins>
20. Minimum Speed to Arrive on Time: <ins>https://leetcode.com/problems/minimum-speed-to-arrive-on-time/</ins>

## UVA

1.  UVA 10277 - Stack 'em Up: This problem involves shuffling and dealing cards using a stack. Modified binary search can be used to determine the final order of the cards after a certain number of shuffles.
2.  UVA 11235 - Frequent values: This problem requires finding the most frequent value in a range of an array. Modified binary search can be used to optimize the query process.
3.  UVA 11935 - Through the Desert: This problem involves simulating a car traveling through a desert with limited fuel. Modified binary search can be used to determine the minimum amount of fuel required to reach the destination.
4.  UVA 12468 - Zapping: This problem involves determining the minimum number of button presses required to change the channel on a TV. Modified binary search can be used to find the shortest distance between two channels.
5.  UVA 1337 - Queen: This problem involves determining the maximum number of queens that can be placed on a chessboard without attacking each other. Modified binary search can be used to find the first position where the number of queens exceeds a certain limit.
6.  UVA 1476 - In Search of the Biggest: This problem involves finding the largest subset of a given array such that the absolute difference between any two elements is less than or equal to a given value. Modified binary search can be used to optimize the search process.
7.  UVA 12192 - Grapevine: This problem involves finding the largest square submatrix in a given matrix such that the absolute difference between any two elements is less than or equal to a given value. Modified binary search can be used to optimize the search process.
8.  UVA 1370 - Biased Standings: This problem involves sorting a list of teams based on their scores while accounting for bias in the reporting of the scores. Modified binary search can be used to determine the position of a team in the final standings.
9.  UVA 1594 - Repeating Decimals: This problem involves determining the repeating decimal expansion of a given fraction. Modified binary search can be used to find the first position where a repeating pattern occurs.
10. UVA 1336 - Alphacode: This problem involves decoding a string of digits into an alphabetic string. Modified binary search can be used to determine the number of valid decoding possibilities.

