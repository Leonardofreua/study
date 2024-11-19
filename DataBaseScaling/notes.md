# Database Scaling

## ACID

- **Atomicity** - Each transaction is all or nothing
- **Consistency** - Any transaction will bring the database from one valid state to another
- **Isolation** - Executing transactions concurrently has the same results as if the transactions were executed serially
- **Durability** - Once a transaction has been committed, it will remain so

## Master-Slave replication

**Disadvantages:**

- Additional logic is needed **to promote a slave to a master**

<p align="center">
  <img alt="mock" width="500" src="images/master_slave_replication.png">
</p>

## Master-Master replication

**Disadvantages:**

- Need a load balancer or make changes to your application logic to determine where to write;
- Most master-master systems are either loosely consistent (violating ACID) or have increased write latency due to synchronization
- Conflict resolution comes more into play as more write nodes are added and as latency increases

**Disadvantages for replication:**

- Potential for loss of data if the master fails before any newly written data can be replicated to other nodes
- The more read slaves, the more you have to replicate, which leads to greater replication lag.
- Writing to the master can spawn multiple threads to write in parallel, whereas read replicas only support writing sequentially with a single thread
- Replication adds more hardware and additional complexity

<p align="center">
  <img alt="mock" width="500" src="images/master_master_replication.png">
</p>

## Federation

**Advantages:**

- Splits up database by function (functional partitioning)
- For example, you could have three databases instead of a monolith database: forums, users, and products. Resulting in less read and write traffic to each database and less replication lag
    - Smaller databases result in more data that can fit in memory
    - Improve cache locality
    - Can write in parallel, increasing throughput

**Disadvantages:**

- Federation is not effective if your schema requires huge functions or tables.
- Adjust application logic to determine which database to read and write
- Joining data from two databases is more complex with a server link
- More hardware and complexity

<p align="center">
  <img alt="mock" width="500" src="images/federation.png">
</p>

## Sharding

**Advantages:**

- Similar to the federation
- less read and write traffic
- less replication
- more cache hits
- index size is also reduced - faster queries
- If one shard goes down, the other shards are still operational
- allow write in parallel
- Commons way to shard a table of users is using User’s last name initial or user’s geographic location

**Disadvantages:**

- Update application logic to work with shard. This could result in complex SQL queries
- Data distribution can become unbalanced. For example, a set of power users on a shard could result in increased load to that shard compared to others:
    - Rebalancing adds additional complexity
    - Consistent hashing can reduce the amount of transferred data
- Joining data from two databases is more complex with a server link
- More hardware and complexity

<p align="center">
  <img alt="mock" width="500" src="images/sharding.png">
</p>

## Denormalization

- Attempts to improve read performance at the expense of some write performance
- Redundant copies of the data are written in multiple tables to avoid expensive joins
- Denormalization might circumvent the need for such complex joins.
- In most systems, reads can heavily outnumber writes 100:1 or even 1000:1
    - Read resulting in a complex database join can be very expensive, spending a significant amount of time on disk operations

**Disadvantages:**

- Data is duplicated
- Constraints can help redundant copies of information stay in sync, increasing the database design's complexity.
- A denormalized database under heavy write load might perform worse than its normalized counterpart.

## SQL Tuning

- MySQL dumps to disk in contiguous blocks for fast access.
- Use `CHAR` instead `VARCHAR` for fixed-length fields
    - `CHAR` effectively allows for fast, random access, whereas with `VARCHAR`, you must find the end of a string before moving onto the next one.
- Use `TEXT` for large blocks of text such as blog posts. `TEXT` also allows for boolean searches. Using a`TEXT` field results in storing a pointer on disk that is used to locate the text block
- Use `INT` for large numbers up to 2^32 or 4 billion
- Use `DECIMAL` for currency to avoid floating point representation errors.
- Avoid storing large `BLOBS`, store the location of where to get the object instead
- `VARCHAR(255)` is the largest number of characters that can be counted in an 8 bit number, often maximizing the use of a byte in some RDBMS.
- Set the `NOT NULL` constraint where applicable to improve search performance

**Use good indices:**

- Columns that you are querying (SELECT, GROUP BY, ORDER BY, JOIN) could be faster with indices;
- Indices are usually represented as self-balancing [B-tree](https://en.wikipedia.org/wiki/B-tree) that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time.
- Placing an index can keep the data in memory, requiring more space
- Writes could also be slower since the index also needs to be updated
- When loading large amounts of data, it might be faster to disable indices, load the data, and then rebuild the indices.

**Partition tables:**

- Break up a table by putting hot spots in a separate table to help keep it in memory
