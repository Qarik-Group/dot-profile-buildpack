# .profile Buildpack

All Cloud Foundry applications can include a `.profile` file which is run in each application container prior to the application commencing. This allows you to create new configuration files, migrate databases, and more. All applications except Java and Binary buildpack applications because you are only uploading the `.jar` or binary, respectively, so the `.profile` file is not uploaded.

This buildpack provides a workaround. If you set the `PROFILED` environment variable, this buildpack will create a `.profile` file in each application container for you.

## Usage

```plain
cf push javaapp -b profiled_buildpack -b java_buildpack --path build/jibs/myapp-1.0.0.jar --no-start
cf set-env javaapp PROFILED "$(cat <<SHELL
#!/bin/bash

cat > config.json <<JSON
{
  "some": "config"
}
JSON
SHELL
)"
cf start javaapp
```

During staging you will see the `profiled_buildpack` kick in and create the `.profile` within your application:

```plain
   -----> Java ProfileD Buildpack version 0.1.0.20190615140303
          Creating .profile
```

Once the application is running we can confirm that `.profile` was created, that it was evaluated, and that it created a `config.json` file:

```plain
$ cf ssh javaapp -c "cat app/config.json"
{
  "some": "config"
}
```

### Updating .profile

If you want to change the `.profile` file then update the `$PROFILED` environment variable. Either:

* Run `cf set-env` and then `cf restage`; or
* Update the `manifest.yml` and re-deploy with `cf push -f manifest.yml`

### Re-run .profile

The `.profile` file will be run once during each application container start. If you want the file to be reevaluated, try `cf restart` command to recreate all containers.
