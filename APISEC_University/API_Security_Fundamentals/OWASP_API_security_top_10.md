# #1 Broken Object Level Authorization (BOLA)

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

*References:*<br />
1 - [API1:2023 Broken Object Level Authorization](https://owasp.org/API-Security/editions/2023/en/0xa1-broken-object-level-authorization/)<br />
2 - [CWE-285: Improper Authorization](https://cwe.mitre.org/data/definitions/285.html)<br />
3 - [CWE-639: Authorization Bypass Through User-Controlled Key](https://cwe.mitre.org/data/definitions/639.html)

# #2 Broken Authentication

> Authentication

### Description

Authentication mechanism is an easy target for attackers since it's exposed to everyone.

Authentication will be considered weak or bad when the follwoing aspects exist:
* Weak password requirements;
* Credencial stuffing, where the attacker uses brute force with a list of valid usernames and passwords;
* Permits brute force attack on the same user account, without presenting captcha/account lockout mechanism;
* Brute forcing of IDs and passwords;
* Lack of rate limiting;
* Putting authentication info into URLs (tokens, passwords, etc), which can be sniffed and reused by malicious parties;
* Changing passwords, email address, or do any other sensitive operations without authentication or verification;
* Uses plain text, non-encrypted, or weakly hashed passwords.
* Doesn't validate the authenticity of tokens;
* Accepts unsigned/weakly signed JWT tokens (`{"alg":"none"}`);
* Doesn't validate the JWT expiration date.

On top of that, a microservice is vulnerable if:
* Other microservices can access it without authentication;
* Uses weak or predictable tokens to enforce authentication.

### Risk

It can lead to improper access to the application, consequently,the leakeage and 
changing of sensive information.

### Prevention

* Know all the possible flows to authenticate to the API (mobile/web/deep links that implement one-click authentication/etc);
* Use the standards! Don't implement an authentication mechanism from scratch. Most framework
 and libraries already provide ways to add authentication mechanism;
* Credential recovery/forgot password endpoints should be treated as login endpoints in terms of brute force, rate limiting, and lockout protections;
* Require re-authentication for sensitive operations (e.g reset password, changing the account owner email address/2FA phone number);
* Where possible, implement multi-factor authentication;
* Implement anti-brute force mechanisms to mitigate credential stuffing, dictionary attacks, and brute force attacks on your authentication endpoints;
* Implement account lockout/captcha mechanisms to prevent brute force attacks agains specific users;
* Implement weak-password checks;
* API keys should not be used for user authentication. They should only be used for API clients (Ref. 4) authentication.

*References:*<br />
1 - [API2:2023 Broken Authentication](https://owasp.org/API-Security/editions/2023/en/0xa2-broken-authentication/)<br />
2 - [CWE-204: Observable Response Discrepancy](https://cwe.mitre.org/data/definitions/204.html)<br />
3 - [CWE-307: Improper Restriction of Excessive Authentication Attempts](https://cwe.mitre.org/data/definitions/307.html)<br />
4 - [Why and When to use API Keys](https://cloud.google.com/endpoints/docs/openapi/when-why-api-key)

# #3 Broken Object Property Level Authorization

> Authorization

### Description

Exploit of endpoints by reading and/or modifying values of objects. 

An API endpoint is vulnerable if:

* The API endpoint exposes properties of an object that are considered sensive and should not be read by the user. For example: an object that represents an entity in the database (Excessive Data Exposure).

* The API endpoint allows a user to change, add/or delete the value of a sensive object's property which the user should not be able to access (Mass Assignment).

### Risk

Unauthorized access to private/sensitive object properties may result in data disclosure, data loss, or data corruption. Under certain circumstances, unauthorized access to object properties can lead to privilege escalation or partial/full account takeover.

### Prevention

* Return only minimum amount of data required for the use case;
* When exposing an object, always make sure that the user should have access to the object's properties you expose;
* Create objects containing specific properties that you want to return;
* Allow changes only to the object's properties that should be updated by the client;
* Implement a schema-based response validation mechanism as an extra layer of security. As part of this mechanism, define and enforce data returned by all API methods (e.g.: https://json-schema.org/).

*References:*<br />
1 - [API3:2023 Broken Object Property Level Authorization](https://owasp.org/API-Security/editions/2023/en/0xa3-broken-object-property-level-authorization/)<br />
2 - [CWE-213: Exposure of Sensitive Information Due to Incompatible Policies](https://cwe.mitre.org/data/definitions/213.html)<br />
3 - [CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes](https://cwe.mitre.org/data/definitions/915.html)