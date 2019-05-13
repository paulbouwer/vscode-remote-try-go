# Hack Notes

When using the **Remote Development - Containers** extension on a non-trivial customer project we ran into the following situation:

- Build using `go` worked from the VS Code Terminal
- The `vendor` folders in the code base were not found by the VS Code Go tooling
- The VS Code Debugger did not work, due to the `vendor` folders not being found

## Why is this important?

The **Remote Development - Containers** extension is a GREAT way to develop Go applications/services. The issues related to `$GOPATH` are largely isolated by developing in a container, and the extension makes that trivial.

It seems like there is a small issue, that when resolved will bring all the Go capabilities of VS Code to the **Remote Development - Containers** extension world and make Go projects super easy to get up and running for new starters on a project.

## Solution

To resolve the above issues, the `vscode-remote-try-go` repo was updated as follows to demonstrate the likely cause.

## Repo/Code Setup

The setup has been mapped to this repo as follows:

### Binary

The `cmd/server/main.go` represents the resulting binary, and is in the `cmd/binary-name/main.go` location.

### Dependencies

Dep was used to managed dependencies. It results in a `vendor` folder that is excluded from git commits via the `.gitignore` file.

The command `dep ensure` that will pull down the dependencies into the `vendor` folder is available via the `make install-dependencies` command. This make command is also called as part of `make build` to ensure the code builds with least amount of effort.

Before you can use the debugger, you will have to run `make install-dependencies` to ensure those have been pulled down.

### Debugger

The `.vscode/launch.json` launch configuration was updated to reference the new location of the binary in `the `cmd/binary-name/main.go` location.

## VS Code Setup

VS Code has been set up as follows:

### Remote Development - Containers Extension

Used the Remote Development - Containers Extension and ensured the following:

1. `.dvcontainer/Dockerfile` file added Dep, some additional utilities, and most importantly added the `cd /go/src/vscode-remote-try-go` command to the `.bashrc` to ensure that the VS Code Terminal opened in the symlinked folder and not the mounted folder. 

2. `.devcontainer/settings.vscode.json` file had gopath set to `/go` for the container.

3. `.devcontainer/devcontainer.json` file had additional `postCreateCommand` to symlink the default mount location of the source code in the container, to the `/go/src/` location in the container. This was to ensure that ALL tooling would work correctly since code was located in `$GOPATH`. This could not be done in the Dockerfile and container since the symlink would have been created before the code was mounted and then `ln -s` got into some recursive wierdness.

Even with all of these items configured, there was still an issue with where VS Code thought the workspace for the source code was.

VS Code still thought that all the files were located at `/workspaces/vscode-remote-try-go` and so because `$GOPATH` was set to `/go`, none of the `vendor` folders could be located by VS Code. Again, the `go` tools all worked correctly in the VS Code Terminal since `$GOPATH` was set to `/go` and the folder `/go/src/vscode-remote-try-go` was understood by the terminal.

When hovering over the tab of any opened file in VS Code, the path was prefixed with `/workspaces/vscode-remote-try-go` proving that VS Code was not able to understand that the working folder was instead `/go/src/vscode-remote-try-go`.

I was not able to find any additional VS Code settings to change the workspace location.

### Multi-root Workspaces

I added a `.code-workspace.code-workspace` file to the repo and manipulated this feature to effectively force a new workspace location onto VS Code.

By doing this, all the tabs of opened files now showed `/go/src/vscode-remote-try-go` as the file location. We could also still build on the command line, and the debugger was operational.

There are still some bits where I don't think the correct workspace `/go/src/vscode-remote-try-go` is being referenced, but by getting to this point, I think I've proved that if there was a way to reset the workspace for ALL of VS Code, then this extension would work solidly for most Go projects using this extension.

