Its important to show we can reflect and accept our own mistakes. Being able to collaborate and be humble.

This is the list of problems found on the code so far:

- It doesn't use the API response data to determine the expiration of records and instead it uses a fixed 10 min age period. This is wasteful.
- The cache is applied to the entire fetch operation when the only problematic API is the first one (random location), having a multi tier cache would improve on this, cache the entire thing, but also cache the problematic API first, only use the higher level cache if the first level fails.
- The circular log length variable does fix the eternal loop problem when every record is expired but it doesnt recover well, a cursor that is X positions from the write cursor will not get there until the buffer is full again, it should get there as soon as there is a single fresh record. This is a bug, we need to cycle completely until we find write again. This can be done by counting the number of spaces (using length) and if we traverse the entire thing once (length) then move the read pointer to write instead of leaving it there.
- 