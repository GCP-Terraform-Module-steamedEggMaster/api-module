# 필수 변수
variable "services" {
  description = "활성화할 GCP 서비스들의 목록 (예: [compute.googleapis.com, storage.googleapis.com])"
  type        = list(string)
}

# 선택적 변수
variable "project" {
  description = "GCP 프로젝트 ID (선택 사항, 제공되지 않으면 Provider 프로젝트 사용)"
  type        = string
  default     = null
}

variable "disable_on_destroy" {
  description = "리소스 삭제 시 서비스 비활성화 여부 (기본값: false)"
  type        = bool
  default     = false
}

variable "disable_dependent_services" {
  description = "종속된 서비스 비활성화 여부 (기본값: false)"
  type        = bool
  default     = false
}

# Timeout 변수
variable "timeout_create" {
  description = "리소스 생성 제한 시간"
  type        = string
  default     = "20m"
}

variable "timeout_read" {
  description = "리소스 읽기 제한 시간"
  type        = string
  default     = "10m"
}

variable "timeout_update" {
  description = "리소스 업데이트 제한 시간"
  type        = string
  default     = "20m"
}

variable "timeout_delete" {
  description = "리소스 삭제 제한 시간"
  type        = string
  default     = "20m"
}
