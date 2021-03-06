# Build numbers for concourse

This is simpler version of [semver](https://github.com/concourse/semver-resource).
Compare to that *semver*, this one produces only build numbers so the versioning
is simpler and you don't have to handle bumping of multiple numbers in your CI.

On the other hand I combine both tools because I need a lot of builds per day
for my applications but real release is done only a few times per year. It's
for environments where agile development is installed and it can also help
to create releases in case of need.

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
        - get: build-number # Read latest build number
          params:
            bump: true
        - ... something important
        - put: build-number # Bump the build number up for next time

## Check

Returns last version. If *bump* parameter is **true** then it bumps the number
up. I prefer to use IN before docker image is building instead of this but it's
completely up to you.

## IN

Downloads current version number and saves it into file called:

    build-number

If *bump* parameter is **true** the build number is increased by 1.

## OUT

This bumps up the number regardless the *source*'s *bump* parameter.
