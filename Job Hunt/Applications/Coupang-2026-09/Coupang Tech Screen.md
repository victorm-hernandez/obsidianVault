Design and implement and LRU cache with get and put operations

## Design problem
Make it distributed

Consistent hashing for keys
Redundancy for resiliency with consensus write and read
	- Change data capture to avoid penalty on write and reads.

## Coding problem

using System;

using System.Collections.Generic;

using System.IO;

using System.Linq;

  

namespace Solution {

class Solution {

static void Main(string[] args) {

/* Enter your code here. Read input from STDIN. Print output to STDOUT */

/*

# Requirements

Write (key, value)

Read(key)

Eviction implicit. Length of the cache (memory constraints).

Both read and write operations make a record fresh.

<T> Generic data

# Design

persist items (dictionary/hashtable) O(1)

Affects Read and Write

Max Heap using the last time used as sort factor. O(logn)

Affect Write and Evict

Heap operations

Pop O(1)

Push O(logn)

Sorted collection

Linked List O(n)

Write operation impact is smaller.

Write

CheckForEviction

Evict

FindOldestLinkedList

DeleteElement

Locks

WriteToColl

Dictionary Write

Update usage time. Linked List upsert

#

Read

Dictionary Read

Update usage time. Linked List upsert

Interfaces, dependency injection

LRUCache : CacheInterface

CacheItem

key, value, referenceLinkedListc

CacheUsage

key, timestamp

*/

}

}

}