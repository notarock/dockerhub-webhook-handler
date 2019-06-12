# gopdater

Il est important de préciser que ce projet est encore en phase "jouet". Il n'est pas fonctionnel, et il
n'y a pas encore de fonctionnement préétabli. Pour le moment, il permet de faire rouler un script
sur une machine distance, mais pas plus.

----------

## Vision

Écoute, c'est l'fun d'avoir plein de gugusses qui fonctionnent automatiquement. Imagine si tu n'avais pas à aller sur un serveur
pour mettre à jour un site web ou un service... Ça serait Nice, non?
Bien voilà, ce projet utilise des outils qui existent déjà pour permettre d'automatiser le déploiement/MAJ sur un serveur.

> Ouin, mais ça existe déjà sur AWS/etc. ?

Peut-être, mais cette implémentation sera utile pour les gens qui n'ont pas leur service sur des IaaS/PaaS.
En résumé, c'est Nice quand tu as ta propre vm/machine ;).


## Résumé

Permet de recevoir un _web-hook_ et de mettre à jours des images dockers suite à un build depuis Dockerhub.

Ce programme doit répondre à une requête de type *POST* en utilisant les informations envoyées pour mettre à jour (pull) les images présentement actives sur le serveur.

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

Pour utiliser Gopdate, il faut simplement l'incorporer dans un docker-compose qui est déjà fonctionnel, et
configurer la variable d'environnement "OWNER" du service.

Ceci est faisable facilement en utilisant le docker-compose qui est fourni en exemple dans ce répo.
Seul le service *gopdater* est nécessaire dans le docker-compose inclut. Il est important de bien laisser
le volume `/var/run/docker.sock` car ceci permet a Gopdater d'utiliser le daemon Docker de la machine.

Le service gopdater configuré dans ce répo construit l'image à partir du Dockerfile inclut, il est donc
recommander d'utiliser l'image publique host sur dockerhub: `notarock/gopdate` 

### Configuration

Pour chaque service inclus dans votre docker-compose, le nom de l'image doit être la même que le nom
du répo sur Dockerhub, et le propriétaire doit être défini dans les variables d'environnements.

Avec dockerhub, allez dans la section *web-hook* de votre image et ajoutez-y l'URL de votre instance
gopdater. Voilà, c'est fini.

Maintenant, lorsque le web-hook se déclenches, Gopdate pourra être notifié et ensuite repartir la nouvelle
image sur votre machine/serveur/vps. _Magie!_

### Utilisation en locale

On peut lancer le programme simplement avec Docker:

```
docker build . -t gopdater
docker run -v /var/run/docker.sock:/var/run/docker.sock --network=host -d gopdater
```

Ensuite, vous pouvez utiliser cette commande pour tester.

    curl -X POST -H 'Content-Type: application/json' -i http://localhost:8000/ --data '    {
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
        "name": "portfolio-ng",
        "namespace": "svendowideit",
        "owner": "notarock",
        "repo_name": "notarock/portfolio-ng",
        "repo_url": "https://registry.hub.docker.com/u/svendowideit/testhook/",
        "star_count": 0,
        "status": "Active"
      }
    }
    
    
