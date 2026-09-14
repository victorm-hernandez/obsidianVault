# **Take Home Assignment** 

# **The Task** 

Create a **production ready** web service that combines two existing web services. 

1. Fetch a random location from https://locations.patch3s.dev/api/ 

2. Lookup up the current weather forecast for that location using the NWS API <u>https://www.weather.gov/documentation/services-web-api</u> 

3. Combine the results and return them to the user. 

# **Time Guidelines** 

We recommend you spend around 2 hours on this task. 

You should aim to at least have running code that meets the basic requirements of the task. 

**_Production ready_** is a broad goal. If you’re unable to meet it fully, then please include suMiciently detailed TODO comments in your code so another engineer could complete the task with minimal design thinking, and the result would match your vision of production readiness. 

Please let us know how much time you spent on the task when you submit your answer. 

# **Further requirements** 

The web service should be written in **Go or Java** 

service. 

The web service should remain responsive under load and be able to support multiple concurrent requests. 

archive and submitted for review by email. 

# **Example** 

# **Fetching a LOCATION** 

$ curl https://locations.patch3s.dev/api/random 

{"locations":[{"name":"Scottsdale","latitude":33.50921,"longitude":-111.89903}]} 

# **Fetching a WEATHER FORECAST - Step 1** 

$ curl -sL <u>https://api.weather.gov/points/33.50921,-111.89903 | jq '.properties.forecast'</u> 

"https://api.weather.gov/gridpoints/PSR/166,60/forecast" 

# **Fetching a WEATHER FORECAST - Step 2** 

$ curl -s <u>https://api.weather.gov/gridpoints/PSR/166,60/forecast | jq</u> 

'.properties.periods[0].detailedForecast' 

"Sunny, with a high near 108. West wind around 5 mph." 

# **Using the new web service** 

$ curl ‘http://localhost:5000’ 

"The weather in Scottsdale is: Sunny, with a high near 108. West wind around 5 mph." 

# **Goals/Evaluation** 

We ultimately want to deploy this application in production, which implies additional concerns beyond just returning correct results to a simple curl call. When we evaluate the production readiness of a service, we think about questions like: 

- How will the application handle large spikes in traMic? 

- What kind of observability exists, and is it enough to debug production issues and aid in capacity planning? 

- What steps have been taken to secure the service? Are there mitigations to common application security issues? 

- How does the application respond if external service dependencies misbehave? 

- How will the application be deployed? 

If your submission compiles, we’ll run it through a load test. 

Remember: implementation is great, but given this is a time-boxed exercise, we’ll also look for discussion of major areas you didn’t have time to implement. 

What else does production ready mean to you? 

# _v. 7 (switch to weather service)_ 

