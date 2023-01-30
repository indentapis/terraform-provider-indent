/**

Auth can be handled either with a Google Service Account JSON file, or with the `indent` cli.

To use the `indent` cli:
$ indent auth login --space my-space-name

To verify you have authenticated successfully, you can check `auth view`:
$ indent auth view --space my-space-name 2>&1 | jq
{
  ...
  "msg": "Current User",
  "user": {
    "Id": "my-email@example.com",
    ...
    "Labels": [
        "indent.com/app/config/id": "space:my-space-name"
    ]
    ...
  }
}

To use a Service Account JSON file, simply reference it in the provider config.
**/

provider "indent" {
  // use cli auth
  space = "my-space-name"
}

provider "indent" {
  // use static service account auth
  space         = "my-space-name"
  json_key_file = "/config/service-account-281938728bdd.json"
}
