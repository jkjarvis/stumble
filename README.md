# Stumble

# How To Run
* `git clone https://github.com/jkjarvis/stumble.git`
* `cd web-service-gin`
* `go run .`

# APIs

* 1. `http://localhost:3000/users` - fetches all users.
* 2. `http://localhost:3000/matches` - fetches all matches.
* 3. `http://localhost:3000/near-matches/user/k` - fetches all users having distance smaller or equal to k related to user , where user is the id of the user and k is location difference.
* 4. `http://localhost:3000/get-similar-name/q` - fetches all users having q in their names.
