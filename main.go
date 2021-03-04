package main
import(
	"./login"
	"net/http"
	"html/template"
	"fmt"
)
const DBpassword = "" // database password for root user

func main() {
	login.DBcreds = fmt.Sprintf("root:%v@tcp(127.0.0.1:3306)/vulnlogin", DBpassword)
	login.Temp = template.Must(template.ParseGlob("templates/*.html"))
	statics := http.FileServer(http.Dir("./statics"))
	http.Handle("/statics/", http.StripPrefix("/statics/", statics))
	http.HandleFunc("/", RedirectLogin)
	http.HandleFunc("/login", login.LoginCheck)
	http.ListenAndServe(":8080", nil)
}

func RedirectLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 302)
}
