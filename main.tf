resource "google_project_service" "enabled_services" {
  # 필수 옵션
  for_each = toset(var.services) # 반복적으로 API를 생성

  # 선택적 옵션
  service                    = each.value                     # 활성화할 GCP 서비스 이름
  project                    = var.project                    # GCP 프로젝트 ID
  disable_on_destroy         = var.disable_on_destroy         # 리소스 삭제 시 서비스 비활성화 여부
  disable_dependent_services = var.disable_dependent_services # 종속된 서비스 비활성화 여부

  # Timeout 설정
  timeouts {
    create            = var.timeout_create # 생성 제한 시간
    read   = var.timeout_read   # 읽기 제한 시간
    update = var.timeout_update # 업데이트 제한 시간
    delete = var.timeout_delete # 삭제 제한 시간
  }
}
