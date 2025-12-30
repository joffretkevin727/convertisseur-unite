# Convertisseur d'Unités - Projet Go

Un convertisseur d'unités simple et robuste développé en **Go**, utilisant des templates HTML et un système de routage dynamique. Ce projet permet de convertir des distances, des masses, des volumes et des températures tout en gérant les erreurs de compatibilité entre catégories.

## 🚀 Fonctionnalités
- **Conversion multi-catégories** : Longueurs (m, km, mi...), Masses (g, kg, lb), Volumes (l, ml) et Températures (°C, °F).
- **Validation stricte** : Empêche la conversion entre unités incompatibles (ex: grammes vers mètres).
- **Interface épurée** : Design moderne avec un fichier CSS externe.
- **Précision adaptative** : Affichage optimisé pour les petites valeurs (ex: millimètres vers miles).

## 📁 Structure du Projet
```text
proj1/
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
1. Cloner le projet ou extraire l'archive.

2. Ouvrir un terminal dans le dossier racine proj1/.

3. Initialiser le module (si nécessaire) :
    go mod init proj1

4. Lancer le serveur :
    go run main.go
5. Ouvrir votre navigateur à l'adresse suivante : http://localhost:8080/home