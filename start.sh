# # Gateway 이미지 빌드
# docker build -t tss-gateway:latest -f Dockerfile.gateway .

# # Party 이미지 빌드
# docker build -t tss-party:latest -f Dockerfile.party .

# # Docker 네트워크 생성
# docker network create tss-network

# # Gateway 서버 실행
# docker run -d --name tss-gateway --network tss-network -p 8080:8080 tss-gateway:latest

# # Party 서버 실행 (3개의 파티 서버)
# docker run -d --name tss-party-1 --network tss-network -p 9091:9090 tss-party:latest
# docker run -d --name tss-party-2 --network tss-network -p 9092:9090 tss-party:latest
# docker run -d --name tss-party-3 --network tss-network -p 9093:9090 tss-party:latest


# 모든 리소스 삭제
kubectl delete -f kubernetes/

# 잠시 대기 (리소스가 완전히 삭제되기를 기다림)
sleep 10

# Docker 이미지 다시 빌드
eval $(minikube -p minukube docker-env)
docker build -t tss-gateway:latest -f Dockerfile.gateway .
docker build -t tss-party:latest -f Dockerfile.party .

# 권한
kubectl delete clusterrolebinding pod-manager-binding
kubectl delete clusterrole pod-manager
kubectl create clusterrole pod-manager --verb=get,list,watch,update --resource=pods
kubectl create clusterrolebinding pod-manager-binding --clusterrole=pod-manager --serviceaccount=default:default

# 리소스 재배포
kubectl apply -f kubernetes/gateway-deployment.yaml
kubectl apply -f kubernetes/gateway-service.yaml
kubectl apply -f kubernetes/party-deployment.yaml
kubectl apply -f kubernetes/party-hpa.yaml
kubectl apply -f kubernetes/party-service.yaml

# 배포 상태 확인
kubectl get pods
kubectl get services


# 잠시 대기 (Pod이 완전히 시작하기를 기다림)
sleep 10

kubectl port-forward service/tss-gateway-service 8080:80

