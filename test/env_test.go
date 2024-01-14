package test

import (
	"os"
	"path"
	"testing"

	"github.com/jonh-dev/go-locEnv/config"
)

/*
TestLoadEnvInCurrentDirectory é uma função de teste que verifica se a função LoadEnv
é capaz de carregar variáveis de ambiente de um arquivo .env no diretório atual.

A função primeiro cria um diretório temporário e um subdiretório dentro dele. Em seguida, cria um arquivo .env.test no subdiretório e define a variável de ambiente APP_ENV para "test".

Depois, a função muda o diretório de trabalho para o subdiretório e cria um novo carregador de ambiente. A função LoadEnv é então chamada para carregar as variáveis de ambiente.

Finalmente, a função verifica se a variável de ambiente TEST_VAR foi definida corretamente. Se a variável de ambiente não for igual a "test", o teste falha.

@params t *testing.T - Um ponteiro para o objeto de teste

@return error - Um erro se a função LoadEnv não conseguir carregar as variáveis de ambiente do arquivo .env
*/
func TestLoadEnvInCurrentDirectory(t *testing.T) {
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
	os.Setenv("APP_ENV", "test")

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

/*
TestLoadEnvInParentDirectory é uma função de teste que verifica se a função LoadEnv
é capaz de carregar variáveis de ambiente de um arquivo .env no diretório pai.

A função primeiro cria um diretório temporário e um subdiretório dentro dele. Em seguida, cria um arquivo .env.test no diretório temporário (que é o diretório pai do subdiretório) e define a variável de ambiente APP_ENV para "test".

Depois, a função muda o diretório de trabalho para o subdiretório e cria um novo carregador de ambiente. A função LoadEnv é então chamada para carregar as variáveis de ambiente.

Finalmente, a função verifica se a variável de ambiente TEST_VAR foi definida corretamente. Se a variável de ambiente não for igual a "test", o teste falha.

@params t *testing.T - Um ponteiro para o objeto de teste

@return error - Um erro se a função LoadEnv não conseguir carregar as variáveis de ambiente do arquivo .env
*/
func TestLoadEnvInParentDirectory(t *testing.T) {
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
	os.Setenv("APP_ENV", "test")

	// Cria um subdiretório dentro do diretório temporário
	subDir := path.Join(tmpDir, "sub")
	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Não foi possível criar o subdiretório: %v", err)
	}

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
