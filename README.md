# VulnLogin
Vulnerable login form
### 
this login form is vulnerable to SQLI, i made this for learning how sql injection works and learning a bit of web development.
this can be useful for beginners, since im a beginner at web development and golang so the source code written in a simple way.

### Installation

```
git clone https://github.com/polarspetroll/VulnLogin.git && cd VulnLogin
```
- create database : 

```sql
CREATE DATABASE vulnlogin;
```

```
mysql -u root -p -D vulnlogin < db.sql

go run main.go

```
---


#### environment variables : 

```DBPWD``` : **Database Password**


- **Server Address : 127.0.0.1:8080**
