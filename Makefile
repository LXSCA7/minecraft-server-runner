# Nome do executável
BINARY_NAME=script-runner.exe

# Pasta de destino
OUTPUT_DIR=script-server

# Arquivo de configuração
CONFIG_FILE=settings.json
KEY_FILE=key.key
TUTORIAL_FILE=tutorial.txt

# Nome do arquivo ZIP
ZIP_FILE=script-server.zip

# Regra padrão
all: build copy

# Regra para compilar o programa
build:
	@echo Building the project...
	@if not exist "$(OUTPUT_DIR)" mkdir "$(OUTPUT_DIR)"
	@go build -o "$(OUTPUT_DIR)\$(BINARY_NAME)"
	@echo Build completed: $(OUTPUT_DIR)\$(BINARY_NAME)

# Regra para copiar o arquivo de configuração
copy:
	@echo Copying configuration files...
	@if exist "$(CONFIG_FILE)" copy "$(CONFIG_FILE)" "$(OUTPUT_DIR)\"
	@if exist "$(KEY_FILE)" copy "$(KEY_FILE)" "$(OUTPUT_DIR)\"
	@echo Configuration files copied to $(OUTPUT_DIR)

# Regra para criar o arquivo ZIP contendo a pasta
deploy: clean build copy
	@echo Creating ZIP archive...
	@if exist "$(ZIP_FILE)" del "$(ZIP_FILE)"
	@powershell Compress-Archive -Path "$(OUTPUT_DIR)" -DestinationPath "$(ZIP_FILE)"
	@if exist "$(TUTORIAL_FILE)" powershell Compress-Archive -Update -Path "$(TUTORIAL_FILE)" -DestinationPath "$(ZIP_FILE)"
	@echo Deployment package created: $(ZIP_FILE)
	
# Limpar os artefatos gerados
clean:
	@echo Cleaning up...
	@if exist "$(OUTPUT_DIR)" rmdir /s /q "$(OUTPUT_DIR)"
	@if exist "$(ZIP_FILE)" del "$(ZIP_FILE)"
	@echo Cleaned.

# Help para exibir os comandos disponíveis
help:
	@echo Makefile Commands:
	@echo   make          - Build the project and copy configuration file
	@echo   make build    - Build the project only
	@echo   make copy     - Copy the configuration file only
	@echo   make deploy   - Build, copy files, and create a ZIP archive
	@echo   make clean    - Remove build artifacts
