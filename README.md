# google-cloud-datastore-example

Download the [Datastore emulator](https://cloud.google.com/datastore/docs/tools/datastore-emulator).

Then run
```bash
# Load the env vars that tell your app to use the emulator (localhost)
eval "$(gcloud beta emulators datastore env-init)"

# Start the emulator service
gcloud beta emulators datastore start

# Run the app (in a new tab)
go run *.go
```

