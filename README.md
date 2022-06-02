# url-health

A service that runs a cron on variable time to check the health of services you're interested in watching.

## Overview

The url-health service has a cron that runs base on a sleep time that can be set in configuration or changed during the service's run-time. The service will sleep for a certain amount of time (in seconds) and check the list of urls is up.

## APIs

Check status of service:

```yaml
api:
  path: /health
  method: GET
  successResponseCode: 200
```

Add new url to check:

```yaml
api:
  path: /add
  method: POST
  requestBody: { "url": "website.com" }
  successResponseCode: 201
```

Delete a url being tracked:

```yaml
api:
  path: /delete
  method: DELETE
  requestBody: { "url": "website.com" }
  successResponseCode: 204
```

List all urls being tracked:

```yaml
api:
  path: /list
  method: GET
  successResponseCode: 200
```

Get status of all urls:

```yaml
api:
  path: /status
  method: GET
  successResponseCode: 200
```

Get status of one url:

```yaml
api:
  path: /status?url=website.com
  method: GET
  successResponseCode: 200
```

Change the sleep value for tracking:

```yaml
api:
  path: /sleep
  method: POST
  requestBody: { "sleep": 600 }
  successResponseCode: 201
```
