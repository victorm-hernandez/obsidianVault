namespace AerostratWorker;

public record Forecast
{
    public float Longitude {get; init;}
    public float Lattitude {get; init;}
    public required string LocationName {get; init;}
    public required string ForecastText {get; init;}
    public DateTime ExpirationData{get; init;}
}