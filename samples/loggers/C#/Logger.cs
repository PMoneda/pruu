using RestSharp;

namespace Infra.Log
{
    public class PruuLog
    {
        private void Log(string key, string message, string level)
        {
            var pruuUrl = "https://pruu.herokuapp.com/log/" + key + "?level=" + level;
            var client = new RestClient(pruuUrl);
            var request = new RestRequest(Method.POST);
            request.AddParameter("text/plain", message, ParameterType.RequestBody);
            client.Execute(request);

        }
        public void LogDebug(string key, string msg)
        {
            Log(key, msg, "DEBUG");
        }

        public void LogError(string key, string msg)
        {
            Log(key, msg, "ERROR");
        }
        public void LogWarning(string key, string msg)
        {
            Log(key, msg, "WARNING");
        }

        public void LogInfo(string key, string msg)
        {
            Log(key, msg, "INFO");
        }
    }
}
