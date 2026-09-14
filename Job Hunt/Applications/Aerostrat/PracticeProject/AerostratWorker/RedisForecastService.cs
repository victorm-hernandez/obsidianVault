namespace AerostratWorker;
using StackExchange.Redis;

public class RedisForecastService
{
    private ConnectionMultiplexer redisMultiplexer;

    public RedisForecastService(ConnectionMultiplexer redisMultiplexer)
    {
        this.redisMultiplexer = redisMultiplexer;    
    }

    public List<Forecast> GetForecast()
    {
        var result = new List<Forecast>();
        var redis = this.redisMultiplexer.GetDatabase();
        redis.StringSet("MyTestKey","Hello world");
        
        return result;
    }

    public void SetForecast(List<Forecast> forecastData)
    {
        
    }
}