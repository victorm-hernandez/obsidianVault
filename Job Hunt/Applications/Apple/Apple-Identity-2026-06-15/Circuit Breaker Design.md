
How many failures before we stop calling the underlying service?
	10 Failures / 1 sec

How to block workers?
	- Location Worker, check circuit breaker is opened. If closed perform the request.
		- If there is an error increase error count in the circuit breaker, include retry after if needed. 
	- Circuit Breaker
		- Inform of Error:
			- Increment count of errors
			- If retry header is present, save this, change state to opened regardless of error count
			- If state is half opened OR ErrorCount > threshold, change to opened
		- Check if allowed
			- If closed, return true
			- If half opened, return false
			- If opened, check cooldown period, if time since last tripped is bigger than cooldown, allow 1 routine (with a compare and swap) and change state to half open.
		- Leak errors over time
		- Inform Of Success, set error count to 0, if opened / half opened, change to closed
