## Installation
- Firstly you have to set database(MySQL) address in connections/connection.go.
- After setting you can run the project


## About The Project
- When you run the project, there will be 3 tables in the database, users, questions and answers.
- First register and then login, otherwise API doesn't allow to get or post datas for user operations.
- To check errors, you can try to get user endpoints before registration
- After registration please login and try to get user by id and then get all users. If you
  write invalid user id you will get an error
- After checking user features, you can add new question and you can add all get questions
- To get special question answer you can send a question id and see what is answer 
- Then you can post all answers of questions by userid and questionid and you can see correct answers
and your rank

## EndPoints

### POST
- http://localhost:8080/api/register

/*To register new user please follow the format below*/
```
{
    "Name": "",
    "Email": "",
    "Password": ""
}
```


- http://localhost:8080/api/login

/*To login please follow the format below*/
```
{
    "Email": "",
    "Password": ""
}
```

- http://localhost:8080/api/addQuestion

/*To add new question please follow the format below*/
```
{
    "Question": "Calculate 3+2=?",
    "ChoiceA": "5",
    "ChoiceB": "4",									
    "ChoiceC": "3",
    "ChoiceD": "2",
    "Answer": "A"

}
```

- http://localhost:8080/api/getResult

/*To post answers please follow the format below*/
```
{
    "userId": 1,
    "answers": [{
        "questionId": 1,
        "chosen": "A"
    },
    {
        "questionId": 2,								//You must send as many answers as the number of questions
        "chosen": "B"
    },
    {
        "questionId": 3,
        "chosen": "C"
    },
    {
        "questionId": 4,
        "chosen": "D"
    },
    {
        "questionId": 5,
        "chosen": "A"
    }
    ]
}
```


- http://localhost:8080/api/logout/{id}

/*user is logged out and token is deleted by given id*/

### GET


- http://localhost:8080/api/getUser/{id}        /*get user by given id*/

- http://localhost:8080/api/getAllUser        /*get All Users*/

- http://localhost:8080/api/getQuestions     /*get All Questions*/

- http://localhost:8080/api/getAnswer/{id}     /*get answer by questions id*/