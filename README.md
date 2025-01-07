# api-module
GCP Terraform enable api module repo

이 모듈은 Google Cloud API 활성화를 위한 모듈입니다. <br>
Google Cloud 프로젝트에서 여러 API를 한 번에 활성화할 수 있도록 설계되었습니다.

## 📋 **모듈 특징**

- Google Cloud 프로젝트에 여러 API를 활성화.
- `disable_on_destroy` 옵션을 통해 API 라이프사이클 관리.(기본값 : false)
- 활성화된 API 목록을 출력하여 인프라에서 활용 가능.

---

## 🔧 사용 방법

### 1. 사전 준비

다음 사항 확인:
1. Google Cloud 프로젝트 준비.
2. 적절한 IAM 권한 필요: `roles/serviceusage.serviceUsageAdmin` (API 활성화 권한 필수).

<br>

### 2. 입력 변수

| 변수명            | 타입   | 필수 여부 | 기본값             | 설명                                   |
|-------------------|--------|-----------|--------------------|----------------------------------------|
| `project_id`  | string | ✅        | 없음               | GCP 프로젝트 ID                        |
| `api_services`      | string | ✅        | 없음  | 활성화할 API 이름 목록          |
| `disable_on_destroy` | string | ❌      | `false`               | true로 설정 시, Terraform destroy 시 API 비 활성화
| | | | | (단, 의존 관계가 있는 API는 삭제되지 않음) |

<br>

### 3. 모듈 호출 예시

```hcl
module "api_enable" {
  source = "git::https://github.com/GCP-Terraform-Module-steamedEggMaster/api-module.git?ref=v1.0.0"

  project_id         = "your-gcp-project-id"
  api_services       = ["compute.googleapis.com", "storage.googleapis.com"]
  disable_on_destroy = false
}
```

<br>

### 4. 출력값 (Outputs)

| 출력명               | 설명                                    |
|----------------------|-----------------------------------------|
| `enabled_apis`  | 프로젝트에서 활성화된 API 목록 |

<br>

### 5. 지원 버전

#### a.  Terraform 버전
| 버전 범위 | 설명                              |
|-----------|-----------------------------------|
| `1.10.3`   | 최신 버전, 지원 및 테스트 완료                  |

#### b. Google Provider 버전
| 버전 범위 | 설명                              |
|-----------|-----------------------------------|
| `~> 4.0`  | 최소 지원 버전                   |


<br>

### 6. 모듈 개발 및 관리

- **저장소 구조**:
  ```
  api-module/
    ├── main.tf        # 모듈의 핵심 구현
    ├── variables.tf   # 입력 변수 정의
    ├── outputs.tf     # 출력 정의
    ├── README.md      # 문서화 파일
    ├── test/          # 테스트 구성 디렉터리
  ```
<br>

---

### Terratest를 이용한 테스트
이 모듈을 테스트하려면 제공된 Go 기반 테스트 프레임워크를 사용하세요. 아래를 확인하세요:

1. Terraform 및 Go 설치.
2. 필요한 환경 변수 설정 (GCP_PROJECT_ID, API_SERVICES 등).
3. 테스트 실행:
```bash
go test -v ./test
```

<br>

## 주요 버전 관리
이 모듈은 [Semantic Versioning](https://semver.org/)을 따릅니다.  
안정된 버전을 사용하려면 `?ref=<version>`을 활용하세요:

```hcl
source = "git::https://github.com/GCP-Terraform-Module-steamedEggMaster/provider-module.git?ref=v1.0.0"
```

<br>

## 기여 (Contributing)
기여를 환영합니다! 버그 제보 및 기능 요청은 GitHub Issues를 통해 제출해주세요.

<br>

## 라이선스 (License)
[MIT License](LICENSE)