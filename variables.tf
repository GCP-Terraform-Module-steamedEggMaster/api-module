variable "project_id" {
  description = "GCP 프로젝트 ID"
  type        = string
}

variable "api_services" {
  description = "활성화할 API 이름 목록"
  type        = list(string)
  default     = [] # 기본값은 빈 리스트
}

variable "disable_on_destroy" {
  description = "Terraform destroy 시 API를 비활성화할지 여부"
  type        = bool
  default     = false
}
