# Project: Quiz app

## End-point: Sign in

### Method: POST

> ```
> http://localhost:8000/auth/signin
> ```

### Body (**raw**)

```json
{
  "name": "user",
  "password": "qwerty123"
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Sign up

### Method: POST

> ```
> http://localhost:8000/auth/signup
> ```

### Body (**raw**)

```json
{
  "name": "user",
  "password": "qwerty123"
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get user

### Method: GET

> ```
> http://localhost:8000/api/users/123
> ```

## End-point: Delete user

### Method: DELETE

> ```
> http://localhost:8000/api/users/2
> ```

## End-point: Update user

### Method: PUT

> ```
> http://localhost:8000/api/users/1
> ```

### Body (**raw**)

```json
{
  "name": "updated",
  "password": "qwerty1234"
}
```

## End-point: Get all users

### Method: GET

> ```
> http://localhost:8000/api/users/
> ```

## End-point: Get all quizzes

### Method: GET

> ```
> http://localhost:8000/api/quizzes/
> ```

## End-point: Create new quizz

### Method: POST

> ```
> http://localhost:8000/api/quizzes/
> ```

### Body (**raw**)

```json
{
  "name": "NewQuiz"
}
```

## End-point: Get quiz by id

### Method: GET

> ```
> http://localhost:8000/api/quizzes/1
> ```

## End-point: Delete quiz by id

### Method: DELETE

> ```
> http://localhost:8000/api/quizzes/3
> ```

## End-point: Update quiz

### Method: PUT

> ```
> http://localhost:8000/api/quizzes/1
> ```

### Body (**raw**)

```json
{
  "name": "QuizNew1"
}
```

## End-point: Take a quiz

### Method: POST

> ```
> http://localhost:8000/api/quizzes/1/take
> ```

### Body (**raw**)

```json
{
  "answers": [
    {
      "question_id": 2,
      "option_id": 3
    }
  ]
}
```

## End-point: Get leaderboard

### Method: GET

> ```
> http://localhost:8000/api/quizzes/4/leaderboard
> ```

## End-point: Create question for specific quiz

### Method: POST

> ```
> http://localhost:8000/api/quizzes/3/questions
> ```

### Body (**raw**)

```json
{
  "description": "Question1"
}
```

## End-point: Get specific question for quiz

### Method: GET

> ```
> http://localhost:8000/api/quizzes/1/questions/2
> ```

## End-point: Get question by id

### Method: GET

> ```
> http://localhost:8000/api/questions/2
> ```

## End-point: Delete question

### Method: DELETE

> ```
> http://localhost:8000/api/questions/16
> ```

## End-point: Update question

### Method: PUT

> ```
> http://localhost:8000/api/questions/15
> ```

### Body (**raw**)

```json
{
  "description": "Question1New",
  "quiz_id": 1
}
```

## End-point: Get all questions for quiz

### Method: GET

> ```
> http://localhost:8000/api/quizzes/1/questions/
> ```

## End-point: Create option for question

### Method: POST

> ```
> http://localhost:8000/api/questions/2/options
> ```

### Body (**raw**)

```json
{
  "content": "correct",
  "correct": true
}
```

## End-point: Get specific option for question

### Method: GET

> ```
> http://localhost:8000/api/questions/2/options/1
> ```

## End-point: Get all options for question Copy

### Method: GET

> ```
> http://localhost:8000/api/questions/2/options
> ```

## End-point: Get specific option

### Method: GET

> ```
> http://localhost:8000/api/options/2/
> ```

## End-point: Update specific option

### Method: PUT

> ```
> http://localhost:8000/api/options/1/
> ```

### Body (**raw**)

```json
{
  "question_id": 2,
  "content": "updated option",
  "correct": false
}
```

## End-point: Delete specific option

### Method: DELETE

> ```
> http://localhost:8000/api/options/2/
> ```
