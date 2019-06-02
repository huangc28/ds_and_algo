# Merge sort

Merge sort belongs to the *divide and conquer* approach. It breaks down a given array to the simplest unit by recursion and merge them back in the sorted order. The time complexity of the worst case is of `O(nLogn)`. The following steps are the general approach to the implementation:

1. If the given length of the array is less then 2, return the array
2. Find the mid position of the array
3. Call `mergeSort` recursively on the left part of the array [:left)
4. Call `mergeSort` recursively on the right part of the array (left:]
5. Merge them in the sorted order. It can be either descending order order acending order, your call.