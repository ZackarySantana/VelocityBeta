curl -X POST \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"email": "example@example.com", "d": "1"}' \
  localhost:8080/v1/admin/users