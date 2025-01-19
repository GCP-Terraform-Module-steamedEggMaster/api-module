# 📘 api-module
GCP Terraform enable api module repo

이 모듈은 Google Cloud API 활성화를 위한 모듈입니다. <br>
Google Cloud 프로젝트에서 여러 API를 한 번에 활성화할 수 있도록 설계되었습니다.

<br>

## 📑 **목차**
1. [모듈 특징](#모듈-특징)
2. [사용 방법](#사용-방법)
    1. [사전 준비](#1-사전-준비)
    2. [입력 변수](#2-입력-변수)
    3. [모듈 호출 예시](#3-모듈-호출-예시)
    4. [출력값 (Outputs)](#4-출력값-outputs)
    5. [지원 버전](#5-지원-버전)
    6. [모듈 개발 및 관리](#6-모듈-개발-및-관리)
3. [테스트](#terratest를-이용한-테스트)
4. [주요 버전 관리](#주요-버전-관리)
5. [기여](#기여-contributing)
6. [라이선스](#라이선스-license)

---

## 모듈 특징

- Google Cloud 프로젝트에 여러 API를 활성화.
- `disable_on_destroy` 옵션을 통해 API 라이프사이클 관리.(기본값 : false)
- 활성화된 API 목록, 고유 식별자 및 프로젝트 정보를 출력하여 인프라에서 활용 가능.

---

## 사용 방법

### 1. 사전 준비

다음 사항 확인:
1. Google Cloud 프로젝트 준비.
2. 적절한 IAM 권한 필요: `roles/serviceusage.serviceUsageAdmin` (API 활성화 권한 필수).

<br>

### 2. 입력 변수

| 변수명                        | 타입         | 필수 여부 | 기본값  | 설명                                                                 |
|-------------------------------|--------------|-----------|---------|----------------------------------------------------------------------|
| `services`                    | list(string) | ✅        | 없음    | 활성화할 GCP 서비스 이름 목록 (예: `["compute.googleapis.com"]`)      |
| `project`                     | string       | ❌        | null    | GCP 프로젝트 ID (제공되지 않으면 Provider 프로젝트 사용)             |
| `disable_on_destroy`          | bool         | ❌        | `false` | 리소스 삭제 시 서비스 비활성화 여부                                   |
| `disable_dependent_services`  | bool         | ❌        | `false` | 종속된 서비스 비활성화 여부                                           |
| `check_if_service_has_usage_on_destroy` | bool | ❌        | `false` | 리소스 삭제 시 30일 내 서비스 사용 여부 확인                          |
| `timeout_create`              | string       | ❌        | `20m`   | 리소스 생성 제한 시간                                                |
| `timeout_read`                | string       | ❌        | `10m`   | 리소스 읽기 제한 시간                                                |
| `timeout_update`              | string       | ❌        | `20m`   | 리소스 업데이트 제한 시간                                            |
| `timeout_delete`              | string       | ❌        | `20m`   | 리소스 삭제 제한 시간                                                |

---

### 3. 모듈 호출 예시

```hcl
module "api_enable" {
  source = "git::https://github.com/GCP-Terraform-Module-steamedEggMaster/api-module.git?ref=v1.0.0"

  services = [
    "compute.googleapis.com",
    "storage.googleapis.com"
  ]

  project                     = "your-gcp-project-id"
  disable_on_destroy          = false
  disable_dependent_services  = false
  check_if_service_has_usage_on_destroy = false

  timeout_create = "15m"
  timeout_read   = "5m"
  timeout_update = "15m"
  timeout_delete = "10m"
}
```

---

### 4. 출력값 (Outputs)

| 출력명               | 설명                                                 |
|----------------------|------------------------------------------------------|
| `enabled_api_services` | 활성화된 GCP API 서비스의 목록                      |
| `enabled_api_ids`      | 활성화된 GCP API 서비스의 고유 식별자 목록          |
| `enabled_api_names`    | 활성화된 GCP API 서비스 이름들                      |
| `project_id`           | 활성화된 GCP API 서비스가 속한 프로젝트 ID          |

<br>

### 5. 지원 버전

#### a. Terraform 버전
| 버전 범위 | 설명                              |
|-----------|-----------------------------------|
| `>= 1.10.3`| 최소 지원 버전                   |

#### b. Google Provider 버전
| 버전 범위 | 설명                              |
|-----------|-----------------------------------|
| `~> 6.0`  | 최소 지원 버전                   |

<br>

### 6. 모듈 개발 및 관리

#### **저장소 구조**
```
router-module/
  ├── .github/workflows/  # github actions 자동화 테스트
  ├── examples/           # 테스트를 위한 루트 모듈 모음 디렉터리
  ├── test/               # 테스트 구성 디렉터리
  ├── main.tf             # 모듈의 핵심 구현
  ├── variables.tf        # 입력 변수 정의
  ├── outputs.tf          # 출력 정의
  ├── README.md           # 문서화 파일
```

<br>

---

### Terratest를 이용한 테스트
이 모듈을 테스트하려면 제공된 Go 기반 테스트 프레임워크를 사용하세요. 아래를 확인하세요:

1. Terraform 및 Go 설치.
2. 필요한 환경 변수 설정.
3. 테스트 실행:
```bash
go test -v ./test
```

<br>

## 주요 버전 관리
이 모듈은 [Semantic Versioning](https://semver.org/)을 따릅니다.  
안정된 버전을 사용하려면 `?ref=<version>`을 활용하세요:

```hcl
source = "git::https://github.com/GCP-Terraform-Module-steamedEggMaster/api-module.git?ref=v1.0.0"
```

### Module ref 버전
| Major | Minor | Patch |
|-----------|-----------|----------|
| `1.0.0`   |    |   |

<br>

## 기여 (Contributing)
기여를 환영합니다! 버그 제보 및 기능 요청은 GitHub Issues를 통해 제출해주세요.

<br>

## 라이선스 (License)
[MIT License](LICENSE)