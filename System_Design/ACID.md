# ACID

## Atomicity 
> Each transaction is all or nothing.

* When you send a query to your database, the entire transaction either succeeds or fails;
* Everything you want to do, must be done at once;
* If any component of write operation fails, the entire operation fails and is rolled back;

## Consistency
> All database rules are enforced, or the entire transaction is rolled back.

- If I have a rule that says a specific field can never be negative, and I try to write a negative value somehow, this transaction will be rolled back to enforce consistency in applying that rule

**OBSERVATION:** The term of Consistency outside of ACID means reading newly written data faithfully. This term can be found in the CAP Theorem.


## Isolation
> No transaction is affected by any other transaction that is still in progress.

- While one command is writing datas, another command can't change it.

## Durability
> Once a transaction has been committed, it will remain so.

- This is applies to durable storages such as disk.