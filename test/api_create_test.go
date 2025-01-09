package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformModule(t *testing.T) {
	// Terraform 모듈 옵션 설정
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/api-test", // 모듈 경로
	}

	// Terraform Destroy 실행 (테스트 완료 후 정리)
	defer terraform.Destroy(t, terraformOptions)
	// TestTerraformModule이 끝나면 실행됨.
	// 에러 상황 발생 시에도 실행됨.

	// Terraform Apply 실행
	terraform.InitAndApply(t, terraformOptions)
}
