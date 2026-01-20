# Chapter 24: Packages and Modules

Ce chapitre couvre les concepts fondamentaux des packages et modules en Go.

## Structure du projet

```
24PackagesAndModules/
├── main.go              # Programme principal utilisant les packages
├── go.mod              # Fichier de configuration du module
├── calculator/         # Package personnalisé calculator
│   └── calculator.go   # Implémentation du package calculator
├── utils/              # Package personnalisé utils
│   └── utils.go        # Implémentation du package utils
└── README.md           # Ce fichier
```

## Concepts abordés

### 1. Packages standards
- Utilisation des packages de la bibliothèque standard (fmt, math, strings)
- Import et utilisation des fonctions exportées

### 2. Packages personnalisés
- Création de packages réutilisables
- Organisation du code en modules logiques
- Package `calculator`: opérations mathématiques
- Package `utils`: fonctions utilitaires

### 3. Visibilité des packages
- **Public** (majuscule): accessible depuis d'autres packages
- **Privé** (minuscule): accessible uniquement dans le package
- Règles de nommage Go

### 4. Initialisation des packages
- Fonction `init()` exécutée automatiquement
- Initialisation des variables au niveau du package
- Ordre d'initialisation

### 5. Import et aliasing
- Import simple: `import "package"`
- Import avec alias: `import alias "package"`
- Import vide: `import _ "package"` (effets de bord)

### 6. Go Modules
- Gestion des dépendances avec `go.mod`
- Versioning sémantique
- Commandes essentielles:
  - `go mod init module-name`
  - `go mod tidy`
  - `go get package@version`

## Exécution

```bash
# Initialiser le module (déjà fait)
go mod init github.com/example/24packagesandmodules

# Exécuter le programme
go run main.go

# Nettoyer les dépendances
go mod tidy
```

## Points clés à retenir

1. **Organisation**: Les packages permettent d'organiser le code de manière logique
2. **Visibilité**: Les noms en majuscule sont publics, en minuscule sont privés
3. **Réutilisabilité**: Les packages personnalisés peuvent être réutilisés dans plusieurs projets
4. **Dépendances**: Go modules gère automatiquement les dépendances et leurs versions
5. **Best practices**: Un package par répertoire, noms descriptifs, API minimale

## Packages tiers populaires

- **Web**: github.com/gin-gonic/gin, github.com/gorilla/mux
- **Testing**: github.com/stretchr/testify
- **CLI**: github.com/spf13/cobra
- **Database**: gorm.io/gorm, github.com/go-redis/redis
- **Configuration**: github.com/spf13/viper
