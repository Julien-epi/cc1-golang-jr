Bonjour Monsieur, j'espère que mon projet vous plaira :D

pour le tester rien de plus simple il suffit de faire un docker-compose up --build à la racine du projet,

vous pouvez aussi faire un go run . pour le tester en local.

Lors du lancement de l'application un fichier csv se créé automatiquement dans le folder "csvgenerate"

dès que le fichier csv est créé un zip de celui-ci est aussi créé.

vous pouvez tester le telechargement de celui-ci avec la route : http://localhost:8081/archive

elle vous renverra un statut 200 avec le binaire de celui-ci.

le dossier .zip et le fichier csv sont aussi stocké dans le container à -> app/csvgenerate grâce aux volumes dans le docker compose.



