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


### Exercise 04 : Implementing a GET Mapping Controller - No Use of Framework

First of all, you can launch this URL : https://maltem-lux.github.io/letzgo-front/#/home
You will quickly notice that it tries to connect to your localhost:10000 URL (which is the Backend GO URL)

So the objective of this exercise is to create into the backend a GET Mapping Controller to display the list of characters into this main screen.

Creation of the Controller, DAO and Handler 
1. Into /pkg/controllers/controllers.go file, 
    1. We need to add the route to the URI and the method called once the URI is reached
    2. We also need to find the way to launch the go program as a Daemon listening on Port 10000
2. Into /internal/character/character-dao.go file, create a Method which will retrieve the list of Characters of a given Player id
3. Into /internal/character/character-controller.go file, create the handler for the route we have just created 

### Exercise 05 : Implementing a GET Mapping Controller with ID as parameter - No Use of Framework

We will reuse what we have just done before, but this time, we must display the details of a selected char.
In the front, once you select a char, you go to his/her details (by now, there is no backend for it, so nothing really correct is displayed)  
We are going to :

Adapt the Route to /characters so that the method called retrieves the ID into param
1. Tips : URI called is : /characters?charId=1
2. Param to retrieve
3. Adapt the existing handler to : 
    1. Either retrieve all the Characters
    2. Or get only details on the given charId

### Exercise 06 : Implementing a POST Mapping Controller to create a Character

On the home page of the front, it is possible to create a Char.
But the POST Controller has not been developed yet.

1. Add the route to map to the char creation. Route is "/newCharacter"
2. Create the method which does the POST (route) in the character-controller.go
3. Call the GORM method to insert into DB the new Character (character-dao.go)
