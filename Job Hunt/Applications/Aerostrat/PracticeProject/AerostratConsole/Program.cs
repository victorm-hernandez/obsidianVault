namespace AerostratConsole;
using StackExchange.Redis;
using StackExchange.Redis.Availability;

class Program
{
    static void Main(string[] args)
    {
        ConnectionMultiplexer redis = ConnectionMultiplexer.Connect("localhost:6379");
        var defaultDB = redis.GetDatabase();
        defaultDB.StringSet("TestStringKey", "Hello World");
        var readBack = defaultDB.StringGet("TestStringKey");

        Console.WriteLine($"Value read from redis is {readBack}");
    }
}
