# letz-go

### Exercise 01 : Just a simple Hello Call !

Can you create a method which will output into either console or log file the following sentence :  
Hello name !

Of course, name shall be a parameter given to the called method. 

### Exercise 02 : An External Structure to be Used !

* Create a Structure into file /internal/character/character.go
* Make it Public to the Application
* It has several fields : 
    * id (number) 
    * age (number)
    * name (string)
    * gender (string)...
* Create a function toString to prettify the way the Character object will be displayed.
    * It must be Public
    * It must be applied to a Pointer of Character Structure (so that we call it like c.toString())
    * It must return a String representation of the Structure    
* in the main.go file, create a new Character with arbitrary values and display it

### Exercise 03 : Connect to a DB using GORM - Go ORM to simplify the Code

Now, it's time to start working on concrete stuffs. First of all, we are going to connect the backend to an External PostgreSQL Server.
We use GORM, a Go ORM Framework, to simplify the way to interact with the DB.

Here are the steps to follow : 

1. Download/install GORM (go get -u github.com/jinzhu/gorm) 
2. Create a const list into /internal/database/db-params.go which must contain
    * Host ("51.254.223.8")
    * Port (5432)
    * User ("maltem-lux-go")
    * Password ("Just ask us, we will give it to you orally")
    * DbName ("letz-go")
2. Follow what's required into file : /internal/database/db-manager.go
3. into main.go file, try to connect to the database and display a Success log
