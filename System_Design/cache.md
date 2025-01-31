# Cache

<p align="center">
  <img alt="mock" width="500" src="images/caching_layer.png">
</p>

- Our Caching layer could be a set of other servers that sit alongside your application servers, whose sole function is to maintain in-memory copies of that data that's going into that database;
- One advantage is that we can resize this cache pool independently of the application pool;
- In case of distributed cache, the consistent hashing function will be used to route the request to where the data will be stored;
- Appropiate for applications with more **reads** than **writes**;
- The *expiration policy* dictates how long data is cached. Too long and your data may go stale; too short and the cache won't do much good;
- *Hotspots* can be a problem (the "celebrity problem"):
  - A cache hotspot identifies instructions that are accessed frequently.
  - It stores these instructions in the L0 cache.
  - Other instructions are stored in the L1 cache.
- Cold-start is also a problem. How do you initially warm up the cache without bringing down whatever your are caching?
  - Let's say that our cache layer goes offline and we have to restart it. All traffic will be redirected to the database. So it will take some time for the cache to warm up;
  - One way to deal with this is to have a separate procedure to warm up the cache before actually exposing it to the outside world
  - So maybe you have a process to actually artificially send traffic to the cache layer for simulated requests;
  - You might be playing back the dealys from the previous day or something;
  - And don't turn the system back on so that the cache layer is not powered until you are sure that it has been properly warmed up.

References:
- [Caching](https://www.geeksforgeeks.org/caching-system-design-concept-for-beginners/#5-types-of-cache)