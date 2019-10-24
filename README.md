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
