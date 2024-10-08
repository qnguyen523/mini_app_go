curl http://localhost:8080/public/login \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{"Email": "user", "Password": "password"}'

curl http://localhost:8080/public/register   \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{"Email": "user", "Name": "user", "Password": "password"}'

curl http://localhost:8080/public/user --request "GET" --data '{"jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjgyMDc0NTEsInN1YiI6IjMifQ.dfoXIpztPfwQsph_fiWDawy59zxLftmnkCV9RucdBGk"}'

1728208639