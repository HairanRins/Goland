# Remarques , explications et utilisations

## Affichage des données dans la console

### 1. fmt.Println

➤ Affiche les valeurs avec un espace entre elles et ajoute un retour à la ligne à la fin.
Usage simple : quand on veut juste afficher quelque chose rapidement.

```
fmt.Println("Bonjour", "le", "monde")
// → Bonjour le monde
// (et saute une ligne après)
```

### 2. fmt.Print

➤ Affiche les valeurs sans saut de ligne automatique, ni espace ajouté.
Usage : quand on veut contrôler la mise en forme soi-même.

```
fmt.Print("Bonjour")
fmt.Print(" le monde")
// → Bonjour le monde
```

(Si on veut une nouvelle ligne, on doit ajouter `\n` soi-même.)

### 3. fmt.Printf

➤ **“Print formatted”** → Permet d’afficher du texte formaté avec des placeholders (`%s`, `%d`, `%f`, etc.).
Usage : quand on veut afficher des variables avec un format précis.

```
nom := "Alice"
age := 22
fmt.Printf("Bonjour %s, tu as %d ans.\n", nom, age)
// → Bonjour Alice, tu as 22 ans.
```

#### Placeholders courants :

| Placeholder | Type                                                    | Exemple   |
| ----------- | ------------------------------------------------------- | --------- |
| `%s`        | string                                                  | `"texte"` |
| `%d`        | entier                                                  | `42`      |
| `%f`        | float                                                   | `3.14`    |
| `%t`        | booléen                                                 | `true`    |
| `%v`        | valeur “par défaut” (pratique pour structs, interfaces) |           |
| `%+v`       | struct avec noms des champs                             |           |
| `%#v`       | valeur avec syntaxe Go (utile pour debug)               |           |

| Fonction  | Saut de ligne auto | Espaces automatiques | Formatage possible |
| --------- | ------------------ | -------------------- | ------------------ |
| `Print`   | ❌                 | ❌                   | ❌                 |
| `Println` | ✅                 | ✅                   | ❌                 |
| `Printf`  | ❌                 | ❌                   | ✅ (placeholders)  |

---

## Les bases solides et les mécanismes internes de Go

### Type Casting

Le **type casting** (ou conversion de type) permet de changer le type d’une variable manuellement.

```
var x float64 = 3.7
var y int = int(x) // cast vers int → 3
```

En Go, **il n’y a pas de conversion implicite** entre types → on doit le faire soi-même.

#### `strconv.Atoi`

- Signification : “ASCII to Integer”.

- Sert à convertir une chaîne (`string`) en entier (`int`).

```
import "strconv"

s := "42"
n, err := strconv.Atoi(s)
if err != nil {
    fmt.Println("Erreur:", err)
}
fmt.Println(n + 8) // → 50
```

_Inverse_ : `strconv.Itoa(42)` → `"42"`

### Maps

Une **map** est une structure clé-valeur, comme un dictionnaire ou objet JSON.

```
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
fmt.Println(ages["Bob"]) // 30

ages["Eve"] = 22 // ajout
delete(ages, "Bob") // suppression
```

Vérifier si une clé existe :

```
value, exists := ages["Bob"]
if exists {
    fmt.Println(value)
}
```

### Range

`range` permet d’itérer sur des tableaux, slices, maps, ou strings.

```
nums := []int{10, 20, 30}
for i, v := range nums {
    fmt.Println(i, v)
}
```

Sur une map :

```
for key, val := range ages {
    fmt.Println(key, val)
}
```

### Recursion

Une **fonction récursive** s’appelle elle-même — utile pour calculs répétés ou structures arborescentes.

```
func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Factorial(n-1)
}
fmt.Println(Factorial(5)) // 120
```

### Slices

Les **slices** sont comme des tableaux dynamiques : taille flexible, références légères.

```
nums := []int{1, 2, 3}
nums = append(nums, 4)
fmt.Println(nums) // [1 2 3 4]

part := nums[1:3]
fmt.Println(part) // [2 3]
```

Les slices pointent vers le **même tableau sous-jacent**, donc les modifications peuvent se propager.

### Structures (`struct`)

Permettent de regrouper plusieurs champs de données sous un même type.

```
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 25}
fmt.Println(p.Name) // Alice
```

Les `struct` **servent souvent à représenter des entités** (user, produit, etc.).

### Pointers

Un **pointeur** stocke l’adresse mémoire d’une variable.
Utile pour **modifier une valeur dans une fonction** ou éviter les copies.

```
x := 10
p := &x     // pointeur sur x
fmt.Println(*p) // → 10 (déréférencement)

*p = 20     // modifie x via le pointeur
fmt.Println(x) // → 20
```

### Scopes

Le **scope** détermine où une variable est visible :

- `{}` délimite un bloc local
- Les variables en dehors d’une fonction → **globales**

```
var a = 10 // global

func main() {
    b := 5 // local à main
    fmt.Println(a, b)
}
```

### Strings

Les chaînes (`string`) sont **immutables**. On peut les parcourir, mesurer, concaténer, etc.

s := "Hello"
fmt.Println(len(s)) // 5
fmt.Println(s[1]) // 101 (byte 'e')
fmt.Println(s + " World")

Fonctions utiles (`strings` package) :

```
import "strings"

strings.ToUpper("go")      // "GO"
strings.Contains("golang", "go") // true
strings.Split("a,b,c", ",") // ["a" "b" "c"]
```

### Méthodes très utiles à maîtriser en Go

| Catégorie   | Fonctions clés                                                 | Usage                      |
| ----------- | -------------------------------------------------------------- | -------------------------- |
| **fmt**     | `Print`, `Println`, `Printf`, `Sprintf`                        | Affichage formaté          |
| **strconv** | `Atoi`, `Itoa`, `ParseFloat`, `FormatInt`                      | Conversion string ↔ nombre |
| **strings** | `Split`, `Join`, `Contains`, `HasPrefix`, `Replace`, `ToLower` | Manipulation de chaînes    |
| **time**    | `Now`, `Format`, `Add`, `Sleep`, `Since`                       | Gestion du temps           |
| **math**    | `Pow`, `Sqrt`, `Round`, `Abs`, `Max`                           | Calculs mathématiques      |
| **sort**    | `Ints`, `Strings`, `Slice`                                     | Tri                        |
| **errors**  | `New`, `Is`, `As`, `Wrap`                                      | Gestion d’erreurs          |
| **reflect** | introspection de types                                         | Avancé (rarement au début) |
