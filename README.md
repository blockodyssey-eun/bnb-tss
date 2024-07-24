tss_project/
├── cmd/
│   ├── gateway/
│   │   └── main.go
│   └── party/
│       └── main.go
├── internal/
│   ├── gateway/
│   │   └── server.go
│   ├── party/
│   │   └── server.go
│   └── tss/
│       └── tss.go
├── handler/
│   ├── gateway/
│   │   ├── keygen_handler.go
│   │   └── sign_handler.go
│   ├── party/
│   │   ├── keygen_handler.go
│   │   └── sign_handler.go
├── proto/
│   └── tss.proto
├── kubernetes/
│   ├── gateway-deployment.yaml
│   ├── gateway-service.yaml
│   ├── party-deployment.yaml
│   ├── party-service.yaml
│   └── party-hpa.yaml
├── Dockerfile.gateway
├── Dockerfile.party
├── go.mod
└── go.sum


이 구조에 대한 설명:
cmd/: 메인 애플리케이션 엔트리 포인트를 포함합니다.
    gateway/main.go: Gateway 서버의 메인 함수
    party/main.go: Party 서버의 메인 함수

internal/: 내부 패키지를 포함합니다.
    gateway/server.go: Gateway 서버 로직
    party/server.go: Party 서버 로직
    tss/tss.go: TSS 관련 핵심 로직

handler/: HTTP 및 gRPC 요청을 처리하는 핸들러를 포함합니다.
    gateway/: Gateway 서버의 핸들러
        keygen_handler.go: 키 생성 요청 핸들러
        sign_handler.go: 서명 요청 핸들러
    party/: Party 서버의 핸들러
        keygen_handler.go: 키 생성 요청 핸들러
        sign_handler.go: 서명 요청 핸들러
proto/: gRPC 프로토콜 정의 파일을 포함합니다.
    tss.proto: TSS 관련 gRPC 서비스 정의

kubernetes/: Kubernetes 배포 관련 YAML 파일들을 포함합니다.
    gateway-deployment.yaml: Gateway 서버 배포 설정
    gateway-service.yaml: Gateway 서버 서비스 설정
    party-deployment.yaml: Party 서버 배포 설정
    party-service.yaml: Party 서버 서비스 설정
    party-hpa.yaml: Party 서버의 Horizontal Pod Autoscaler 설정
    
Dockerfile.gateway  Gateway서버를 위한 Docker 이미지 빌드 파일
Dockerfile.party: Party 서버를 위한 Docker 이미지 빌드 파일

go.mod Go 모듈 의존성 관리 파일
go.sum: Go 모듈 의존성 관리 파일


1. 프로젝트 초기 설정
    프로젝트 디렉토리 생성 및 초기화
        tss_project 디렉토리를 생성하고 go mod init tss_project 명령어로 Go 모듈을 초기화합니다.
    의존성 설치
        필요한 패키지들을 go get 명령어로 설치합니다. 예: go get google.golang.org/grpc, go get github.com/bnb-chain/tss-lib, go get github.com/gorilla/mux 등.
2. 프로토콜 버퍼 정의 및 gRPC 설정
    proto 파일 작성
        proto/tss.proto 파일을 작성하여 gRPC 서비스와 메시지를 정의합니다.
    프로토콜 버퍼 컴파일
        protoc 명령어를 사용하여 Go 코드를 생성합니다. 예: protoc --go_out=. --go-grpc_out=. proto/tss.proto
3. 내부 로직 구현
    TSS 로직 구현
        internal/tss/tss.go 파일에 TSS 관련 핵심 로직을 구현합니다.
4. 핸들러 구현
    Gateway 핸들러 구현
        handler/gateway/keygen_handler.go 및 handler/gateway/sign_handler.go 파일에 HTTP 요청을 처리하는 핸들러를 구현합니다.
    Party 핸들러 구현
        handler/party/keygen_handler.go 및 handler/party/sign_handler.go 파일에 gRPC 요청을 처리하는 핸들러를 구현합니다.
5. 서버 설정 및 구현
    Gateway 서버 구현
        internal/gateway/server.go 파일에 Gateway 서버 로직을 구현합니다.
        cmd/gateway/main.go 파일에 Gateway 서버의 메인 함수를 작성합니다.
    Party 서버 구현
        internal/party/server.go 파일에 Party 서버 로직을 구현합니다.
        cmd/party/main.go 파일에 Party 서버의 메인 함수를 작성합니다.
6. Docker 설정
    Dockerfile 작성
        Dockerfile.gateway 및 Dockerfile.party 파일을 작성하여 각각의 서버를 위한 Docker 이미지를 빌드할 수 있도록 합니다.
    Docker 이미지 빌드
        docker build -t tss_project/gateway -f Dockerfile.gateway .
        docker build -t tss_project/party -f Dockerfile.party .
7. Kubernetes 설정
    Kubernetes 배포 파일 작성
        kubernetes/gateway-deployment.yaml, kubernetes/gateway-service.yaml, kubernetes/party-deployment.yaml, kubernetes/party-service.yaml, kubernetes/party-hpa.yaml 파일을 작성하여 Kubernetes에 배포할 수 있도록 합니다.
    Kubernetes 클러스터에 배포
        kubectl apply -f kubernetes/ 명령어를 사용하여 Kubernetes 클러스터에 배포합니다.
8. 테스트 및 디버깅
    단위 테스트 작성
        각 컴포넌트에 대한 단위 테스트를 작성합니다.
    통합 테스트
        시스템 전체의 통합 테스트를 수행하여 모든 컴포넌트가 올바르게 작동하는지 확인합니다.
    디버깅 및 최적화
        발견된 문제를 해결하고 성능을 최적화합니다.
9. 문서화 및 배포
    문서화
        프로젝트의 README.md 파일을 작성하여 프로젝트 개요, 설치 방법, 사용 방법 등을 문서화합니다.
    배포
        최종적으로 시스템을 프로덕션 환경에 배포합니다.


# Gateway 이미지 빌드
docker build -t tss-gateway:latest -f Dockerfile.gateway .

# Party 이미지 빌드
docker build -t tss-party:latest -f Dockerfile.party .

# Docker 네트워크 생성
docker network create tss-network

# Gateway 서버 실행
docker run -d --name tss-gateway --network tss-network -p 8080:8080 tss-gateway:latest

# Party 서버 실행 (3개의 파티 서버)
docker run -d --name tss-party-1 --network tss-network -p 9091:9090 tss-party:latest
docker run -d --name tss-party-2 --network tss-network -p 9092:9090 tss-party:latest
docker run -d --name tss-party-3 --network tss-network -p 9093:9090 tss-party:latest


# 테스트 
curl -X POST http://localhost:8080/keygen \
-H "Content-Type: application/json" \
-d '{
  "threshold": 2,
  "total_parties": 3
}'


#### 쿠버네티스 명령어

1. POD 삭제
```bash
kubectl delete pod <pod-name>
# 강제 
--grace-period=0 --force
```



# Troubleshooting

Failed to create party pods: failed to create deployment: deployments.apps is forbidden: User "system:serviceaccount:default:default" cannot create resource "deployments" in API group "apps" in the namespace "default"


1. 적절한 권한을 가진 ClusterRole을 생성합니다:
```bash
kubectl create clusterrole deployment-manager --verb=create,delete,get,list,update,watch --resource=deployments
// clusterrole.rbac.authorization.k8s.io/deployment-manager created
```


2.  ClusterRole을 default 서비스 계정에 바인딩합니다:
```bash
kubectl create clusterrolebinding deployment-manager-binding --clusterrole=deployment-manager --serviceaccount=default:default
//clusterrolebinding.rbac.authorization.k8s.io/deployment-manager-binding created
```



Failed to create party pods: failed to mark pod as in use: pods "tss-party-5fb595b4f7-g2ffz" is forbidden: User "system:serviceaccount:default:default" cannot update resource "pods" in API group "" in the namespace "default"

1. 이미 권한이 존재하는 경우 삭제 후 재생성
```bash
kubectl delete clusterrolebinding pod-manager-binding
kubectl delete clusterrole pod-manager
kubectl create clusterrole pod-manager --verb=get,list,watch,update --resource=pods
kubectl create clusterrolebinding pod-manager-binding --clusterrole=pod-manager --serviceaccount=default:default
```



### 

네, 모든 pods를 삭제하겠습니다. 다음 명령어를 사용하여 모든 pods를 삭제할 수 있습니다:

```bash
kubectl delete pods --all
```

하지만 Deployment나 ReplicaSet에 의해 관리되는 pods는 삭제 후 자동으로 다시 생성될 수 있습니다. 따라서 모든 리소스를 완전히 제거하려면 다음 명령어들을 순서대로 실행하는 것이 좋습니다:

1. Deployments 삭제:
```bash
kubectl delete deployments --all
```

2. ReplicaSets 삭제:
```bash
kubectl delete replicasets --all
```

3. Services 삭제:
```bash
kubectl delete services --all
```

4. Pods 삭제:
```bash
kubectl delete pods --all
```

5. HorizontalPodAutoscalers 삭제:
```bash
kubectl delete hpa --all
```

이렇게 하면 모든 관련 리소스가 삭제되고, 새로운 상태에서 다시 시작할 수 있습니다. 삭제 후에는 다음 명령어로 모든 리소스가 제거되었는지 확인할 수 있습니다:

```bash
kubectl get all
```
