resource "google_project_service" "enabled_services" {
  for_each           = toset(var.api_services) # 활성화할 API 리스트
  project            = var.project_id          # 프로젝트 ID
  service            = each.value              # 활성화할 API 이름
  disable_on_destroy = var.disable_on_destroy  # 삭제 시 API 비활성화 여부
}
## Terraform이 GCP API 활성화 시에 roles/serviceusage.serviceUsageAdmin 권한 필요