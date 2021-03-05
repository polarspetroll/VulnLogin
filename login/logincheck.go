package login
import (
  "net/http"
  "fmt"
  "database/sql"
_ "github.com/go-sql-driver/mysql"
  "html/template"
)
var DBcreds string
var Temp *template.Template
type res struct {
  Resault string
}
func LoginCheck(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open("mysql", DBcreds)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    Temp.ExecuteTemplate(w, "500.html", nil)
    fmt.Println("Cannot connect to database\n", err.Error())
    return
  }
  defer db.Close()
  Temp.ExecuteTemplate(w, "login.html", nil)

  if r.Method == "POST" {
    r.ParseForm()
    username := r.PostForm.Get("username")
    password := r.PostForm.Get("password")
    query := fmt.Sprintf("SELECT username, password FROM vulnlogin WHERE username='%v' AND password='%v'", username, password)
    q, err := db.Query(query)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      Temp.ExecuteTemplate(w, "500.html", nil)
      fmt.Fprintf(w, err.Error())
      return
    }
    defer q.Close()
    if q.Next() == true {
      stat := fmt.Sprintf("success! you logged in as %v", username)
      Temp.ExecuteTemplate(w, "done.html", stat)
    }else {
      Temp.ExecuteTemplate(w, "fail.html", "Login Failed")
    }
  }
}
