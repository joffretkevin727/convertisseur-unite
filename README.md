# Convertisseur d'Unités - Projet Go

Un convertisseur d'unités simple et robuste développé en **Go**, utilisant des templates HTML et un système de routage dynamique. Ce projet permet de convertir des distances, des masses, des volumes et des températures tout en gérant les erreurs de compatibilité entre catégories.

## 🚀 Fonctionnalités
* **Conversion multi-catégories** : Longueurs (m, km, mi...), Masses (g, kg, lb), Volumes (l, ml) et Températures (°C, °F).
* **Validation stricte** : Empêche la conversion entre unités incompatibles (ex: grammes vers mètres).
* **Interface épurée** : Design moderne avec un fichier CSS externe.
* **Précision adaptative** : Affichage optimisé pour les petites valeurs (ex: millimètres vers miles).

## 📁 Structure du Projet
```text
convertisseur-unite/
├── controller/
│   └── controller.go   # Logique métier et validation
├── router/
│   └── router.go       # Configuration des routes et fichiers statiques
├── structure/
│   └── structure.go    # Définition de la structure PageData
├── static/
│   └── style.css       # Design et mise en page
├── template/
│   └── home.html       # Interface utilisateur (Go Templates)
└── main.go             # Point d'entrée de l'application

🛠️ Installation et Lancement
1. Cloner le dépôt :
    git clone https://github.com/joffretkevin727/convertisseur-unite.git

2. Lancer le serveur :
    cd .\convertisseur-unite\
    go run main.go

3. Accéder à l'application :
    Rendez-vous sur http://localhost:8080/home