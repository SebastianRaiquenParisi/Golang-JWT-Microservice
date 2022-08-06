# JSON Web Token Microservice using Golang
#### Self imposed Golang challenge after applying to GRID eSports as a Jr Dev<br>
Microservice aimed to create authorization in [.NETCoreChallenge](https://github.com/SebastianRaiquenParisi/.NETCoreChallenge), although it could be used on any project that I make from now on!<br>
Working as if I had a team that is currently working on the .NET side, focusing on implementing the Microservice first and then working on the .NET Project receiving end.<br><br>
The objective of this project is to show knowledge about:<br>
* GoRoutines, pointers and structs :heavy_check_mark:
* User Validation with JSON Web Token :heavy_check_mark:
* JSON marshalization/demarshalization :negative_squared_cross_mark:
* Microservice architecture :negative_squared_cross_mark:
* Azure service bus :negative_squared_cross_mark:
* Simple Docker containers :negative_squared_cross_mark:

Currently has fixed values, now working on creating the users "database" from this [JSON user values](https://reqres.in/api/users) creating the passwords manually<br>
The next implementation shall be to create a change password page, then an email service for the same purpose<br><br>

 
  Implemented Golang JWToken for correct authorization management, tested using postman<br>

![alt text](https://github.com/SebastianRaiquenParisi/JWT-Golang-Microservice/blob/main/documentation-images/JWT-pm-login.png?raw=true)

![alt text](https://github.com/SebastianRaiquenParisi/JWT-Golang-Microservice/blob/main/documentation-images/JWT-pm-validate.png?raw=true)






