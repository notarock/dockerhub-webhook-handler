# gopdater

Il est important de préciser que ce projet est encore en phase "jouet". Il n'est pas fonctionnel, et il
n'y a pas encore de fonctionnement pré-établit. Pour le moment, il permet de faire rouller un script
sur une machine distance, mais pas plus.

----------

## Vision

Écoute, c'est l'fun d'avoir pleins de guguss qui fonctionne automatiquement. Imagine si tu n'avais pas à aller sur un serveur
pour mettre à jours un site web ou un service... Ça serait nice, non?
Bien voilà, ce projet utilise des outils qui existent déjà pour permettre d'automatiser le déploiement/MAJ sur un serveurs.

> Ouin mais ça existe déjà sur AWS/etc... ?

Peut-être, mais cette implémentation sera utile pour les gens qui n'ont pas leur service sur des IaaS/PaaS.
En résumé, c'est nice quand tu as ta propre vm/machine ;).


## Résumé

Permet de recevoir un _web-hook_ et de mettre à jours des images dockers suite à la MAJ du build depuis Dockerhub.

Ce programme doit répondre à une requête de type *POST* en utilisant les informations envoyés pour mettre à jours (pull) les images présentement actives sur le serveurs.

Voici un exemple de JSON envoyé par DockerHub lorsqu'un build fonctionne:


    {
      "callback_url": "https://registry.hub.docker.com/u/svendowideit/testhook/hook/2141b5bi5i5b02bec211i4eeih0242eg11000a/",
      "push_data": {
        "images": [
            "27d47432a69bca5f2700e4dff7de0388ed65f9d3fb1ec645e2bc24c223dc1cc3",
            "51a9c7c1f8bb2fa19bcd09789a34e63f35abb80044bc10196e304f6634cc582c",
            "..."
        ],
        "pushed_at": 1.417566161e+09,
        "pusher": "trustedbuilder",
        "tag": "latest"
      },
      "repository": {
        "comment_count": 0,
        "date_created": 1.417494799e+09,
        "description": "",
        "dockerfile": "#\n# BUILD\u0009\u0009docker build -t svendowideit/apt-cacher .\n# RUN\u0009\u0009docker run -d -p 3142:3142 -name apt-cacher-run apt-cacher\n#\n# and then you can run containers with:\n# \u0009\u0009docker run -t -i -rm -e http_proxy http://192.168.1.2:3142/ debian bash\n#\nFROM\u0009\u0009ubuntu\n\n\nVOLUME\u0009\u0009[/var/cache/apt-cacher-ng]\nRUN\u0009\u0009apt-get update ; apt-get install -yq apt-cacher-ng\n\nEXPOSE \u0009\u00093142\nCMD\u0009\u0009chmod 777 /var/cache/apt-cacher-ng ; /etc/init.d/apt-cacher-ng start ; tail -f /var/log/apt-cacher-ng/*\n",
        "full_description": "Docker Hub based automated build from a GitHub repo",
        "is_official": false,
        "is_private": true,
        "is_trusted": true,
        "name": "testhook",
        "namespace": "svendowideit",
        "owner": "svendowideit",
        "repo_name": "svendowideit/testhook",
        "repo_url": "https://registry.hub.docker.com/u/svendowideit/testhook/",
        "star_count": 0,
        "status": "Active"
      }
    }


## Utilisation

Très simple pour le momment:

    go build
    ./gopdater


Une fois que le programme roulle, il peut recevoir une requête HTTP de type *POST* contenant:

    {
    "service": "nom-de-service"
    }

Pour le moment, le programme tente de rouller le script `nom-de-service` qui se trouve dans son dossier
 scripts.
 
