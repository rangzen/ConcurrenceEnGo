# Introduction

Vous trouverez dans ce dépôt les sources qui correspondent à l’article du reste de ce README sur la concurrence en Go.
Ce repo est très inspiré de la vidéo [Concurrency in Go](https://www.youtube.com/watch?v=LvgVSSpwND8) de [Jake Wright](https://www.youtube.com/channel/UCc1Pn7FxieMohCZFPYEbs7w).

# Code

## Appel en série

Appelle de la fonction compte pour compter des moutons puis des poissons.  
La boucle infinie du premier appel, empêche l’exécution de l’appel pour les poissons.

## Création d’une goroutine

L’ajout du mot clé **go** permet de créer une goroutine. Une goroutine est un appel à une fonction qui sera exécuté en arrière-plan.  
C’est en fait la deuxième goroutine du programme puisque main est également une goroutine.

## Tout en goroutine

Les goroutines sont très efficaces pourquoi ne pas tout mettre en goroutines ?  
Au lancement du programme, vous allez voir... rien. Les goroutines ne sont pas bloquantes, lorsque la goroutine de main va se terminer (puisqu’elle continue d’avancer dans les instructions), le programme va se terminer. Quelque soit l’état des goroutines lancées.

## Attendre à la fin de main

En ajoutant une attente de données au clavier, les deux goroutines peuvent s’exécuter. Le programme se termine si l’on appuie sur entrée (validation de Scanln).