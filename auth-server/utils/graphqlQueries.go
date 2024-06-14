package utils

var UserQuery string = `
    query UserQuery($username: User_Varchar) {
      user_users(where: {username: {_eq: $username}}) {
        id
        username
        password
      }
    }
	`
