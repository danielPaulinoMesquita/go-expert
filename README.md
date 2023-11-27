## FOUNDATION / FUNDAÇÃO

### 1 - Entendendo a primeira linha
* Tudo que estiver dentro do mesmo package é visivel para os outros arquivos go do mesmo pacote/package
* package tem o nome do pacote que ele tiver inserido, caso não tem, ele vai ser main.

### 2- Declaração e atribuição
* you can use shorthand to create var ":="

### 3- Criando um type(tipo) de variável
* type ID int
* type IdString string:

  f ID      = 21
  g IdString = "Testando novo tipo variável String"


### 4- Importando fmt e tipagem
importando o fmt e imprimindo valores e tipos utilizando as func, como
Printf("test type %T or value %v")  and Println

### 5- Percorrer Arrays
verificar o tamanho do array len(meuArray) <br />  
for i, v := range meuArray {
fmt.Printf("The Value of index is %d and the value is %d\n", i, v)
}

### 6- Slice
anyArray[:] this is slice "[:]", you can get an interval of values from some array
starting for the left trough the right or start index to final index.

### 7- Maps
var mapNamesOld = make(map[int]string) < -second way to create a map
var mapNamesNew = map[int]string{} <- third way to create, initializing map with empty
to print all elements of the map
for chave, valor := range mapNames {
fmt.Printf("Verificando chaves:(%s) valores=%d\n", chave, valor)
}

### 8 - functions
Please, take aa lo ok at code, directory 8

### 9 - Variadic Functions
func sum(numbers ...int) int {

### 10 - Closure
it's like anonymous class from JAVA


### 11 - Iniciando Com Structs.
É como se fosse Class of java

### 12 - Structs composite
Você pode ver mais no dir 12

### 13 - Métodos em Structs
Pode criar funções dentro do structs como se fosse um comportamento(methods from java)

### 14 - Interfaces
Interface só tem métodos, não possui variáveis/atributos. Toda struct(classe)
que cria um method(func) igual o da interface, implementa ela "por de baixo dos panos".

### 15 - pointers
'*' get the value from reference, & get the reference 

### 16 - pointers and structs
you will see in the dir pointer-structs and empty-interfac

### 17 - modules and packages
when you want to use a module, you need to execute the go mod init "nameOfModule"
- Rules about functions, variables and structs from another file inside a module
- - When the first letter is uppercase means it is visible in another file, like a 'export' in javascript
- - if first letter is lowercase, only his file can see it. 

### 18 - Generics
You can use T for receive any types of argues, you can understand better if you access the dir name Generics

### 19 - Install packages
you must use go get 'package name' 

### 20 - Compiling projects
you can run 
* go build <-- if you are into a module, this command goes to find the right file(where is placed the main) to compile, and the name will be the same as the module
* go build -o 'new name of build file' <-- this will change the name, and doesn't will the same of module
* go build 'name of file or main file' <- this will create a file compile 
* GOOS=windows go build 'name of file or main file' <- this will create a file compile(executable) to run on windows
* GOOS=linux go build 'name of file or main file' <- this will create a file compile(executable) to run on linux
* go tool dist list <-- to see all platforms
* go env GOOS GOARCH


### 27 and 28 Request using http
dirs with requests in Http 
* calling-http-requests
* busca-CEP
* calling-http-requests

Reference to know others platforms:
https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures


### Create a module
go mod init github.com/devfullcycle/goexpert/7-packing/1
this url allow you to access and get this module in another project

####  example
<pre>
    <code>
module github.com/danielPaulinoMesquita/go-expert/package_modules

go 1.21.3
    </code>
</pre>

