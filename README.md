# Build numbers for concourse

This is simpler version of [semver](https://github.com/concourse/semver-resource).
Compare to that *semver*, this one produces only build numbers so the versioning
is simpler and you don't have to handle bumping of multiple numbers in your CI.

Feel free to post support for other new backends (only S3 is supported right now)
or support for *YEAR.BUILD* format. That's something I would like to add but I
don't need in near future.

## Build & tests

To build the code you need:

    * [Golang](https://golang.org/)
    * [dep](https://github.com/golang/de)
    * [Make](https://www.gnu.org/software/make/)

If you have it all just run:

    make


It runs the tests and builds the binaries.

## Resource configuration

I tested this resource only in DigitalOcean's Spaces but it should be compatible
with everything else similar to S3.

At first add new resource type:

    resource_types:
      - name: build-number-resource
      type: docker-image
      source:
          repository: rosti/concourse-build-number-resource
          tag: latest

And then the resource:

    - name: build-number
      type: build-number-resource
      source:
        endpoint: "ams3.digitaloceanspaces.com"
        access_key: "xxxxx"
        secret_key: "xxxxxx"
        use_ssl: true
        bucket: "my-bucket"
        project: "test"
        initial_value: 1
        bump: false

Usage in the jobs is easy:

    jobs:
      - name: some-job
        plan:
        - get: build-number # Read lastest build number
        - ... something important
        - put: build-number # Bump the build number up for next time

## Check

Returns last version. If *bump/ parameter is **true** then it bumps the number
up. I prefer to use OUT instead of this but it depends on the usage.

## IN

Downloads current version number and saves it into file called:

    build-number

## OUT

This bumps up the number regardless the *source*'s *bump* parameter.