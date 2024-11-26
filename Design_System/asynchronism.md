# Asynchronism

<p align="center">
  <img alt="mock" width="500" src="images/asynchronism.png">
</p>

* Reduce request times for expensive operations that would otherwise be performed in-line;
* Also help by doing time-consuming work in advance, such as periodic aggregation of data.

## Kafka

> Is a Distributed Streaming Platform or a Distributed Commit Log

**Distributed**:
* Works as a cluster of one or more nodes that can live in different Datacenters
* It's inherently scalable, available, and fault-tolerant

**Streaming Platform:**
* Stores data as a stream of continuous records which can be processed in different methods.

**Commit Log:**
* When you push data to Kafka it takes and appends them to a stream of records, like appending logs in a log file or if you’re from a Database background like the [WAL (Write-ahead logging)](https://en.wikipedia.org/wiki/Write-ahead_logging)
* This stream of data can be Replayed or read from any point in time

### Message/Event

* Is the atomic unit of data for Kafka
* Might have an associate "Key" which is nothing but some metadata
  * Is used to determine the destination partition for message
* Events are immutable. It's impossible to change the past

### Topic

<p align="center">
  <img alt="mock" width="500" src="images/kafka-topic-partitions-producers-consumers.png">
</p>

* A topic is a log of events.
* Apache Kafka’s most fundamental unit of organization is the topic, which is something like a table in a relational database.
* You create different topics to hold different kinds of events and different topics to hold filtered and transformed versions of the same kind of event.
* Every event is appended to the end of your topic
* They aren't deleted from topic until a configurable amount of time has elapsed, even if they've been read;
* Topics are properly logs, NOT QUEUES; they are durable, replicated, fault-tolerant records of the events stored in them.
* Topics are stored as log files on disk

### Partitions

* Systematic way of breaking the one topic log file into many logs, each of which can be hosted on a separed server;
* Partition is analogous to shard in the database and is the core concept behind Kafka's scaling capabilities.

Scenario:

* Our system becomes really popular and hence there are millions of log messages per seconds;
* So now, the node on which appLogs topic is present, is unable to hold all the data that is coming in;
* We initially solve this by adding more storage to our node i.e vertical scaling;
  * But as we all know vertical scaling has its limit
* Which means we need to add more nodes and split the data between the nodes
  * When we split data of a topic into multiple stream, we call all of those streams the "Partition" of the topic

<p align="center">
  <img alt="mock" width="500" src="images/kafka">
</p>

* The blocks avobe are the different messages in that partition
* Imagine the topic to be an array, now due to memory constraint we have split the single array into 4 different smaller array;
* The numbers on the blocks in this picture denote the **Offset** the first block is at the 0th offset and the last block would on the (n-1)th offset;
* Please note that on Kafka it is not going to be an actual array but a symbolic one
