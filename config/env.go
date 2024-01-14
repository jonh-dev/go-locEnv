package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/jonh-dev/go-logger/logger"
)

var ErrEnvFound = errors.New("env found")

/*
IEnvLoader é uma interface que define as funções necessárias para carregar variáveis de ambiente de um arquivo .env

LoadEnv é responsável por localizar e carregar o arquivo .env apropriado com base no ambiente atual.
Se o arquivo .env não puder ser encontrado ou carregado, ele retornará um erro.
@return error - Um erro se o arquivo .env não puder ser encontrado ou carregado

GetEnv retorna o ambiente atual que foi definido ao carregar o arquivo .env.
@return string - O ambiente atual
*/
type IEnvLoader interface {
	LoadEnv() error
	GetEnv() string
}

/*
FileEnvLoader é uma estrutura que implementa a interface IEnvLoader para carregar variáveis de ambiente de um arquivo .env

A estrutura FileEnvLoader contém um campo Env, que armazena o ambiente atual que foi definido ao carregar o arquivo .env.

Env string - O ambiente atual, que é definido ao carregar o arquivo .env
*/
type FileEnvLoader struct {
	Env string
}

/*
NewEnvLoader cria um novo carregador de ambiente que implementa a interface IEnvLoader

A função NewEnvLoader cria uma nova instância de FileEnvLoader, que implementa a interface IEnvLoader.
Ela define o ambiente atual chamando a função getEnvironment e armazena o resultado no campo Env da nova instância de FileEnvLoader.

@return IEnvLoader - Uma nova instância de FileEnvLoader que implementa a interface IEnvLoader
*/
func NewEnvLoader() IEnvLoader {
	return &FileEnvLoader{
		Env: getEnvironment(),
	}
}

/*
getEnvironment obtém o ambiente atual a partir da variável de ambiente GO_ENV

A função getEnvironment tenta obter o valor da variável de ambiente GO_ENV.
Se GO_ENV estiver definido, ele chama as funções productionWarning e developmentProfile para registrar mensagens de aviso ou perfil, dependendo do valor de GO_ENV.
Finalmente, ele retorna o valor de GO_ENV.

@return string - O valor da variável de ambiente GO_ENV
*/
func getEnvironment() string {
	env := os.Getenv("APP_ENV")
	return env
}

/*
LoadEnv carrega as variáveis de ambiente a partir de um arquivo .env

A função LoadEnv primeiro chama a função findEnvFile para localizar o arquivo .env apropriado e obter o ambiente correspondente.
Se findEnvFile não encontrar um arquivo .env ou ocorrer um erro durante a busca, LoadEnv retorna um erro.

Se um arquivo .env for encontrado, LoadEnv define o campo Env da estrutura FileEnvLoader para o ambiente obtido e carrega as variáveis de ambiente do arquivo .env chamando a função loadEnvFile.

@return error - Um erro se o arquivo .env não puder ser encontrado, ocorrer um erro durante a busca, ou o arquivo .env não puder ser carregado
*/
func (f *FileEnvLoader) LoadEnv() error {
	envFile, env, err := f.findEnvFile()
	if err != nil {
		return err
	}
	if envFile == "" {
		return fmt.Errorf("arquivo .env não encontrado")
	}
	f.Env = env
	return f.loadEnvFile(envFile)
}

/*
findEnvFile é um método da estrutura FileEnvLoader que procura um arquivo .env no diretório atual e nos diretórios pais.

O método começa obtendo o diretório de trabalho atual. Se houver um erro ao obter o diretório de trabalho atual, o método retorna um erro.

Em seguida, o método chama a função searchInCurrentAndParentDirectories, passando o diretório de trabalho atual. Esta função procura um arquivo .env no diretório atual e nos diretórios pais, retornando o caminho do arquivo e o ambiente correspondente. Se ocorrer um erro durante a busca, o método retorna esse erro.

@return string - O caminho do arquivo .env encontrado
@return string - O ambiente correspondente ao arquivo .env encontrado
@return error - Um erro se o diretório de trabalho atual não puder ser obtido, ou se ocorrer um erro durante a busca
*/
func (f *FileEnvLoader) findEnvFile() (string, string, error) {
	filePath := ""
	env := ""

	currentDir, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	filePath, env, err = f.searchInCurrentAndParentDirectories(currentDir)
	if err != nil {
		return "", "", err
	}

	return filePath, env, nil
}

/*
searchInCurrentAndParentDirectories é um método da estrutura FileEnvLoader que procura um arquivo .env no diretório atual e nos diretórios pais.

O método recebe o diretório atual como argumento. Inicializa duas variáveis, filePath e env, para armazenar o caminho do arquivo .env encontrado e o ambiente correspondente.

Em seguida, entra em um loop infinito. Dentro do loop, o método chama a função searchInDirectory, passando o diretório atual. Se um arquivo .env for encontrado, o loop é interrompido.

Se nenhum arquivo .env for encontrado, o método obtém o diretório pai do diretório atual. Se o diretório pai for a raiz ("/") ou o diretório atual (".") o loop é interrompido.

Se ocorrer um erro durante a busca, o método retorna esse erro.

@return string - O caminho do arquivo .env encontrado
@return string - O ambiente correspondente ao arquivo .env encontrado
@return error - Um erro se ocorrer um erro durante a busca
*/
func (f *FileEnvLoader) searchInCurrentAndParentDirectories(currentDir string) (string, string, error) {
	filePath := ""
	env := ""

	for {
		var err error
		filePath, env, err = f.searchInDirectory(currentDir)
		if err != nil {
			return "", "", err
		}
		if filePath != "" {
			break
		}

		currentDir = filepath.Dir(currentDir)
		if currentDir == "/" || currentDir == "." {
			break
		}
	}

	return filePath, env, nil
}

/*
searchInDirectory é um método da estrutura FileEnvLoader que procura um arquivo .env no diretório fornecido.

O método recebe o diretório como argumento. Inicializa duas variáveis, filePath e env, para armazenar o caminho do arquivo .env encontrado e o ambiente correspondente.

Em seguida, o método chama a função filepath.Walk, passando o diretório e uma função anônima. A função anônima é chamada para cada arquivo e diretório no diretório fornecido.

Se o arquivo atual for um diretório, a função anônima retorna e passa para o próximo arquivo. Se o arquivo atual for um arquivo e seu nome começar com ".env.", a função anônima verifica se o ambiente correspondente ao arquivo .env (obtido removendo ".env." do nome do arquivo) corresponde ao ambiente atual. Se corresponder, define filePath para o caminho do arquivo e env para o ambiente correspondente, e retorna um erro especial para parar a função filepath.Walk.

Se ocorrer um erro durante a busca, o método retorna esse erro.

@return string - O caminho do arquivo .env encontrado
@return string - O ambiente correspondente ao arquivo .env encontrado
@return error - Um erro se ocorrer um erro durante a busca
*/
func (f *FileEnvLoader) searchInDirectory(dir string) (string, string, error) {
	filePath := ""
	env := ""

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasPrefix(info.Name(), ".env.") {
			env = strings.TrimPrefix(info.Name(), ".env.")
			if env == f.Env {
				filePath = path
				return ErrEnvFound
			}
		}

		return nil
	})

	if err == ErrEnvFound {
		err = nil
	}

	return filePath, env, err
}

/*
loadEnvFile carrega as variáveis de ambiente de um arquivo .env específico

A função loadEnvFile usa a biblioteca godotenv para carregar as variáveis de ambiente do arquivo .env especificado.
Se ocorrer um erro ao carregar o arquivo .env, ele registra o erro e retorna um erro.

Se o arquivo .env for carregado com sucesso, ele registra uma mensagem de sucesso indicando que o ambiente foi carregado.

@param envFile string - O caminho do arquivo .env a ser carregado

@return error - Um erro se o arquivo .env não puder ser carregado
*/
func (f *FileEnvLoader) loadEnvFile(envFile string) error {
	err := godotenv.Load(envFile)
	if err != nil {
		logger.Error(fmt.Sprintf("Erro ao carregar variáveis de ambiente: %s", err.Error()))
		return fmt.Errorf("erro ao carregar variáveis de ambiente: %s", err.Error())
	}

	logger.Success(fmt.Sprintf("Ambiente de %s carregado", f.Env))
	return nil
}

/*
GetEnv retorna o ambiente atual que foi definido ao carregar o arquivo .env

A função GetEnv retorna o valor do campo Env da estrutura FileEnvLoader, que foi definido ao carregar o arquivo .env.

@return string - O ambiente atual que foi definido ao carregar o arquivo .env
*/
func (f *FileEnvLoader) GetEnv() string {
	return f.Env
}
