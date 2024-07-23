<!-- tss_project/
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
        최종적으로 시스템을 프로덕션 환경에 배포합니다. -->
