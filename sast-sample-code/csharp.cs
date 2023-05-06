using System.Web;
using System.Data.SqlClient;

public class SampleBadCode : IHttpHandler
{
  public void ProcessRequest(HttpContext context)
  {
    using (var conn = new SqlConnection("connString"))
    {
      string queryParam = HttpContext.Current.Request.Form["queryParam"];

      var sqlString = "SELECT * FROM Products WHERE CategoryID = " + queryParam;
      var command = new SqlCommand(sqlString, conn);
      command.ExecuteReader();
    }

    context.Response.Redirect(context.Request.QueryString["page"]);
    context.Response.Write("The page \"" + context.Request.QueryString["page"] + "\" was not found.");
  }

  public bool IsReusable
  {
    get { return true; }
  }
}
