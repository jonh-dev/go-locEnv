package test

import (
	"os"
	"path"
	"testing"

	"github.com/jonh-dev/go-locEnv/config"
)

/*
TestLoadEnv testa a função LoadEnv da interface IEnvLoader

Para testar a função LoadEnv, é necessário criar um diretório temporário e um arquivo .env.test dentro dele.
O arquivo .env.test deve conter a variável de ambiente TEST_VAR=test.

O teste altera o diretório de trabalho para o diretório temporário e define a variável de ambiente GO_ENV para "test".
Em seguida, cria um novo carregador de ambiente e carrega as variáveis de ambiente.

Por fim, verifica se a variável de ambiente TEST_VAR foi definida corretamente.

@params t *testing.T - Ponteiro para o objeto de teste

@return void
*/
func TestLoadEnv(t *testing.T) {
	// Cria um diretório temporário
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Não foi possível criar o diretório temporário: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Limpa o diretório no final

	// Cria um arquivo .env.test no diretório temporário
	envFile := path.Join(tmpDir, ".env.test")
	err = os.WriteFile(envFile, []byte("TEST_VAR=test"), 0644)
	if err != nil {
		t.Fatalf("Não foi possível criar o arquivo .env: %v", err)
	}

	// Define a variável de ambiente GO_ENV para "test"
	os.Setenv("GO_ENV", "test")

	// Altera o diretório de trabalho para o diretório temporário
	os.Chdir(tmpDir)

	// Cria um novo carregador de ambiente
	loader := config.NewEnvLoader()

	// Carrega as variáveis de ambiente
	err = loader.LoadEnv()
	if err != nil {
		t.Errorf("Erro ao carregar variáveis de ambiente: %s", err)
	}

	// Verifica se a variável de ambiente TEST_VAR foi definida corretamente
	testVar := os.Getenv("TEST_VAR")
	if testVar != "test" {
		t.Errorf("Esperado %s, obtido %s", "test", testVar)
	}
}

/*
TestLoadEnvInSubdirectory testa a função LoadEnv da interface IEnvLoader em um subdiretório

Para testar a função LoadEnv, é necessário criar um diretório temporário e um subdiretório dentro dele.
O subdiretório deve conter um arquivo .env.test dentro dele.
O arquivo .env.test deve conter a variável de ambiente TEST_VAR=test.

O teste altera o diretório de trabalho para o subdiretório e define a variável de ambiente GO_ENV para "test".
Em seguida, cria um novo carregador de ambiente e carrega as variáveis de ambiente.

Por fim, verifica se a variável de ambiente TEST_VAR foi definida corretamente.

@params t *testing.T - Ponteiro para o objeto de teste

@return void
*/
func TestLoadEnvInSubdirectory(t *testing.T) {
	// Cria um diretório temporário
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Não foi possível criar o diretório temporário: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Limpa o diretório no final

	// Cria um subdiretório dentro do diretório temporário
	subDir := path.Join(tmpDir, "sub")
	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Não foi possível criar o subdiretório: %v", err)
	}

	// Cria um arquivo .env.test no subdiretório
	envFile := path.Join(subDir, ".env.test")
	err = os.WriteFile(envFile, []byte("TEST_VAR=test"), 0644)
	if err != nil {
		t.Fatalf("Não foi possível criar o arquivo .env: %v", err)
	}

	// Define a variável de ambiente GO_ENV para "test"
	os.Setenv("GO_ENV", "test")

	// Altera o diretório de trabalho para o subdiretório
	os.Chdir(subDir)

	// Cria um novo carregador de ambiente
	loader := config.NewEnvLoader()

	// Carrega as variáveis de ambiente
	err = loader.LoadEnv()
	if err != nil {
		t.Errorf("Erro ao carregar variáveis de ambiente: %s", err)
	}

	// Verifica se a variável de ambiente TEST_VAR foi definida corretamente
	testVar := os.Getenv("TEST_VAR")
	if testVar != "test" {
		t.Errorf("Esperado %s, obtido %s", "test", testVar)
	}
}
