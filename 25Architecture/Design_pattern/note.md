# Les piliers de l'artisanat logiciel (Software Craftsmanship). 

- KISS : Keep It Simple,Stupid
- YAGNI : You Aren't Gonna Need It
- SOLID : Single Responsibility Principle - Open-Closed Principle - Liskov Substitution Principle - Interface Segregation Principle - Dependency Inversion Principle
- DRY : Don't Repeat Yourself

Ils ne sont pas des règles rigides, mais des guides pour éviter que le code ne devienne un sac de nœuds indémêlable avec

---

## 1. KISS : Keep It Simple, Stupid

**Le concept :** La complexité est l'ennemi. Si une solution peut être écrite en trois lignes lisibles plutôt qu'en dix lignes "astucieuses", choisissez la simplicité.

* **Pourquoi ?** Un code simple est plus facile à tester, à maintenir et à transmettre à un collègue.
* **Exemple :** Au lieu de créer une usine à gaz avec des classes imbriquées pour calculer une remise, utilisez une simple fonction si le besoin s'arrête là.

## 2. YAGNI : You Aren't Gonna Need It

**Le concept :** Ne développez pas une fonctionnalité tant qu'elle n'est pas nécessaire "maintenant".

* **Pourquoi ?** On passe souvent du temps à coder des options "au cas où" qui ne servent jamais, polluant le projet inutilement.
* **Exemple :** Ne créez pas un système de support multi-langues si vous n'avez que des clients français pour le moment. Attendez le premier client étranger.

## 3. DRY : Don't Repeat Yourself

**Le concept :** Toute connaissance ou logique doit avoir une représentation unique et non ambiguë au sein d'un système.

* **Pourquoi ?** Si vous copiez-collez un calcul à trois endroits et que la formule change, vous devrez modifier trois fichiers (et vous en oublierez sûrement un).
* **Exemple :** Si vous validez un format d'email à la création ET à la modification d'un profil, centralisez cette logique dans une fonction `validateEmail()`.

---

## 4. SOLID : Les 5 commandements de l'Objet

C'est le framework le plus technique, conçu pour rendre les systèmes extensibles.

| Principe | Explication simple | Exemple d'usage |
| --- | --- | --- |
| **S**ingle Responsibility | Une classe ne doit avoir qu'une seule raison de changer (une seule mission). | Une classe `Facture` calcule le prix, mais c'est une classe `FactureImprimante` qui gère l'impression. |
| **O**pen-Closed | Le code doit être ouvert à l'extension, mais fermé à la modification. | On ajoute de nouveaux comportements en créant de nouvelles classes, pas en modifiant les anciennes. |
| **L**iskov Substitution | On doit pouvoir remplacer une classe mère par une classe fille sans casser le programme. | Si une classe `Autruche` hérite de `Oiseau` mais ne peut pas voler, la méthode `voler()` va bugger. C'est une violation de LSP. |
| **I**nterface Segregation | Mieux vaut plusieurs petites interfaces spécifiques qu'une seule grosse interface généraliste. | Ne forcez pas une imprimante simple à implémenter une fonction `faxer()` dont elle n'a pas besoin. |
| **D**ependency Inversion | Dépendre des abstractions (interfaces), pas des implémentations concrètes. | Votre code ne doit pas dépendre de "MySQL", mais d'une interface "BaseDeDonnées". |

---

## Les langages qui "brillent" par ces principes

Bien que ces principes soient applicables partout, certains langages les encouragent par leur structure :

* **Java & C# :** Les rois du **SOLID**. Ils imposent une structure stricte de classes et d'interfaces qui rend l'application de ces principes naturelle (parfois même trop complexe, attention au KISS !).
* **Python & Go :** Les champions du **KISS**. Go, en particulier, a été conçu pour être minimaliste. Il n'y a souvent qu'une seule façon de faire les choses, ce qui limite la sur-ingénierie.
* **Rust :** Excellent pour le **DRY** et la gestion des abstractions grâce à son système de "traits" (similaires aux interfaces mais plus flexibles).
* **TypeScript :** Apporte la rigueur de SOLID au monde du Web, permettant de structurer de gros projets JavaScript sans que cela devienne un chaos technique.

---

### En résumé

* **KISS** : Reste simple.
* **YAGNI** : Ne code pas pour le futur imaginaire.
* **DRY** : Centralise ta logique.
* **SOLID** : Structure ton code pour qu'il soit robuste et évolutif.
