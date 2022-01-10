# vecty-examples
Basic usage of Vecty framework examples.

## Instructions
1. Change directory to the folder with the example you wish to run
2. Run `wasmserve`. To install:
    ```shell
    go install github.com/hajimehoshi/wasmserve@latest
    ```
3. Navigate in browser to [localhost:8080](http://localhost:8080)

### IDE Protip
Add build tags `js,wasm` to go tooling. In VSCode you may add a `.vscode` folder with a `settings.json` configuration file with the following contents:

```json
{
    "go.buildTags": "js,wasm"
}
```

This will cause intellisense to work correctly.