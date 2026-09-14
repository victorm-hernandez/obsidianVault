## ### OWASP Password Storage Recommendations

Source: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html

- **Avoid Fast Hashing Algorithms:** Fast cryptographic hashes like MD5, SHA-1, or standard SHA-256/SHA-512 must **not** be used for password storage. Because they are designed to compute quickly, they allow attackers to run massive brute-force and dictionary guesses via GPUs.
    
- **Use Slow, Adaptive Hashing Functions:** Passwords must be protected using modern, memory-hard, or computationally intensive hashing functions. The preferred choices are:
    
    - **Argon2id** (Recommended primary choice; e.g., configured with adequate memory and iteration parameters).
        
    - **scrypt** (A strong memory-hard alternative).
        
    - **bcrypt** (A widely adopted standard, keeping in mind its 72-byte password length limitation).
        
    - **PBKDF2** (Recommended primarily if FIPS-140 compliance is required, using a high work factor/iteration count, such as 600,000+ iterations with HMAC-SHA-256).
        
- **Mandatory Unique Salting:** A unique, cryptographically strong random salt must be generated and appended to every single credential upon creation to completely mitigate rainbow table and precomputed lookup attacks.

## NIST COMPLIANCE

The document that governs password storage requirements is **NIST Special Publication (SP) 800-63B**, titled **_Digital Identity Guidelines: Authentication and Lifecycle Management_**.
### Key Password Storage Requirements in SP 800-63B:

- **Resistance to Offline Attacks:** Verifiers must store memorized secrets (passwords) in a way that protects against offline cracking if the database is compromised.
    
- **Salting and Key Derivation Functions (KDFs):** Passwords must be salted and hashed using a suitable one-way key derivation function or password hashing scheme (such as **PBKDF2**, **Argon2**, **bcrypt**, or **scrypt**) rather than standard, fast cryptographic hashes (like plain SHA-256 or MD5).
    
- **Salt Length:** The salt must be at least **32 bits** in length and chosen arbitrarily to minimize collisions among stored hashes.
    
- **Cost Factor / Iteration Count:** Algorithms must use a configurable cost factor or iteration count (e.g., iterations in PBKDF2) to intentionally slow down computation, making brute-force and dictionary guessing attacks computationally expensive for an attacker.

## Related FIPS 

While there isn't a single FIPS document that specifies how to store passwords in a web application database, **FIPS 140** (specifically **FIPS 140-2** and **FIPS 140-3**) governs the cryptographic modules and approved algorithms that federal systems and regulated industries must use.

When it comes to FIPS compliance and password hashing, the relationship works as follows:

- **The FIPS Limitation:** Modern, highly recommended password hashing algorithms like **Argon2id** and **bcrypt** are _not_ FIPS-approved algorithms. Because they are not part of the official NIST-validated cryptographic primitive suite (like SHA-2 or AES), systems operating under strict FIPS mode cannot use them out-of-the-box.
    
- **The FIPS-Approved Solution:** **PBKDF2** (Password-Based Key Derivation Function 2) _is_ a NIST-approved key derivation function under FIPS guidelines (often built using FIPS-approved underlying hash primitives like HMAC-SHA-256).
    
- **The Rule for Regulated Environments:** If an application requires strict **FIPS 140 compliance**, developers typically must use **PBKDF2** with a high iteration count (e.g., 600,000+ iterations using HMAC-SHA-256) rather than Argon2 or bcrypt.

## NIST FIPS 

Source: https://www.nist.gov/standardsgov/compliance-faqs-federal-information-processing-standards-fips

**What are Federal Information Processing Standards (FIPS)?**

FIPS are standards for federal computer systems that are developed by the National Institute of Standards and Technology (NIST) and approved by the Secretary of Commerce in accordance with the Information Technology Management Reform Act of 1996 and Computer Security Act of 1987. These standards are developed when there are no acceptable industry standards or solutions for a particular government requirement. Although FIPS are developed for use by the Federal Government, many in the private sector voluntarily use these standards.

**What are the current FIPS?**

The list of current FIPS—those that have been published, plus draft FIPS posted for comment—can be found on NIST’s [Current FIPS](https://csrc.nist.gov/publications/fips) webpage.

**Are All FIPS mandatory?**

No. FIPS are not always mandatory for federal agencies. The applicability section of each FIPS details when the standard is applicable and mandatory. FIPS do not apply to national security systems (as defined in Title III, Information Security, of the Federal Information Security Management Act (FISMA) of 2002).

State agencies administering federal programs like unemployment insurance, student loans, Medicare, and Medicaid must comply with FISMA. Private sector companies with government contracts must also comply with FISMA, which mandates the use of FIPS.

**Can federal agencies waive mandatory FIPS?**

No. The Computer Security Act of 1987 contained a waiver process for FIPS; however, this Act was superseded by FISMA of 2002 (as amended by the Federal Information Security Modernization Act (FISMA) of 2014), which no longer allows this practice. Some FIPS may still contain language referring to the “waiver process,” but this no longer valid.

**How can FIPS be used by non-government organizations?**

While FIPS are required for Federal Government organizations, the standards are valuable resources for non-government organizations looking to secure their information and systems and establish strong information security programs.

**How are FIPS developed and when are they withdrawn?**

Please visit [Procedures for Developing FIPS (Federal Information Processing Standards) Publications](https://www.nist.gov/itl/procedures-developing-fips-federal-information-processing-standards-publications) for current information on how FIPS are developed and when they are withdrawn.