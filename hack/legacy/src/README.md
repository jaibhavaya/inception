# inception
> automate the automation

## Introduction
This project automates the curation of pipelines in ConcourseCI. It watches all
branches on a project's git repository and takes action for those which match
the desired naming criteria (e.g. `project-[0-9]*`).

On branch creation this system will create a new CI pipeline that matches the
name of the branch. That pipeline will then introspect the branch to discover
which projects it should test, build, push container images for. It also creates
isolated Kubernetes namespaces to deploy into.

On branch deletion, this system will delete all resources associated with the
pipeline it created.

All of this is controlled by per-project configuration defined below.

### Configuration
Projects wishing to benefit from the automation provided by this system must
include an `Inceptionfile` in the root directory of the project and format it
according to these conventions:

```yaml
project:
  # A unique name within the repository. Typically this should exactly match the
  # name of the folder the project resides in. The artifacts produced by this
  # project are named with the value used here.
  name: example-project

  # A boolean flag indicating if this project is under active development. This
  # controls if the project will be built and deployed into a fully isolated and
  # ephemeral feature branch environment. Future versions of Inceptionfiles will
  # include the notion of project dependencies such that signalling `wip` for
  # one project automatically flags those it requires as `wip` as well.
  wip: true

  # Controls the container used by CI for this project.
  image: build

  # Controls which directories are cached between builds for pipelines generated
  # with CACHING_ENABLED=true. Only recommended for use in ephemeral branch
  # pipelines. Caches are specific to each pipeline and project that appears
  # within them. Any work that relies on a cached file or directory must be able
  # to succeed in its absence. Paths are relative to the Inceptionfile.
  caches:
  - ../.gradle
  - build

  # Criteria for CI to run various actions for this project.
  triggers:
    # The project build job will be run in CI when commits are detected that
    # change files listed here. Paths are relative to the Inceptionfile.
    # Note: build triggers only affect the dev CI environment.
    build:
    - src/*
    - ../shared/components/<language>/<library>
    - Dockerfile
    - Makefile
    # The project deploy job will be run in CI when commits are detected that
    # change files listed here. Paths are relative to the Inceptionfile.
    deploy:
    - infrastructure/deployments/${DEPLOYMENT_NAME}/*
    # The project test job will be run in CI when commits are detected that
    # change files listed here. Paths are relative to the Inceptionfile.
    # Note: only use this if you are not using build already
    test:
    - src/*
    - Makefile
    # The project docker image will be pushed in CI when commits are detected
    # that change files listed here. Paths are relative to the Inceptionfile.
    # Note: only use this if you are not using build already
    push:
    - src/*
    - Makefile

  # Commands for this project that will be run by automation in some fashion.
  commands:
    # CI will run this in the root directory of the project as a part of build
    # jobs when changes to build trigger files are detected. This is commonly
    # used to run tests and produce the inputs to a Dockerfile (which will
    # automatically be built and pushed to a registry by CI).
    build: make
```
