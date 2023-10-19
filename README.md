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