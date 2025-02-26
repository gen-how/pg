package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ProjectConfig 包含生成專案所需的配置信息
type ProjectConfig struct {
	Name     string
	Language string
	IsApp    bool
	IsLib    bool
	UseMake  bool
	UseCMake bool
	License  string
}

// Generate 根據配置生成專案
func Generate(config ProjectConfig) error {
	// 創建專案根目錄
	err := os.MkdirAll(config.Name, 0755)
	if err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// 創建基本目錄結構
	dirs := []string{"src"}

	// 如果是庫專案，需要創建 include 目錄
	if config.IsLib {
		dirs = append(dirs, "include")
	}

	// 如果使用 Makefile，需要創建 build 目錄
	if config.UseMake {
		dirs = append(dirs, "build")
	}

	for _, dir := range dirs {
		dirPath := filepath.Join(config.Name, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// 生成源碼文件
	if err := generateSourceFiles(config); err != nil {
		return err
	}

	// 生成建置系統文件
	if config.UseMake {
		if err := generateMakefile(config); err != nil {
			return err
		}
	}

	if config.UseCMake {
		if err := generateCMakeFiles(config); err != nil {
			return err
		}
	}

	// 生成 README
	if err := generateReadme(config); err != nil {
		return err
	}

	// 生成許可證文件
	if config.License != "" {
		if err := generateLicense(config); err != nil {
			return err
		}
	}

	return nil
}

// generateSourceFiles 生成源碼文件
func generateSourceFiles(config ProjectConfig) error {
	ext := ".c"
	if config.Language == "cpp" {
		ext = ".cpp"
	}

	// 生成應用程式源碼
	if config.IsApp {
		mainFile := filepath.Join(config.Name, "src", "main"+ext)
		// TODO: Use text/templates to generate the main source file
		mainContent := templates.GetMainTemplate(config.Language, config.Name, config.IsLib)
		if err := os.WriteFile(mainFile, []byte(mainContent), 0644); err != nil {
			return fmt.Errorf("failed to create main source file: %w", err)
		}
	}

	// 生成庫源碼
	if config.IsLib {
		// 源碼文件
		libSrcFile := filepath.Join(config.Name, "src", config.Name+ext)
		libSrcContent := templates.GetLibSourceTemplate(config.Language, config.Name)
		if err := os.WriteFile(libSrcFile, []byte(libSrcContent), 0644); err != nil {
			return fmt.Errorf("failed to create library source file: %w", err)
		}

		// 頭文件
		headerExt := ".h"
		if config.Language == "cpp" {
			headerExt = ".hpp"
		}
		headerFile := filepath.Join(config.Name, "include", config.Name+headerExt)
		headerContent := templates.GetHeaderTemplate(config.Language, config.Name)
		if err := os.WriteFile(headerFile, []byte(headerContent), 0644); err != nil {
			return fmt.Errorf("failed to create header file: %w", err)
		}
	}

	return nil
}

// generateMakefile 生成 Makefile
func generateMakefile(config ProjectConfig) error {
	makefilePath := filepath.Join(config.Name, "Makefile")
	content := templates.GetMakefileTemplate(config)
	return os.WriteFile(makefilePath, []byte(content), 0644)
}

// generateCMakeFiles 生成 CMake 文件
func generateCMakeFiles(config ProjectConfig) error {
	cmakePath := filepath.Join(config.Name, "CMakeLists.txt")
	content := templates.GetCMakeTemplate(config)
	return os.WriteFile(cmakePath, []byte(content), 0644)
}

// generateReadme 生成 README.md
func generateReadme(config ProjectConfig) error {
	readmePath := filepath.Join(config.Name, "README.md")
	content := templates.GetReadmeTemplate(config)
	return os.WriteFile(readmePath, []byte(content), 0644)
}

// generateLicense 生成許可證文件
func generateLicense(config ProjectConfig) error {
	license := strings.ToLower(config.License)
	content := templates.GetLicenseText(license)

	if content == "" {
		return fmt.Errorf("unsupported license type: %s", license)
	}

	licensePath := filepath.Join(config.Name, "LICENSE")
	return os.WriteFile(licensePath, []byte(content), 0644)
}
