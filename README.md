go run gateway/cmd/main.go





// party
docker build -t gino0/tss-party:latest .
docker push gino0/tss-party:latest







kubectl delete pod -l app=tss-party

kubectl describe pod tss-party-0

kubectl logs -f tss-party-0


curl -X POST http://localhost:8080/keygen -H "Content-Type: application/json" -d '{"n": 1, "m": 2}'