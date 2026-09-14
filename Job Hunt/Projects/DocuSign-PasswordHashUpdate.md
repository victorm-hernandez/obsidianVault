
# Password Hash Update - Interview Project Story

## Interview Story

### 1. Context

In my architect role, I deliberately took on some of the organization's
harder security and reliability problems as hands-on engineering
projects. One example was a password-hashing migration where I was
responsible for both the architecture and direct implementation.

The authentication platform was using a password-hashing approach that
was no longer aligned with current security guidance. We had a
compliance deadline, so this was not simply a cryptography upgrade. We
had to change a security-critical primitive across a large, distributed
authentication system without breaking authentication, creating a
significant availability problem, or leaving a long tail of accounts on
the legacy scheme.

My role was to take the problem from requirements and security guidance
through algorithm selection, benchmarking, implementation, migration,
and production rollout.

### 2. The Problem

The first question was not simply, "What is the newest hashing
algorithm?"

I needed to establish what the system actually needed to comply with and
then translate that guidance into engineering requirements.

I reviewed the applicable NIST Digital Identity Guidelines, particularly
**NIST SP 800-63B**, as well as NIST guidance around password-based key
derivation and cryptographic modules. I also compared those requirements
with the **OWASP Password Storage Cheat Sheet**, which provides
practical guidance for password storage.

The important security principle was that passwords must be stored in a
form resistant to offline attacks. Modern password storage therefore
uses a password hashing/KDF scheme with a unique salt and an
intentionally expensive cost factor rather than a fast general-purpose
cryptographic hash.

NIST SP 800-63B requires salted password hashing using a suitable
password hashing scheme and specifies that the salt be at least 32 bits.
The current revision also recommends storing information identifying the
hashing scheme and cost factor so that future migrations can be
performed safely.

For environments requiring FIPS-validated cryptography, there is an
additional constraint. FIPS 140-3 addresses validation of cryptographic
modules, rather than defining a general web-application password-storage
recipe. That distinction was important when evaluating algorithm
choices.

### 3. Algorithm Analysis

I evaluated the available password-hashing approaches based on several
dimensions:

-   Security properties and resistance to offline password cracking
-   NIST guidance and applicability to our compliance requirements
-   FIPS constraints for environments where validated cryptography was
    required
-   Library and platform support
-   CPU and memory consumption
-   Authentication latency
-   Throughput under realistic concurrency
-   Operational impact during migration
-   Ability to version the stored representation for future migrations

OWASP recommends modern adaptive password hashing schemes such as
**Argon2id**, **scrypt**, **bcrypt**, and **PBKDF2**, with the exact
choice depending on the environment. In a FIPS-constrained environment,
PBKDF2 is commonly the practical choice because it can be implemented
using approved cryptographic primitives within a validated cryptographic
module.

The important architectural point was that algorithm selection could not
be made independently of the deployment environment. The strongest
algorithm on paper is not necessarily the correct choice if it cannot
satisfy the organization's cryptographic validation requirements or
introduces unacceptable production latency.

### 4. Benchmarking

Before changing production behavior, I built benchmarks around the
candidate configurations.

I measured:

-   Hashing latency
-   Verification latency
-   CPU consumption
-   Memory consumption where applicable
-   Throughput
-   Impact under concurrency
-   Effect on authentication-service latency

The goal was to find a cost factor that materially increased the cost of
offline attacks while remaining operationally safe for legitimate
authentication traffic.

This is an important distinction in password hashing: the cost factor
should not simply be "as high as possible." It needs to be tuned against
the verifier's production capacity. NIST similarly describes the cost
factor as something that should be as high as practical without
negatively impacting verifier performance.

### 5. Finding All Password Entry Points

The cryptographic change was only one part of the problem.

The organization had multiple customer-facing and internal flows that
could establish or change a password. I therefore mapped the complete
password lifecycle and identified every code path that could create,
update, verify, reset, or otherwise process password credentials.

That included:

-   New account creation
-   Password changes
-   Password resets
-   Administrative or recovery flows
-   Authentication and verification
-   Internal APIs
-   Service-to-service paths that interacted with credential state
-   Data migration and backfill jobs

This was critical because changing only the primary login path would
have created inconsistent credential formats and potentially
reintroduced the legacy algorithm through another entry point.

I introduced a versioned representation for the password hash so that
the verifier could understand which scheme and parameters were
associated with each credential.

Conceptually, the stored representation needed to communicate something
like:

``` text
algorithm/version + parameters + salt + derived hash
```

This also made the system extensible for future cryptographic
migrations.

### 6. Migration Strategy

One of the most interesting design decisions was how to migrate existing
credentials.

We considered the traditional approach of upgrading a user's password
when they successfully logged in:

``` text
User login
    |
    v
Verify legacy password
    |
    +---- success ----> Re-hash using new algorithm
                           |
                           v
                       Store new hash
```

This approach has one major operational problem: **it depends on user
activity**.

If a customer does not log in, their credential remains on the legacy
scheme indefinitely. For a large customer population, that makes it
difficult to guarantee completion by a compliance deadline.

I therefore proposed a migration that did not require the user to log
in.

The key observation was that we could transform the existing stored
credential without knowing the user's plaintext password.

### 7. The Two-Stage Hashing Migration

The migration effectively introduced a second layer of hashing around
the existing password verifier.

Conceptually:

``` text
Existing credential:

Password
   |
   v
LegacyHash(password)
   |
   v
ExistingStoredHash


Migrated credential:

Password
   |
   v
LegacyHash(password)
   |
   v
NewHash(existing_legacy_hash)
   |
   v
NewStoredHash
```

The migration process operated on the existing stored verifier rather
than requiring access to the user's plaintext password.

At authentication time, the verifier could reproduce the same
transformation:

``` text
Password supplied by user
        |
        v
LegacyHash(password)
        |
        v
NewHash(result)
        |
        v
Compare with migrated stored value
```

This allowed us to migrate inactive users as well as active users.

An important security nuance is that this is a **migration
construction**, not a claim that applying an additional generic hash
automatically makes an otherwise weak password-storage scheme secure.
The final construction, parameters, salt handling, and cryptographic
primitives still have to satisfy the applicable security requirements
and threat model.

### 8. Why I Chose This Approach

The main alternatives were:

#### Option A: Upgrade on login

**Advantages**

-   Simple conceptual model
-   The plaintext password is available at the time of migration
-   The resulting credential can immediately use the new scheme directly

**Disadvantages**

-   Inactive users never migrate
-   Migration completion becomes dependent on customer behavior
-   Difficult to guarantee compliance by a fixed deadline
-   Requires maintaining legacy verification indefinitely

#### Option B: Offline migration using the existing verifier

**Advantages**

-   Can migrate inactive users
-   Migration can be measured and controlled
-   Can be performed in batches
-   Allows us to work toward a deterministic compliance deadline

**Disadvantages**

-   Requires careful cryptographic design
-   Requires compatibility between old verification and the new
    representation
-   Requires additional logic during the transition
-   The migration must be designed so that the resulting credential
    remains resistant to offline attacks

Given the compliance deadline and the size of the customer population, I
selected the second approach.

### 9. Phased Production Rollout

I did not treat the migration as a single database operation.

We designed it as a controlled, multi-phase rollout:

``` text
Discover
   |
   v
Benchmark
   |
   v
Implement
   |
   v
Validate
   |
   v
Canary
   |
   v
Small migration batch
   |
   v
Measure health
   |
   +---- Problems ----> Pause / rollback / investigate
   |
   v
Increase batch size
   |
   v
Complete migration
```

The migration worker processed credentials in batches rather than
attempting to update the entire population at once.

We monitored both security and operational signals, including:

-   Authentication success rate
-   Authentication latency
-   Error rate
-   CPU utilization
-   Database load
-   Migration throughput
-   Number of credentials remaining on the legacy representation
-   Number of successfully migrated credentials
-   Retry/failure counts

The migration rate could then be increased progressively as the system
demonstrated that it had sufficient capacity.

This was especially important because password hashing is intentionally
computationally expensive. A migration that is secure but overwhelms the
authentication infrastructure is still an unsuccessful production
design.

### 10. Handling Failure and Rollback

Another important architectural consideration was making the migration
observable and recoverable.

The migration process was designed to be idempotent at the credential
level. A credential that had already been migrated could be detected and
skipped rather than being transformed repeatedly.

We also separated migration state from the core authentication path so
that a migration failure did not prevent customers from authenticating.

The rollout therefore had clear control points:

-   Pause migration
-   Resume migration
-   Reduce or increase batch size
-   Detect already migrated records
-   Retry failed records
-   Monitor authentication behavior independently from migration
    throughput

This made the migration an operationally controlled process rather than
a one-time data manipulation.

### 11. Security and Compliance Outcome

The project accomplished more than replacing one hashing function.

It established a repeatable password-storage architecture with:

-   Modern password hashing/KDF practices
-   Per-credential salts
-   Configurable work factors
-   Versioned credential representations
-   A migration strategy that did not depend exclusively on user login
    activity
-   Benchmark-driven configuration
-   Controlled production rollout
-   Operational telemetry
-   A path for future cryptographic migrations

Most importantly, I translated a compliance requirement into an
engineering solution.

The difficult part was not calling a different cryptographic API. The
difficult part was understanding the security requirements, choosing an
algorithm compatible with the organization's compliance environment,
finding every place where passwords entered the system, designing a
migration that could reach inactive users, and deploying it without
destabilizing the authentication platform.

------------------------------------------------------------------------

# Short Interview Version

If I only have two or three minutes, I would tell the story this way:

> One project I took on as an architect was a password-hashing
> migration. Our authentication platform was using a legacy
> password-hashing approach that was no longer aligned with current
> security guidance, and we had a compliance deadline to address it.
>
> I owned the problem end to end, including the architecture and
> hands-on implementation. I first studied the applicable NIST guidance,
> particularly SP 800-63B, and compared that with OWASP recommendations.
> I also had to understand our FIPS requirements because that affected
> which cryptographic approaches were viable.
>
> I evaluated the candidate algorithms based not only on security, but
> also on operational characteristics. I built benchmarks to measure
> hashing and verification latency, CPU impact, throughput, and
> concurrency so we could select a cost factor that made offline attacks
> expensive without creating unacceptable authentication latency.
>
> The next challenge was discovering every place in the organization
> where a password could be created, changed, reset, or verified. I
> mapped those flows and introduced a versioned credential
> representation so the system could support the new scheme and future
> migrations.
>
> The most interesting architectural decision was the migration
> strategy. We considered upgrading passwords when users logged in, but
> that would leave inactive users on the legacy scheme indefinitely.
> Because we had a compliance deadline, I designed a migration that
> could operate on the existing stored credential without requiring the
> user to log in. We used a staged hashing construction so the existing
> verifier could be transformed into the new representation, and
> authentication could reproduce the transformation when the user
> supplied their password.
>
> Finally, I treated the migration as a production distributed-system
> problem rather than a database script. We migrated credentials in
> batches, monitored authentication latency, error rates, CPU and
> database load, and progressively increased the migration rate as the
> system demonstrated capacity.
>
> The result was a controlled migration that addressed the security and
> compliance requirement while minimizing customer impact. The bigger
> lesson for me was that cryptographic migrations are rarely just
> cryptography problems. They are architecture, compatibility,
> distributed-systems, observability, and operational-rollout problems
> at the same time.

------------------------------------------------------------------------

# Technical Concepts to Be Ready to Discuss

## Password hashing vs. encryption

Password verification should generally use a one-way password
hashing/KDF scheme rather than reversible encryption.

The security goal is to make each offline password guess expensive while
avoiding storage of the plaintext password.

## Salt

A salt is a unique, random value associated with each password.

The salt prevents attackers from efficiently reusing precomputed
password guesses across accounts.

NIST SP 800-63B requires at least 32 bits of salt and requires the salt
and resulting hash to be stored with the credential.

## Work factor

The work factor controls how expensive password verification is.

The correct value is a performance/security tradeoff:

``` text
Higher cost
    |
    +--> Higher attacker cost
    |
    +--> Higher legitimate authentication cost
```

Therefore the value must be benchmarked against real production
capacity.

## Adaptive password hashing

Password hashing schemes such as Argon2id, scrypt, bcrypt, and PBKDF2
are designed to make password guessing substantially more expensive than
using a fast general-purpose hash.

OWASP currently recommends Argon2id as the primary choice when the
environment permits it and identifies PBKDF2 as an option when FIPS-140
requirements apply.

## FIPS 140-3

FIPS 140-3 is about the security requirements and validation of
cryptographic modules.

It should not be described as a password-storage standard by itself.

A more precise interview statement is:

> "We had to account for FIPS-validated cryptography in environments
> where that requirement applied, which constrained our algorithm
> choices."

## NIST SP 800-63B-4

As of 2026, the current NIST Digital Identity Guidelines are **SP 800-63
Revision 4**, finalized in July 2025. The authentication requirements
are in **SP 800-63B-4**.

The current guidance requires passwords to be stored using salted
password hashing resistant to offline attacks, with a cost factor that
is as high as practical without negatively affecting verifier
performance. It also recommends storing the password-hashing scheme and
cost factor so future migrations are possible.

## Future migration

A strong architecture should not hard-code the password algorithm into
the entire authentication stack.

Instead, treat the credential as a versioned structure:

``` text
Credential
├── scheme
├── version
├── parameters
├── salt
└── derived value
```

That allows the verifier to support:

``` text
Legacy scheme
      |
      v
Migration scheme
      |
      v
Future scheme
```

without requiring another organization-wide rewrite.

------------------------------------------------------------------------

# Interview Questions You Should Expect

### "Why couldn't you just rehash everyone's password?"

Because a secure password hash is intentionally one-way. We cannot
recover the plaintext password from the existing stored verifier.

That is why password migration requires either:

1.  Rehashing when the user supplies the password, or
2.  A carefully designed migration construction that can transform the
    existing verifier while preserving the ability to authenticate.

### "Why not just migrate users at login?"

That works well when there is no hard migration deadline. The problem is
that inactive accounts may never log in, leaving a potentially
indefinite tail of legacy credentials.

### "How did you choose the work factor?"

I benchmarked candidate configurations against production-like
authentication workloads and selected a cost that increased the
attacker's computational burden while keeping legitimate authentication
within the service's latency and capacity targets.

### "What happens if the migration job runs too aggressively?"

Password hashing consumes CPU and potentially memory. An uncontrolled
migration could compete with authentication traffic and create latency
or availability problems.

That's why the migration was batched and governed by production health
signals.

### "What did you monitor?"

At minimum:

-   Authentication success/error rate
-   Authentication latency
-   CPU utilization
-   Database utilization
-   Migration throughput
-   Migration failures
-   Remaining legacy credentials
-   Retry rate

### "What was the hardest part?"

The hardest part was not changing the hashing API. It was understanding
the entire credential lifecycle and designing a migration that satisfied
a security/compliance deadline without requiring every customer to
become active.

### "How did compliance influence engineering?"

Compliance gave us the security requirements and deadline, but
engineering still had to determine the implementation.

I translated the normative requirements into concrete engineering
constraints:

``` text
Regulatory / security guidance
          |
          v
Algorithm constraints
          |
          v
Performance requirements
          |
          v
Architecture
          |
          v
Migration strategy
          |
          v
Production rollout
```

That is the part I would emphasize in a senior/principal-level
interview.
