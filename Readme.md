# tags - a go app to handle tags generically

[![pipeline](https://gitlab.com/meinto/tags/badges/master/pipeline.svg)](https://gitlab.com/meinto/tags)

This go app provides a simple api to save "tags" in json format. Use it for example to count pageviews of your website.

## Example

Save the following tag by calling the `/create` endpoint of the go app with this payload:

```js
{
  uri: "/the/uri/of/the/subpage",
  some: "other",
  info: "mations",
  foo: "bar"
}
```

Later you can count tags by csalling the `/count` endpoint with the following payload:

```js
{
  uri: "/the/uri/of/the/subpage"
}

// or if you want more specific tags:
{
  uri: "/the/uri/of/the/subpage",
  foo: "bar"
}
```

Both json payloads above would match the saved tag. The following tag would not match:

```js
// this would NOT match
{
  uri: "/the/uri/of/the/subpage",
  unknown: "property"
}
```