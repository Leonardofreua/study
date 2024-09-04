# #1 Broken Object Level Authorization (BOLA)

> Authorization

*Attacker authenticates as User A and then retrieves data on User B.*

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

# #4 Unrestricted Resource Consumption

Abuse of APIs due to high volumes of API calls, large requests, etc. Multiple concurrent
requests performed from a single local computer or by using cloud computing resources. Most of the automated tools available are designed to cause DoS via high loads of traffic, impacting APIs service rate.

An API will be considerd vulnerable if at least one of the following limits is missing or set inappropriately (e.g. too low/high):

* Execution timeouts;
* Maximum allocable memory;
* Maximum number of files descriptors;
* Maximum number of processos;
* Maximum upload file size;
* Number of operations to perform in a single API client request (e.g. batches);
* Number of records per page to return in a single request-response;
* Third-party service providers' speding limit.

### Risks

* Daniel of Service (DoS);
* Operation costs due to higher CPU demand, increasing cloud storage needs;
* Performance Impact;
* Mass data hervesting.

### Prevention

* Implement traffic controls;
* Test effectiveness of controls
* Use containers / Serverless code (e.g. Lambdas) to limit memory, CPU, number of restarts, file descriptors and processes;
* Define and enforce a maximum size of data on all incoming parameters and payloads, such as
  maximum length for strings, numbers of elements in arrays, and maximum upload file size (regardless of whether it is stored locally or in cloud storage);
* Implement a limit on how often a client can interact with the API within a defined timeframe (RATE LIMITING);
* Rate limiting should be fine tuned based on the business needs. Some API endpoints might require stricter policies;
* Limit/throttle how many times or how often a single API client/user can execute a single operation (e.g. validate an OTP, or request password recovery without visiting the one-time URL);
* Add proper server-side validation for query string and request body parameters, specifically the one that controls the number of records to be returned in the response;
* Configure spending limits for all service providers/API integrations. When setting spending limits is not possible, billing alerts should be configured instead.

*References:*<br />
1 - [API4:2023 Unrestricted Resource Consumption](https://owasp.org/API-Security/editions/2023/en/0xa4-unrestricted-resource-consumption/)<br />
2 - [CWE-770: Allocation of Resources Without Limits or Throttling](https://cwe.mitre.org/data/definitions/770.html)<br />
3 - [CWE-400: Uncontrolled Resource Consumption](https://cwe.mitre.org/data/definitions/400.html)<br />
4 - [CWE-799: Improper Control of Interaction Frequency](https://cwe.mitre.org/data/definitions/799.html)

# #5 Broken Function Level Authorization

Abuse of API functionality to improperly modify objects (create, update, delete).

Often involves replaceing passive methods (GET) with active (PUT, DELETE).

An API will be vulnerable if regular users with low-level roles can access resources owned by users with higher roles.

### Risks

Administrative functions are keys to targets for this type of attack.

* May be used to escalate privilege;
* Can be exploited to modify account details;
* May lead to data disclosure;
* Data loss;
* Data corruption;
* Service disruption.

### Prevention

* Identify functions that expose high sensitivity capability and develop controls to limit access;
* Implement continuous release testing to ensure proper behavior;
* The enforcement mechanism(s) should deny all access by default, requiring explicit grants to specific roles for access to every function;
* Make sure that all of your administrative controllers inherit from an administrative abstract controller that implements authorization checks based on the user's group/role.

*References:*<br />
1 - [API5:2023 Broken Function Level Authorization](https://owasp.org/API-Security/editions/2023/en/0xa5-broken-function-level-authorization/)<br />
2 - [CWE-285: Improper Authorization](https://cwe.mitre.org/data/definitions/285.html)

# #6 Unrestricted Access to Sensitive Business Flows

Abuse of a legitimate business workflow through excessive, automated use Rate limiting, captchas not always effective agains fraudulent trafic.

Rapid IP rotation makes detection difficult.

Typically a result of application logic flaw.

Exploitation usually involves understanding the business model backed by the API, finding sensitive business flows, and automating access to these flows, causing harm to the business.

### Risks

* Loss of critical business activity;
* Prevent legitimate users from purchasing a product (e.g. automated ticket purchasing):
  * An attacker can buy all the stock of a high-demand item at once and resell for a higher price (scalping)
* Creating a comment/post flow:
  * an attacker can spam the system
* making a reservation:
  * an attacker can reserve all the available time slots and prevent other users from using the system
* Lead to inflation in the internal economy of a game.

An API Endpoint is vulnerable if it exposes a sensitive business flow, without appropriately restricting the access to it.

### Prevention

* Implement fraudulent traffic detection and control:
  * Using either captcha or more advanced biometric solutions (e.g. typing patterns)
  * Analyze the user flow to detect non-human patterns (e.g. the user accessed the "add to cart" and "complete purchase" functions in less than one second)
  * Consider blocking IP addresses of Tor exit nodes and well-known proxies
* Setup and automate testing of control mechanisms;
* Identify the business flows that might harm the business if they are excessively used.

*References:*<br />
1 - [API6:2023 Unrestricted Access to Sensitive Business Flows](https://owasp.org/API-Security/editions/2023/en/0xa6-unrestricted-access-to-sensitive-business-flows/)

# #7 Server Side Request Forgery

*Exploiting URL inputs to make a request a malicious servers.*

Exploitation requires the attacker to find an API endpoint that access a URI that's provided by the client. 

* SSRF occur when an API is fetching a remote resource without validating the user-supplied URL. It enables an attacker to coerce the application to send a crafted request to an unexpected destination, even when protected by a firewall or a VPN;
* More common: The following concepts encourage developers to access an external resource based on user input: Webhooks, file fetching from URLs, custom SSO, and URL previews;
* More Dangerous: Modern technologies like cloud providers, Kubernetes, and Docker expose management and control channels over HTTP on predictable, well-known paths. Those channels are an easy target for an SSRF attack;
* Local file injection (LFI) can lead to a SSRF.

### Risks

* Might lead to internal services enumeration (e.g. port scanning);
* Information disclosure;
* Bypassing firewalls or other security mechanisms;
* In some cases, it can lead to DoS or the server being used as a proxy to hide malicious activities.

### Prevention

* Validate and sanitize ALL user-supplied information, including URL parameters;
* Ensure communication only permitted with trusted resources;
* Test URL validation effectiveness;
* Whenever possible, use allow lists of:
  * Remote origins users are expected to download resources from (e.g. Google Drive, Gravatar, etc.)
  * URL schemes and ports
  * Accepted media types for a given functionality
* Disable HTTP redirections;
* Use a well-tested and maintained URL parser to avoid issues caused by URL parsing inconsistencies.
* Do not send raw responses to clients.


*References:*<br />
1 - [API7:2023 Server Side Request Forgery](https://owasp.org/API-Security/editions/2023/en/0xa7-server-side-request-forgery/)<br />
2 - [CWE-918: Server-Side Request Forgery (SSRF)](https://cwe.mitre.org/data/definitions/918.html)

# #8 Security Misconfiguration

This is a pretty broad category that basically encompasses all kinds of vulnerabilities that can result from misconfiguration, whether it's on the servers, the infrastructure the network, the application themselves. It could be unpatched systems or services that shouldn't be running, common endpoints, services running with insecure default configurations, or unprotected files and directories to gain unauthoerized access or knowledge of the system.

Missing security patches on libraries and applications used within our service. Unnecessary features enabled, no need to have SFTP enabled on a server if that's not a requirement for the application to run. Missing encryption, transport layer security. You want to make sure that your APIs can only be accessed by properly identified and approved sources. So if you've got missing CORS (Cross Origin Resource Sharing) policies, that's another potential source of a vulnerability.

* Appropriate security hardening is missing across any part of the API stack, or if there are improperly configured permissions on cloud services;
* The latest security patches are missing, or the systems are out of date;
* Unnecessary features are enabled (e.g. HTTP verbs, logging features);
* There are discrepancies in the way incoming requests are processed by servers in the HTTP server chain;
* Transport Layer Security (TLS) is missing;
* Security or cache control directives are not sent to clients;
* A Cross-Origin Resource Sharing (CORS) policy is missing or improperly set;
* Error messages include stack traces, or expose other sensitive information.

### Risks

* Expose sensitive data;
* Expose system details that can lead to ful server compromise.

### Prevention

* Implement hardening procedures;
* Routinely review configurations;
* Implement automated, continuos security testing;
* Be specific about which HTTP verbs each API can be accessed by: all other HTTP verbs should be disabled (e.g. HEAD);
* APIs expecting to be accessed from browser-based clients (e.g., WebApp front-end) should, at least:
    * implement a proper Cross-Origin Resource Sharing (CORS) policy
    * include applicable Security Headers
* Restrict incoming content types/data formats to those that meet the business/ functional requirements.
* Where applicable, define and enforce all API response payload schemas, including error responses, to prevent exception traces and other valuable information from being sent back to attackers.

*References:*<br />
1 - [API8:2023 Security Misconfiguration](https://owasp.org/API-Security/editions/2023/en/0xa8-security-misconfiguration/)<br />
2 - [CWE-2: Environmental Security Flaws](https://cwe.mitre.org/data/definitions/2.html)<br />
3 - [CWE-16: Configuration](https://cwe.mitre.org/data/definitions/16.html)<br />
4 - [CWE-209: Generation of Error Message Containing Sensitive Information](https://cwe.mitre.org/data/definitions/209.html)<br />
5 - [CWE-319: Cleartext Transmission of Sensitive Information](https://cwe.mitre.org/data/definitions/319.html)<br />
6 - [CWE-388: Error Handling](https://cwe.mitre.org/data/definitions/388.html)<br />
7 - [CWE-444: Inconsistent Interpretation of HTTP Requests ('HTTP Request/Response Smuggling')](https://cwe.mitre.org/data/definitions/444.html)<br />
8 - [CWE-942: Permissive Cross-domain Policy with Untrusted Domains](https://cwe.mitre.org/data/definitions/942.html)<br />