# JSON Web Token Microservice using Golang
#### Self imposed Golang challenge after applying to GRID eSports as a Jr Dev<br>
Microservice aimed to create authorization in [.NETCoreChallenge](https://github.com/SebastianRaiquenParisi/.NETCoreChallenge), although it could be used on any project that I make from now on!<br>
Working as if I had a team that is currently working on the .NET side, focusing on implementing the Microservice first and then working on the .NET Project receiving end.
The objective of this project is to show knowledge about:<br>
* GoRoutines, pointers and structs :heavy_check_mark:
* User Validation with JSON Web Token :heavy_check_mark:
* JSON marshalization/demarshalization :writing_hand:
* Microservice architecture :writing_hand:
* Azure service bus :writing_hand:

Currently has fixed values, now working on creating the users "database" json (later it will be a non relational database) from this [JSON user values](https://reqres.in/api/users) creating the passwords manually<br>
The next implementation shall be to create a change password page, then an email service for the same purpose<br><br>
* *Implemented Golang JWToken for correct authorization management*<br><br>

![alt text](https://github.com/SebastianRaiquenParisi/JWT-Golang-Microservice/blob/main/documentation-images/JWT-pm-login.png?raw=true)

![alt text](https://github.com/SebastianRaiquenParisi/JWT-Golang-Microservice/blob/main/documentation-images/JWT-pm-validate.png?raw=true)





