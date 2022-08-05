# JSON Web Token Microservice using Golang

User Validation with JSON Web Token<br>

Implementing Golang JWToken and JSON serialization/deserealization

This project is the continuation of [.NETCoreChallenge](https://github.com/SebastianRaiquenParisi/.NETCoreChallenge)
<br>at this moment it has no authentication, I will first create the Microservice functionalities and then send a request from the .NET project to validate logins to the system. Then this token will let the user handle data in the CRUD .NET application regarding its corresponding session.

Working as if I had a team that is currently working on the .NET side, focusing on implementing the Microservice first and then working on the .NET Project.

The register should be on this microservice, managing the users data entirely on this Microservice for rehusable puproses

Here I'll create the users json (later it will be a non relational database) from a [JSON file](https://reqres.in/api/users) but creating the passwords internally
