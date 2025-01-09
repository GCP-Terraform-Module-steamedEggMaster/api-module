module "gcp_services" {
  source = "../../" # 모듈의 경로를 지정

  # 활성화할 GCP 서비스 목록
  services = [
    "compute.googleapis.com",
    "storage.googleapis.com",
    "sqladmin.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "serviceusage.googleapis.com"
  ]
}