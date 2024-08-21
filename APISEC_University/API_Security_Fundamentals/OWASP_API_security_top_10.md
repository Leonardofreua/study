# Broken Object Level Authorization (BOLA)

> Authorization

*Attacker authenticates as User A and then retrieves data on User B.*

### Description

Most common and damaging API vulnerability.

Manipulation of APIs to access data/objects belonging to other user. Usually through
manipulating the ID of an object that is sent within the requets. Objects Ids can be anything
from sequencial integers, UUIDs, or generic strings

Can be explore through path or query string parameters, request headers, or event 
as part of the request payload (POST)

### Risk

Can lead to data loss, data disclosure to unauthorized parties, or data manipulation. Sometimes, unauthorized access to objects can also lead to full account takeover.

### Prevention

- Define data access policies and implement associated controls;
- Enforce data access controls at application logic layer. Check if the logged in user has access to perform the requested action on the record in every function that uses an input from the client to access a record in the database;
- Use random and unpredictable values as GUIDs for record's IDs;
- Write automated testing to find BOLA flaws.

Avoid:
- Compare user ID extracted from JWT token with the vulnerable ID parameter.
