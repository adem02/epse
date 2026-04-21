package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/common"
	"github.com/adem02/epse/internal/templates"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func CreateAuthFileFromTmpl(projectType typeutils.ProjectType, tmplName, destPath string) error {
	if osutils.FileOrDirectoryExists(destPath) {
		logutils.Logger{}.Warning(fmt.Sprintf("⚠️ File already exists, skipping: %s", destPath))
		return nil
	}

	if err := osutils.CreateDirectory(filepath.Dir(destPath)); err != nil {
		return err
	}

	tmplPath := GetAuthTemplatePath(projectType, tmplName)
	data := &struct{}{}

	return osutils.CreateFileFromTmpl(templates.FS, tmplPath, destPath, data)
}

func CreateAllLiteAuthFiles() error {
	files := []struct {
		tmpl string
		dest string
	}{
		{"jwt.config.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("config", "jwt.config.ts")},
		{"auth.types.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("types", "auth.types.ts")},
		{"jwt.service.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("services", "jwt.service.ts")},
		{"auth.service.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("services", "auth.service.ts")},
		{"auth.middleware.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("middlewares", "auth.middleware.ts")},
		{"login.controller.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("controllers", "login.controller.ts")},
		{"register.controller.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("controllers", "auth", "register.controller.ts")},
		{"auth.routes.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("routes", "auth.routes.ts")},
		{"express.d.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("types", "express.d.ts")},
	}

	for _, f := range files {
		if err := CreateAuthFileFromTmpl(typeutils.LiteProjectType, f.tmpl, f.dest); err != nil {
			return err
		}
	}

	return nil
}

func CreateAllCleanAuthFiles() error {
	files := []struct {
		tmpl string
		dest string
	}{
		{"jwt.config.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("config", "jwt.config.ts")},
		{"auth.types.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("entities", "types", "auth.types.ts")},
		{"JwtManager.interface.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("useCases", "gateway", "JwtManager.interface.ts")},
		{"JwtManager.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("adapters", "gateway", "JwtManager.ts")},
		{"authentication.ts.tmpl", common.GetFileOrDirectoryPathFromSrcPath("frameworks", "tsoa", "authentication.ts")},
		{"Login.controller.ts.tmpl", GetCleanAuthControllerPath("login", "Login.controller.ts")},
		{"Login.dto.ts.tmpl", GetCleanAuthControllerPath("login", "Login.dto.ts")},
		{"Register.controller.ts.tmpl", GetCleanAuthControllerPath("register", "Register.controller.ts")},
		{"Register.dto.ts.tmpl", GetCleanAuthControllerPath("register", "Register.dto.ts")},
		{"Login.useCase.ts.tmpl", GetCleanAuthUseCasePath("login", "Login.useCase.ts")},
		{"Login.request.ts.tmpl", GetCleanAuthUseCasePath("login", "Login.request.ts")},
		{"Login.response.ts.tmpl", GetCleanAuthUseCasePath("login", "Login.response.ts")},
		{"Register.useCase.ts.tmpl", GetCleanAuthUseCasePath("register", "Register.useCase.ts")},
		{"Register.request.ts.tmpl", GetCleanAuthUseCasePath("register", "Register.request.ts")},
		{"Register.response.ts.tmpl", GetCleanAuthUseCasePath("register", "Register.response.ts")},
	}

	for _, f := range files {
		if err := CreateAuthFileFromTmpl(typeutils.CleanProjectType, f.tmpl, f.dest); err != nil {
			return err
		}
	}

	return nil
}

func AddAuthRouteInIndexFile() error {
	indexFilePath := filepath.Join(osutils.GetCurrentDirPath(), "src", "routes", "index.ts")

	bytes, err := os.ReadFile(indexFilePath)
	if err != nil {
		return fmt.Errorf("❌ File `routes/index.ts` not found in project")
	}

	content := string(bytes)

	importStatement := `import authRoutes from './auth.routes';`
	newRouteUse := `router.use('/auth', authRoutes);`

	common.AddImportStatementInFile(&content, importStatement)
	if !strings.Contains(content, newRouteUse) {
		common.InsertContentBeforeStatementInFile(&content, "export default router;", newRouteUse)
	}

	if err := os.WriteFile(indexFilePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("❌ Error while writing in routes/index.ts: %w", err)
	}

	return nil
}

func UpdateApiErrorTypes() error {
	if err := updateApiErrorCode(); err != nil {
		return err
	}
	return updateApiErrorKey()
}

func updateApiErrorCode() error {
	path := common.GetFileOrDirectoryPathFromSrcPath("entities", "error", "ApiErrorCode.enum.ts")

	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("❌ ApiErrorCode.enum.ts not found")
	}

	content := string(bytes)

	additions := ""
	if !strings.Contains(content, "Unauthorized") {
		additions += "  Unauthorized = 401,\n"
	}
	if !strings.Contains(content, "Forbidden") {
		additions += "  Forbidden = 403,\n"
	}

	if additions == "" {
		return nil
	}

	content = strings.Replace(content, "}", additions+"}", 1)
	return os.WriteFile(path, []byte(content), 0644)
}

func updateApiErrorKey() error {
	path := common.GetFileOrDirectoryPathFromSrcPath("entities", "error", "ApiErrorKey.type.ts")

	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("❌ ApiErrorKey.type.ts not found")
	}

	content := string(bytes)

	if strings.Contains(content, "auth/") {
		return nil
	}

	authKeys := "\n\n  // Auth\n" +
		"  | 'auth/missing-header'\n" +
		"  | 'auth/invalid-token'\n" +
		"  | 'auth/token-expired-or-invalid'\n" +
		"  | 'auth/insufficient-role'\n" +
		"  | 'auth/invalid-security-name'"

	lastSemicolon := strings.LastIndex(content, ";")
	if lastSemicolon == -1 {
		return fmt.Errorf("❌ invalid ApiErrorKey.type.ts format")
	}

	content = content[:lastSemicolon] + authKeys + ";\n"

	return os.WriteFile(path, []byte(content), 0644)
}

func UpdateDiConstants() error {
	path := common.GetFileOrDirectoryPathFromSrcPath("utilities", "di.constants.ts")

	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("❌ di.constants.ts not found")
	}

	content := string(bytes)

	if strings.Contains(content, "JwtManager") {
		return nil
	}

	content = strings.TrimRight(content, "\n")
	content += "\nexport const JwtManagerToken = Symbol('JwtManager');\n"

	return os.WriteFile(path, []byte(content), 0644)
}
