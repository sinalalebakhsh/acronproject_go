# acronproject_go

-----------------------------------------------------------------------------------------------
Make Code cleaner with
package level variables
Defined at the top outside all functions
They can be accesses inside any of the functions
And in all files, which are in the same package
Belike Python language is global variable
Local Variables:
> Difined inside a function or a block
> They can be accessed only inside that function or block of code
-----------------------------------------------------------------------------------------------
More Use Cases of functions
Another important purpose which is that, same block of code can be reused in many different places in your Application
by calling the function name just like variables where we define something once and reuse it hundred times so far example
if you were hosting 10 conferences at the same time and needed to ask for and validate user input you can reuse that code
for all 10 conferences instead of writing that same login 10 times.
-----------------------------------------------------------------------------------------------
Packages in Go:
Until now we worked only with 1 file.
What if I'm writing a more complex Go application with a lot of logic?
For run this To file:
go run main.go helper.go
Or write:
go run .
-----------------------------------------------------------------------------------------------
You can export: variables, functions, constant, types
-----------------------------------------------------------------------------------------------
Scope Rules (Variable Scopes) in Go

Variable Scope Definetion:
    Scope is the region of a program, where a defined variable can be accessed


3 Levels of Scope:
    Local:
        1-Declaration within function (Or assertion Or statement):
            Can be used only within that function
        2- Declaration within block
            Can be used only within that block
    Package:
        1- Declaration outside all functions:
            Can be used everywhere in the same package
    Global:
        1-Declaration outside all functions & uppercase first letter:
            Can be used everywhere across all package

Revert "helper" package in tag 0.0.43  (Revert like = return, reflux, reversal, regression, relapse, turnover  )

--------------------------------------------------------
Maps
    > Maps unique keys to values
    > You can retrieve the value by using its key later
    > All keys have the same data type 
    > All values have the same data type
    Example, Slice Like List PYTHON:
    var my_slice []string
    Example, Map Like Dictionary PYTHON:
    var my_map map[string]string
    Or
    var my_map = make(map[string]string)

Definetion:
    > my_map["Keys"] = Values
    > my_map["Keys1"] = Values2
    > my_map["Keys3"] = Values3

For adding another type, We need to convert that to default Keys and Values Type:
    For this target, write this code:
        > strconv.FormatUint(uint64(get_age_number), 10)
        10 in the end is Decimal number
    For Memorize :
	    > user_Data["age_number"] = strconv.FormatUint(uint64(get_age_number), 10)

Map supports only 1 data type
* How can we save mixed data types values for every issues ?
--------------------------------------------------------
Struct Data Type
    - Stands for "Structure"
    - Can hold mixed data types
    - Collect different data types of data on users 
    - For Create use like this:
        type user_data struct {
            first_name string
            country_name string
            age_numner string 
        }
    --------------------------------------------------------
    - "type" statment-Custom Types
    - The type keyword creates a new type, with the name you specify.
    - In fact, you could also create a type based on every other data type like int, string etc.
    --------------------------------------------------------
    - Defining a structure:
      Mixed data type
      Defining a structure (which fields) of the User Type
    - It's like a lightweight class, which e.g doesn't support inheritance از وراثت پشتیبانی نمی کند
    --------------------------------------------------------
    - For append:
      Go suggests the field names of our User type
    --------------------------------------------------------
    When to Use “e.g.” and “i.e.” While Writing Your Paper
    We see them often in text, usually in parentheses, and we can usually figure out the context from the text before them, but what do those letters mean? The abbreviation “e.g.” stands for the Latin exempli gratia, which means “for example” or “for the sake of example.” The abbreviation “i.e.” stands for the Latin phrase id est, which means “that is to say” or “in other words.” When writing, we often use these terms like examples (e.g.) to emphasize a point or use (i.e.) to state the point in a different way without a long explanation.
--------------------------------------------------------
Go Routines - Concurrency
    - Simulate a long processing task
    - Generate and send ticket
--------------------------------------------------------
"time" - functionality for time
    - The sleep function stops or blocks the current "tread" (gorotine) execution for the defined duration.
    Handle this blocking code with Goroutines
    - "go" keyword
    - "go ..." - starts a new goroutine
    - A goroutine is a lightweight thread managed by the Go runtime 
-----------------------------------------------------------------------------------------------
Synchronizing the Goroutines:
    Waitgroup:
        * Waits for the launched goroutine to finish
        * Package "sync" provides basic synchronization functionality
        * Add: Sets the number of goroutines to wait for (increase the counter by provided number)
        * Wait: blocks until the Waitgroup counter is 0
        * Done: Decrements the Waitgroup counter by 1 so this is called by the goroutine to indicate that it's finished.
-----------------------------------------------------------------------------------------------
Comparsion to other languages:
    Concurency available in other languages 
        Java
    Why?What exactly is different?
        * Go is using, what's called a "Green thread"
        * Abstraction of an actual thread
        * Managed by the go routine, we are only interacting with these high level goroutine.
        * You can run hundreds of thousands or millions goroutines without affecting the performance.
        * Built-in functionality for goroutines to talk with one another.
        * Better connection between two or more thread goroutines.
------------------------------------------------------------------------------------------------
for add module:
1-  write in import inside "here"
Example:
import (
	"github.com/jalaali/go-jalaali"
)
2-  Command in Terminal:
       go mod tidy
3-  result:
       go: finding module for package github.com/jalaali/go-jalaali
       go: downloading github.com/jalaali/go-jalaali v0.0.0-20210801064154-80525e88d958
       go: found github.com/jalaali/go-jalaali in github.com/jalaali/go-jalaali v0.0.0-20210801064154-80525e88d958
-----------------------------------------------------------------------
for import your package after you wrote it
you must use uppercase name function(or variable etc) like this = Sina
                                                     like this = Mina
                                                     like this = PrintSina