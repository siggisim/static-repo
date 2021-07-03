Get started with this starter pack
==================================

Below you will see commands to install, run, and deploy a go static server.
---------------------------------------------------------------------------

In just 3 steps you'll have a go server that can serve static files on your local machine or kubernetes, all built with bazel.

1. Install flame
```
$ curl tryflame.com/install.sh | sh
```

2. Download the flame go-static server starter pack.
```
$ flame new [APP_DIRECTORY]  "https://github.com/tryflame/flame-starter-go-static"
```

readme

3. Run the app locally.
```
$ cd [APP_DIRECTORY} && flame run
```

Deploying the app is no more difficult than running locally, but it does require a google cloud (GCE) or amazon web services (AWS) account setup. Go get one of those and come back when you have kubectl ready to run!

* [Setup kubectl on Google](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-access-for-kubectl) | [Setup kubectl on Amazon](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-access-for-kubectl)

4. Deploy the app to kubernetes
```
$ flame run :deploy.create
```

