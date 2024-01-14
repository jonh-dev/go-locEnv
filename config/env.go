package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/jonh-dev/go-logger/logger"
)

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
	productionWarning(env)
	developmentProfile(env)
	return env
}

/*
LoadEnv carrega as variáveis de ambiente a partir de um arquivo .env

A função LoadEnv primeiro chama a função findEnvFile para localizar o arquivo .env apropriado e obter o ambiente correspondente.
Se findEnvFile não encontrar um arquivo .env, LoadEnv retorna um erro.

Se um arquivo .env for encontrado, LoadEnv define o campo Env da estrutura FileEnvLoader para o ambiente obtido e carrega as variáveis de ambiente do arquivo .env chamando a função loadEnvFile.

@return error - Um erro se o arquivo .env não puder ser encontrado ou carregado
*/
func (f *FileEnvLoader) LoadEnv() error {
	envFile, env := f.findEnvFile()
	if envFile == "" {
		return fmt.Errorf("arquivo .env não encontrado")
	}
	f.Env = env
	return f.loadEnvFile(envFile)
}

/*
findEnvFile procura um arquivo .env no diretório atual e em todos os subdiretórios

A função findEnvFile usa a função filepath.Walk para percorrer todos os arquivos no diretório atual e em todos os subdiretórios.
Para cada arquivo que encontra, verifica se o nome do arquivo começa com ".env.".
Se encontrar um arquivo que corresponda a esse padrão, define filePath para o caminho do arquivo e env para o ambiente correspondente (que é a parte do nome do arquivo após ".env.").

@return string - O caminho do arquivo .env encontrado
@return string - O ambiente correspondente ao arquivo .env encontrado
*/
func (f *FileEnvLoader) findEnvFile() (string, string) {
	filePath := ""
	env := ""

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasPrefix(info.Name(), ".env.") {
			filePath = path
			env = strings.TrimPrefix(info.Name(), ".env.")
		}

		return nil
	})

	return filePath, env
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

/*
productionWarning exibe um aviso se o ambiente atual for "production" ou "prod"

A função productionWarning verifica se o ambiente atual é "production" ou "prod".
Se for, ela registra uma mensagem de aviso indicando que o ambiente de produção está sendo usado.

@param Environment string - O ambiente atual, que é verificado para ver se é "production" ou "prod"
*/
func productionWarning(Environment string) {
	if Environment == "production" || Environment == "prod" {
		logger.Warning("Você está rodando em produção")
	}
}

/*
developmentProfile exibe uma mensagem informativa se o ambiente atual for "development", "dev" ou vazio

A função developmentProfile verifica se o ambiente atual é "development", "dev" ou vazio.
Se for, ela registra uma mensagem informativa indicando que o ambiente de desenvolvimento está sendo usado.

@param Environment string - O ambiente atual, que é verificado para ver se é "development", "dev" ou vazio
*/
func developmentProfile(Environment string) {
	if Environment == "development" || Environment == "dev" || Environment == "" {
		logger.Info("Você está rodando em desenvolvimento")
	}
}
